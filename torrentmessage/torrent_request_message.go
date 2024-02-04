package torrentmessage

import (
	"encoding/binary"
)

// RequestMessage represents a message requesting a specific piece of data from a peer.
type RequestMessage struct {
	PieceIndex int
	Begin      int
	Length     int
}

// ToBytes converts the RequestMessage to its byte representation.
func (req RequestMessage) ToBytes() []byte {
	bytes := make([]byte, 0)
	bytes = binary.BigEndian.AppendUint32(bytes, uint32(req.PieceIndex))
	bytes = binary.BigEndian.AppendUint32(bytes, uint32(req.Begin))
	bytes = binary.BigEndian.AppendUint32(bytes, uint32(req.Length))
	return bytes
}
