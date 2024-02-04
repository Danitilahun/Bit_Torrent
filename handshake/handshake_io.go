package handshake

// Package handshake provides functionality related to BitTorrent protocol handshakes.

import "io"

// ReadHandShake reads a HandShake from the given reader.
// It attempts to read the components of a BitTorrent protocol handshake message from the provided io.Reader.
// The expected structure of the handshake message is as follows:
//   - 20 bytes representing the header text (protocol identifier).
//   - 8 reserved bytes (usually set to zero).
//   - 20 bytes representing the info hash of the torrent.
//   - 20 bytes representing the peer id of the client.
// The function returns a HandShake struct containing the parsed information or an error if the reading process fails.
func ReadHandShake(reader io.Reader) (HandShake, error) {
	var headerText [20]byte
	var reserved [8]byte
	var infoHash [20]byte
	var peerId [20]byte

	_, err := reader.Read(headerText[:])
	if err != nil {
		return HandShake{}, err
	}
	_, err = reader.Read(reserved[:])
	if err != nil {
		return HandShake{}, err
	}
	_, err = reader.Read(infoHash[:])
	if err != nil {
		return HandShake{}, err
	}
	_, err = reader.Read(peerId[:])
	if err != nil {
		return HandShake{}, err
	}

	return HandShake{
		HeaderText: string(headerText[:]),
		InfoHash:   infoHash,
		PeerId:     peerId,
	}, nil
}
