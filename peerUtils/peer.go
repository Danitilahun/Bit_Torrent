package peerutils

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/Danitilahun/Bit_Torrent/common"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
)

func getTrackerRequestUrl(manifest torrentmodels.TorrentManifest, announce string, peerId [20]byte, port int) (string, error) {
	baseUrl, err := url.Parse(announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(manifest.InfoHash[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(int(manifest.TotalLength))},
	}

	baseUrl.RawQuery = params.Encode()
	return baseUrl.String(), nil
}

func ConnectToPeer(peerAddress peer.PeerAddress, port int, timeout time.Duration) (conn net.Conn, err error) {
	return net.DialTimeout("tcp", peerAddress.IP.String()+":"+strconv.Itoa(int(peerAddress.Port)), timeout)
}

func EstablishConnection(peerAddress peer.PeerAddress, manifest torrentmodels.TorrentManifest) (peerInstance *peer.Peer) {
	var conn net.Conn = nil
	var timeout = time.Duration(10 * time.Second)

	fmt.Printf("Connecting to peer %v:%v\n", peerAddress.IP, peerAddress.Port)

	for conn == nil {
		timeout *= 2
		conn, _ = ConnectToPeer(peerAddress, common.Port, timeout)

		if timeout > time.Duration(60*time.Second) {
			fmt.Printf("Can't connect to peer %v:%v\n", peerAddress.IP, peerAddress.Port)
			return nil
		}
	}

	fmt.Printf("Connected to peer %v:%v\n", peerAddress.IP, peerAddress.Port)

	// When a peer is "choking" another peer, it means that it is deliberately limiting the rate at which it sends data to that peer.
	// When a peer is "choked" by another peer, it means that it is currently being limited in the rate at which it can receive data from that peer.

	peerInstance = &peer.Peer{
		Conn:       conn,
		Address:    peerAddress,
		Interested: false,
		IsChoked:   true,
		IsChoking:  false,
		BitField:   make([]byte, len(manifest.PieceHashes)),
	}
	return peerInstance
}
