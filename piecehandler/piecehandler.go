// Package piecehandler provides utilities for handling and processing pieces in the BitTorrent protocol.
package piecehandler

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"github.com/Danitilahun/Bit_Torrent/download"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
	"os"
)

// GetPieceLength calculates the length of a piece based on the given index, piece length, and total length.
func GetPieceLength(index int, pieceLength int, totalLength int) int {
	if pieceLength < (totalLength - index*pieceLength) {
		return pieceLength
	}
	return totalLength - index*pieceLength
}

// ReadPieceMessage extracts piece information from a given payload in a piece message.
func ReadPieceMessage(payload []byte) (int, int, []byte, error) {
	if len(payload) < 8 {
		return 0, 0, nil, errors.New("invalid payload length during piece message")
	}
	pieceIndex := binary.BigEndian.Uint32(payload[0:4])
	begin := binary.BigEndian.Uint32(payload[4:8])
	block := payload[8:]
	return int(pieceIndex), int(begin), block, nil
}

// CheckPieceHash verifies if the hash of a piece matches the expected hash.
func CheckPieceHash(piece []byte, hash [20]byte) bool {
	sha1Hash := sha1.Sum(piece)
	return bytes.Equal(sha1Hash[:], hash[:])
}

// WritePieceToFile writes a piece to a file at the appropriate offset.
func WritePieceToFile(manifest *torrentmodels.TorrentManifest, pieceJobResult *download.PieceJobResult, blobFile *os.File) {
	pieceOffset := int64(pieceJobResult.PieceIndex) * manifest.PieceLength
	blobFile.WriteAt(pieceJobResult.PieceData, pieceOffset)
}

// WritePieceMessage creates a piece message with the provided index, begin offset, and block data.
func WritePieceMessage(index int, begin int, block []byte) *peercommunication.Message {
	payload := make([]byte, 8)
	binary.BigEndian.PutUint32(payload[0:4], uint32(index))
	binary.BigEndian.PutUint32(payload[4:8], uint32(begin))
	payload = append(payload, block...)
	return &peercommunication.Message{
		Type:    peercommunication.MsgTypePiece,
		Payload: payload,
	}
}
