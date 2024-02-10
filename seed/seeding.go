package seed

import (
	"fmt"

	"github.com/Danitilahun/Bit_Torrent/bitfield"
	messageutils "github.com/Danitilahun/Bit_Torrent/messageUtils"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
	"github.com/Danitilahun/Bit_Torrent/piecehandler"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
)

func HandleSeedingRequest(req *SeedRequest, blobFile []byte, currentBitField *bitfield.Bitfield, manifest *torrentmodels.TorrentManifest) {
	if req.Peer.IsChoking {
		messageutils.SendMessageWithRetry(req.Peer, peercommunication.Message{Type: peercommunication.MsgTypeUnChoke})
		return
	}

	index, begin, length, err := messageutils.ReadRequestMessage(req.Message.Payload)

	if err != nil {
		fmt.Printf("Error reading request message from peer %v:%v, %v\n", req.Peer.Address.IP, req.Peer.Address.Port, err)
		return
	}

	if index >= len(*currentBitField) {
		fmt.Printf("Received request message from peer %v:%v with invalid index %v\n", req.Peer.Address.IP, req.Peer.Address.Port, index)
		return
	}

	if !(*currentBitField).HasPiece(index) {
		fmt.Printf("Received request message from peer %v:%v for a piece that is not available at index %v\n", req.Peer.Address.IP, req.Peer.Address.Port, index)
		return
	}

	if begin+length > int(manifest.PieceLength) {
		fmt.Printf("Received request message from peer %v:%v with invalid begin %v\n", req.Peer.Address.IP, req.Peer.Address.Port, begin)
		return
	}

	pieceOffset := int64(index) * int64(manifest.PieceLength)
	blockOffset := int64(begin)
	block := blobFile[pieceOffset+blockOffset : pieceOffset+blockOffset+int64(length)]

	if err != nil {
		fmt.Printf("Error reading block from blob file %v\n", err)
		return
	}

	// Print a message after the block is sent
	fmt.Printf("Sent block to peer %v:%v, PieceIndex: %v, Begin: %v, Length: %v\n", req.Peer.Address.IP, req.Peer.Address.Port, index, begin, length)

	messageutils.SendMessageWithRetry(req.Peer, *piecehandler.WritePieceMessage(index, begin, block))
}
