package handshake

// handshake is an initial communication between two peers (a client and a server)
// that establishes a connection. The handshake is crucial for peers to exchange
// information about the torrent they are sharing and to ensure compatibility between them.

// HandShake represents the structure of a BitTorrent protocol handshake.
type HandShake struct {
	HeaderText string   // HeaderText is the protocol identifier, usually set to "BitTorrent protocol".
	InfoHash   [20]byte // InfoHash is the SHA-1 hash of the info key in the torrent file.
	PeerId     [20]byte // PeerId uniquely identifies the client making the handshake.
}

// New creates a new HandShake instance with the provided infoHash and peerId.
func New(infoHash, peerId [20]byte) HandShake {
	return HandShake{
		HeaderText: "BitTorrent protocol",
		InfoHash:   infoHash,
		PeerId:     peerId,
	}
}

// ToBytes converts the HandShake struct to a byte slice.
// The handshake message consists of:
//   - 1 byte representing the length of the header text
//   - The header text itself
//   - 8 reserved bytes (usually set to zero)
//   - 20 bytes representing the info hash of the torrent
//   - 20 bytes representing the peer id of the client
// The total length of the handshake message is thus len(headerText) + 49 bytes.
func (handShake *HandShake) ToBytes() []byte {
	buf := make([]byte, len(handShake.HeaderText)+49)
	buf[0] = byte(len(handShake.HeaderText))
	curr := 1
	curr += copy(buf[curr:], []byte(handShake.HeaderText))
	curr += copy(buf[curr:], make([]byte, 8)) // 8 reserved bytes
	curr += copy(buf[curr:], handShake.InfoHash[:])
	curr += copy(buf[curr:], handShake.PeerId[:])
	return buf
}
