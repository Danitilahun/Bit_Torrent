package peercommunication

type MessageType byte

const (
	// Tells a peer to stop sending requests.
	MsgTypeChoke MessageType = 0
	// Tells a peer that it can start sending requests.
	MsgTypeUnChoke MessageType = 1
	// Indicates interest in receiving data from a peer.
	MsgTypeInterested MessageType = 2
	// Indicates no interest in receiving further data from a peer.
	MsgTypeNotInterested MessageType = 3
	// Notifies peers that the sender has received a piece of data.
	MsgTypeHave MessageType = 4
	//  Provides a bitfield representing the data pieces a peer has.
	MsgTypeBitField MessageType = 5
	// Requests a piece of data from a peer.
	MsgTypeRequest MessageType = 6
	// Delivers a piece of data in response to a request.
	MsgTypePiece MessageType = 7
	// Cancels a previously made request.
	MsgTypeCancel MessageType = 8
	// A keep-alive message, likely to maintain the connection alive in periods of inactivity.
	MsgTypeKeepAlive MessageType = 9
)
