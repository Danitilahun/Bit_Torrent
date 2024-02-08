package seed

import (
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/peercommunication"
)

type SeedRequest struct {
	Peer    *peer.Peer
	Message *peercommunication.Message
}
