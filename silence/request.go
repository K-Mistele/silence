package silence

import (
	"encoding/binary"
	"fmt"
)
// RequestMessageType DEFINES TYPE CODES FOR SILENCE REQUESTS
type RequestMessageType uint8
const (

	ReadyForCommand			RequestMessageType = 0x01
	CommandAcknowledged		RequestMessageType = 0x02
	CommandOutput			RequestMessageType = 0x03

	// 0xF0-0xFF ARE ERROR CODES
	ErrorWithDebug			RequestMessageType = 0xF0
	ErrorGoBackToMessage	RequestMessageType = 0xF1

)

// RequestMessageBody DEFINES AN INTERFACE FOR DIFFERENT REQUEST MESSAGE BODY TYPES
type RequestMessageBody interface {
	Marshall() []byte
}

// RequestMessage IS THE MESSAGE FOR A SILENCE REQUEST TO THE SERVER
type RequestMessage struct {
	Type 				RequestMessageType	// PROTOCOL MESSAGE TYPE
	SequenceNumber		uint8				// SEQUENCE NUMBER, GOES 0->1->...255->0->1...
	Nonce				uint32				// A RANDOM 32-BIT INTENER TO XOR WITH THE MESSAGE BODY
	Body				RequestMessageBody	// A REQUEST MESSAGE BODY, DEPENDING ON WHAT THE BODY IS
}

// Marshall WILL BUILD OUT THE RequestMessage INTO A STRING OF BYTES, PERFORMING ENCODING AS APPROPRIATE
func (r *RequestMessage) Marshall() []byte {

	headerBytes := make([]byte, 6)	// 6 BYTES FOR 48 BITS
	headerBytes[0] = uint8(r.Type)
	headerBytes[1] = r.SequenceNumber

	binary.LittleEndian.PutUint32(headerBytes[2:6], r.Nonce)
	fmt.Println(headerBytes)

	return headerBytes
}

// NullBody IS A NULL BODY THAT'S JUST 4 NULL BYTES
type NullBody struct {
}

// Marshall WILL SERIALIZE THE NULL BODY AND RETURN IT
func (nb *NullBody) Marshall() []byte {
	return []byte {0, 0, 0, 0}
}
