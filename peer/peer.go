package peer

import (
	"github.com/Danitilahun/Bit_Torrent/bitfield"
	"net"
)

// Peer represents a peer in a peer-to-peer network.
type Peer struct {
	// Conn is the network connection to the peer.
	Conn net.Conn
	// Address represents the address information of the peer.
	Address PeerAddress
	// Interested indicates whether the peer is interested in the content.
	Interested bool
	// IsChoked is true if the peer is not allowed to send data to us.
	IsChoked bool
	// IsChoking is true if we are not allowed to send data to the peer.
	IsChoking bool
	// BitField is a list of booleans indicating whether the peer has the corresponding piece.
	BitField bitfield.Bitfield
}
