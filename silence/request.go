package silence

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

type SilenceMessage interface {
	Marshall() []byte
	Unmarshall([]byte)
}

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
////////////////////////////////////////////////////////////////////
// RequestMessage type
////////////////////////////////////////////////////////////////////

// RequestMessage IS THE MESSAGE FOR A SILENCE REQUEST TO THE SERVER
// IMPLEMENTS SilenceMessage
type RequestMessage struct {
	Type 				RequestMessageType	// PROTOCOL MESSAGE TYPE
	SequenceNumber		uint8				// SEQUENCE NUMBER, GOES 0->1->...255->0->1...
	Nonce				uint32				// A RANDOM 32-BIT INTENER TO XOR WITH THE MESSAGE BODY
	Body				RequestMessageBody	// A REQUEST MESSAGE BODY, DEPENDING ON WHAT THE BODY IS
}

// Marshall WILL BUILD OUT THE RequestMessage INTO A STRING OF BYTES, PERFORMING ENCODING AS APPROPRIATE
func (r *RequestMessage) Marshall() []byte {

	headerBytes := make([]byte, 2)	// 6 BYTES FOR 48 BITS
	headerBytes[0] = uint8(r.Type)
	headerBytes[1] = r.SequenceNumber

	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, r.Nonce)
	headerBytes = append(headerBytes, nonceBytes...)

	bodyBytes := r.Body.Marshall()
	// TODO: XOR THIS WITH THE NONCE
	headerBytes = append(headerBytes, bodyBytes...)

	fmt.Println(headerBytes)

	return headerBytes
}

// Unmarshall WILL ATTEMPT TO UNMARSHALL THE BYTES INTO THE REQUEST MESSAGE. IF THERE'S AN ERROR, THE MESSAGE WILL BE NIL
func (r *RequestMessage) Unmarshall(data []byte) {

	// IF THE MESSAGE IS MALFORMED, RECOVER AND DON'T DO ANYTHING
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic while unmarshalling: ", r)
			r = nil
		}
	}()

	r.Type = RequestMessageType(data[0])
	r.SequenceNumber = data[1]
	r.Nonce = binary.LittleEndian.Uint32(data[2:6])

	// TODO DECRYPT THE BODY
	if r.Type == ReadyForCommand {
		r.Body = &NullRequestBody{}
		r.Body.Unmarshall(data[6:])
	} else {
		// DEFAULT
		r.Body = &NullRequestBody{}
	}


}

// NewRequestMessage WILL BUILD A NEW RequestMessage WITH A RANDOM NONCE
func NewRequestMessage(t RequestMessageType, seqNo uint8, body RequestMessageBody) *RequestMessage {
	return &RequestMessage{
		Type: t,
		SequenceNumber: seqNo,
		Nonce: rand.Uint32(),
		Body: body,
	}
}

////////////////////////////////////////////////////////////////////
// RequestMessageBody type
////////////////////////////////////////////////////////////////////

// RequestMessageBody DEFINES AN INTERFACE FOR DIFFERENT REQUEST MESSAGE BODY TYPES
type RequestMessageBody interface {
	Marshall() []byte
	Unmarshall([]byte)

}

// NullRequestBody IS A NULL BODY THAT'S JUST 4 NULL BYTES
type NullRequestBody struct {
	Data []byte
}

// Marshall WILL SERIALIZE THE NULL BODY AND RETURN IT
func (nb *NullRequestBody) Marshall() []byte {
	return []byte {0, 0, 0, 0}
}

// Unmarshall WILL UPDATE THE POINTER WITH THE PROPERTIES FROM THE BYTES
func (nb *NullRequestBody) Unmarshall(b []byte) {
	nb.Data = []byte {0, 0, 0, 0}
}

