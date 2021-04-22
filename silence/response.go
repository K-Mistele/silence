package silence

// THE ResponseMessageType DEFINES TYPE CODES FOR SILENCE REPLIES
type ResponseMessageType uint8
const (
	ResponseMessageTypeNoop = 0x00
	ResponseMessageTypeExecuteCommands = 0x01

)

// THE ResponseMessage IS THE MESSAGE FOR A REPLY FROM THE SERVER
type ResponseMessage struct {
	Type ResponseMessageType
}

