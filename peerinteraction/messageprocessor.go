package peerinteraction

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Danitilahun/Bit_Torrent/download"
	messageutils "github.com/Danitilahun/Bit_Torrent/messageUtils"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
	"github.com/Danitilahun/Bit_Torrent/piecehandler"
	"github.com/Danitilahun/Bit_Torrent/seed"
)

func processIncomingMessages(peer *peer.Peer, connReader io.Reader, progress *download.PieceJobProgress, seedRequestChannel *chan *seed.SeedRequest) (peercommunication.MessageType, error) {
	message, err := messageutils.ReadMessage(connReader)

	if err != nil {
		fmt.Printf("Error reading message from peer %v:%v, %v\n", peer.Address.IP, peer.Address.Port, err)
		return peercommunication.MsgTypeKeepAlive, err
	}

	if message == nil {
		return peercommunication.MsgTypeKeepAlive, nil
	}

	if message.Type != peercommunication.MsgTypePiece {
		fmt.Printf("Received message from peer %v:%v, %v\n", peer.Address.IP, peer.Address.Port, message.Type.String())
	}

	switch message.Type {
	case peercommunication.MsgTypeUnChoke:
		peer.IsChoking = false
	case peercommunication.MsgTypeChoke:
		peer.IsChoking = true
	case peercommunication.MsgTypeInterested:
		peer.Interested = true
	case peercommunication.MsgTypeNotInterested:
		peer.Interested = false
	case peercommunication.MsgTypeHave:
		pieceIndex := binary.BigEndian.Uint32(message.Payload)
		peer.BitField.MarkPiece(int(pieceIndex))
	case peercommunication.MsgTypeBitField:
		peer.BitField = message.Payload
	case peercommunication.MsgTypeCancel:
		fmt.Printf("Received cancel message from peer %v:%v\n", peer.Address.IP, peer.Address.Port)
	case peercommunication.MsgTypePiece:
		index, begin, block, err := piecehandler.ReadPieceMessage(message.Payload)
		if err != nil {
			fmt.Printf("Error reading piece job result from peer %v:%v, %v\n", peer.Address.IP, peer.Address.Port, err)
			return peercommunication.MsgTypePiece, err
		}

		if progress == nil {
			fmt.Printf("Received piece job result from peer %v:%v with no job in progress\n", peer.Address.IP, peer.Address.Port)
			return peercommunication.MsgTypePiece, err
		}

		if index != progress.PieceIndex {
			fmt.Printf("Received piece job result from peer %v:%v with wrong piece index %v\n", peer.Address.IP, peer.Address.Port, index)
			return peercommunication.MsgTypePiece, err
		}

		if begin+len(block) > progress.PieceLength {
			fmt.Printf("Received piece job result from peer %v:%v with wrong begin %v\n", peer.Address.IP, peer.Address.Port, begin)
			return peercommunication.MsgTypePiece, err
		}

		progress.TotalDownloaded += len(block)
		copy(progress.Buffer[begin:], block)
	case peercommunication.MsgTypeRequest:
		*seedRequestChannel <- &seed.SeedRequest{
			Peer:    peer,
			Message: message,
		}
		return peercommunication.MsgTypeRequest, nil
	}

	return message.Type, nil
}
