package peer

import "net"

// PeerAddress represents the address information of a peer.
type PeerAddress struct {
	// IP is the IP address of the peer.
	IP net.IP
	// Port is the port number of the peer.
	Port uint16
}
