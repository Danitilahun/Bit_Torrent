package handShakeUtils

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Danitilahun/Bit_Torrent/handshake"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
)

func ReadHandShake(reader io.Reader) (*handshake.HandShake, error) {
	var headerText [20]byte
	var err error = nil

	_, err = io.ReadFull(reader, headerText[:])

	if err != nil {
		return nil, err
	}

	if string(headerText[:]) != string(rune(19))+"BitTorrent protocol" {
		return nil, errors.New("invalid header, expected 'BitTorrent protocol' but got '" + string(headerText[:]) + "'")
	}

	var reserved [8]byte
	_, err = io.ReadFull(reader, reserved[:])

	if err != nil {
		return nil, err
	}

	var infoHash [20]byte
	_, err = io.ReadFull(reader, infoHash[:])

	if err != nil {
		return nil, err
	}

	var peerId [20]byte
	_, err = io.ReadFull(reader, peerId[:])

	if err != nil {
		return nil, err
	}

	return &handshake.HandShake{
		InfoHash: infoHash,
		PeerId:   peerId,
	}, nil
}

func EstablishHandShake(peerId [20]byte, peer peer.Peer, manifest torrentmodels.TorrentManifest) (int, error) {
	peer.Conn.SetDeadline(time.Now().Add(5 * time.Second))
	// Reset deadline
	defer peer.Conn.SetDeadline(time.Time{})

	retries := 0
	var err error = nil
	handShake := handshake.New(manifest.InfoHash, peerId)

	for retries < 10 {
		if retries > 0 {
			fmt.Printf("Retrying handshake with peer %v:%v for %v time\n", peer.Address.IP, peer.Address.Port, retries)
		}

		_, err := peer.Conn.Write(handShake.ToBytes())
		if err == nil || err == io.EOF {
			return retries, err
		}

		retries++
	}

	return retries, err
}
