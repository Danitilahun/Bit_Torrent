package peerinteraction

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Danitilahun/Bit_Torrent/bitfield"
	"github.com/Danitilahun/Bit_Torrent/common"
	"github.com/Danitilahun/Bit_Torrent/download"
	"github.com/Danitilahun/Bit_Torrent/handShakeUtils"
	messageutils "github.com/Danitilahun/Bit_Torrent/messageUtils"
	"github.com/Danitilahun/Bit_Torrent/peer"
	peerutils "github.com/Danitilahun/Bit_Torrent/peerUtils"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
	"github.com/Danitilahun/Bit_Torrent/piecehandler"
	"github.com/Danitilahun/Bit_Torrent/seed"
	"github.com/Danitilahun/Bit_Torrent/torrentmessage"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
)

func StartPeerWorker(peers []*peer.Peer, i int, peerAddress peer.PeerAddress, peerId [20]byte, manifest torrentmodels.TorrentManifest, port int, workChannel *chan download.PieceJob, currentBitField *bitfield.Bitfield, pieceJobResultChannel *chan *download.PieceJobResult, seedRequestChannel *chan *seed.SeedRequest, conn *net.Conn) {
	// Establish connection
	var p *peer.Peer = nil

	if conn != nil {
		p = &peer.Peer{
			Address:    peerAddress,
			Conn:       *conn,
			Interested: false,
			IsChoking:  true,
			IsChoked:   false,
			BitField:   make(bitfield.Bitfield, len(*currentBitField)),
		}
	} else {
		p = peerutils.EstablishConnection(peerAddress, manifest)
	}

	if p == nil {
		return
	}
	peers[i] = p
	defer p.Conn.Close()

	// Establish handshake
	_, err := handShakeUtils.EstablishHandShake(peerId, *p, manifest)
	if err != nil {
		fmt.Printf("Error establishing handshake with peer %v:%v\n", p.Address.IP, p.Address.Port)
		return
	}

	connReader := io.Reader(p.Conn)

	// Read handshake
	shouldReturn := readHandShake(connReader, p, manifest)
	if shouldReturn {
		return
	}

	// Send Interested
	_, err = messageutils.SendMessageWithRetry(p, peercommunication.Message{
		Type: peercommunication.MsgTypeInterested,
	})
	if err != nil {
		fmt.Printf("Error sending interested to peer %v:%v\n", p.Address.IP, p.Address.Port)
		return
	}
	fmt.Printf("Interested sent to peer %v:%v\n", p.Address.IP, p.Address.Port)

	err = messageutils.SendUnchokeMessage(p)
	if err != nil {
		fmt.Printf("Error sending unchoke to peer %v:%v\n", p.Address.IP, p.Address.Port)
		return
	}

	// Receive bitfield
	_, err = processIncomingMessages(p, connReader, nil, seedRequestChannel)
	if err != nil {
		fmt.Printf("Error processing incoming messages from peer %v:%v, %v\n", p.Address.IP, p.Address.Port, err)
		return
	}

	for {
		// Check if peer is choking us and wait for unchoke
		if p.IsChoking {
			time.Sleep(1 * time.Second)
			_, err = processIncomingMessages(p, connReader, nil, seedRequestChannel)
			if err != nil {
				fmt.Printf("Error processing incoming messages from peer %v:%v, %v\n", p.Address.IP, p.Address.Port, err)
				return
			}
			continue
		}

		for pieceJob := range *workChannel {
			// Check if peer has piece
			if !p.BitField.HasPiece(pieceJob.PieceIndex) {
				*workChannel <- pieceJob
				continue
			}

			if p.IsChoking {
				*workChannel <- pieceJob
				break
			}

			fmt.Printf("Sending piece job to peer %v:%v, piece index %v\n", p.Address.IP, p.Address.Port, pieceJob.PieceIndex)

			pieceJobProgress := download.PieceJobProgress{
				PieceIndex:      pieceJob.PieceIndex,
				Buffer:          make([]byte, pieceJob.PieceLength),
				TotalDownloaded: 0,
				PieceLength:     pieceJob.PieceLength,
			}

			p.Conn.SetDeadline(time.Now().Add(30 * time.Second))

			for pieceJobProgress.TotalDownloaded < pieceJob.PieceLength {
				if p.IsChoking {
					*workChannel <- pieceJob
					break
				}

				currentBlockSize := common.BlockSize
				if pieceJob.PieceLength-pieceJobProgress.TotalDownloaded < common.BlockSize {
					currentBlockSize = pieceJob.PieceLength - pieceJobProgress.TotalDownloaded
				}

				// Send request
				_, err = messageutils.SendMessageWithRetry(p, peercommunication.Message{
					Type: peercommunication.MsgTypeRequest,
					Payload: torrentmessage.RequestMessage{
						PieceIndex: pieceJob.PieceIndex,
						Begin:      pieceJobProgress.TotalDownloaded,
						Length:     currentBlockSize,
					}.ToBytes(),
				})

				if err != nil {
					fmt.Printf("Error sending request to peer %v:%v\n", p.Address.IP, p.Address.Port)
					*workChannel <- pieceJob
					return
				}

				var msgType peercommunication.MessageType = peercommunication.MsgTypeKeepAlive

				for msgType != peercommunication.MsgTypePiece {
					msgType, err = processIncomingMessages(p, connReader, &pieceJobProgress, seedRequestChannel)
					if p.IsChoking {
						*workChannel <- pieceJob
						break
					}

					if err != nil {
						fmt.Printf("Error processing incoming messages from peer %v:%v, %v\n", p.Address.IP, p.Address.Port, err)
						*workChannel <- pieceJob
						return
					}
				}
			}

			// Done with piece job
			// Reset deadline
			p.Conn.SetDeadline(time.Time{})

			if pieceJobProgress.TotalDownloaded == pieceJob.PieceLength {
				// check if piece is valid
				if !piecehandler.CheckPieceHash(pieceJobProgress.Buffer, pieceJob.PieceHash) {
					fmt.Printf("Piece hash doesn't match for piece %v\n", pieceJob.PieceIndex)
					*workChannel <- pieceJob
					continue
				}

				*pieceJobResultChannel <- &download.PieceJobResult{
					PieceIndex: pieceJob.PieceIndex,
					PieceData:  pieceJobProgress.Buffer,
				}
			}
		}
	}
}
