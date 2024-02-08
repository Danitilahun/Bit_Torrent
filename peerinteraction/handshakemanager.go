package peerinteraction

import (
	"bytes"
	"fmt"
	"io"

	"github.com/Danitilahun/Bit_Torrent/handShakeUtils"
	messageutils "github.com/Danitilahun/Bit_Torrent/messageUtils"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
)

func readHandShake(connReader io.Reader, peer *peer.Peer, manifest torrentmodels.TorrentManifest) bool {
	handshake, err := handShakeUtils.ReadHandShake(connReader)
	if err != nil {
		fmt.Printf("Error reading handshake from peer %v:%v, %v\n", peer.Address.IP, peer.Address.Port, err)
		return true
	}

	if !bytes.Equal(handshake.InfoHash[:], manifest.InfoHash[:]) {
		fmt.Printf("Handshake info hash doesn't match with manifest info hash from peer %v:%v\n", peer.Address.IP, peer.Address.Port)
		return true
	}
	fmt.Printf("Handshake established with peer %v:%v\n", peer.Address.IP, peer.Address.Port)
	return false
}

func sendChoke(peer *peer.Peer) bool {
	_, err := messageutils.SendMessageWithRetry(peer, peercommunication.Message{
		Type: peercommunication.MsgTypeChoke,
	})
	if err != nil {
		fmt.Printf("Error sending choke to peer %v:%v\n", peer.Address.IP, peer.Address.Port)
		return true
	}
	fmt.Printf("Choke sent to peer %v:%v\n", peer.Address.IP, peer.Address.Port)
	return false
}
