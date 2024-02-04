package torrentmessage

import (
	"encoding/binary"
	"errors"
)

// ParseRequestMessage parses a byte slice into a RequestMessage.
func ParseRequestMessage(payload []byte) (RequestMessage, error) {
	if len(payload) < 12 {
		return RequestMessage{}, errors.New("invalid payload length for request message")
	}

	// Extract information from the payload
	pieceIndex := binary.BigEndian.Uint32(payload[0:4])
	begin := binary.BigEndian.Uint32(payload[4:8])
	length := binary.BigEndian.Uint32(payload[8:12])

	// Create and return a RequestMessage
	return RequestMessage{
		PieceIndex: int(pieceIndex),
		Begin:      int(begin),
		Length:     int(length),
	}, nil
}
