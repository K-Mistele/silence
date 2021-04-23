package silence

type SilenceMessage interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

// SilenceMessageType DEFINES THE CODES THE IDENTIFY WHICH MESSAGE THIS IS, EITHER  A REQUEST OR RESPONSE
// PROBABLY REDUNDANT BECAUSE OF UNDERLYING TRANSPORT ON ICMP WHICH IS DIRECTIONAL, BUT ALLOWS FOR EXTENSIBILITY
type SilenceMessageType uint8
const (
	SilenceMessageRequest		SilenceMessageType = 0x01
	SilenceMessageResponse		SilenceMessageType = 0x02
)

