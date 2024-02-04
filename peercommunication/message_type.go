package peercommunication

type MessageType byte

const (
	MsgTypeChoke         MessageType = 0
	MsgTypeUnChoke       MessageType = 1
	MsgTypeInterested    MessageType = 2
	MsgTypeNotInterested MessageType = 3
	MsgTypeHave          MessageType = 4
	MsgTypeBitField      MessageType = 5
	MsgTypeRequest       MessageType = 6
	MsgTypePiece         MessageType = 7
	MsgTypeCancel        MessageType = 8
	MsgTypeKeepAlive     MessageType = 9
)
