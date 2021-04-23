package silence

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

type SilenceMessage interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

// RequestMessageType DEFINES TYPE CODES FOR SILENCE REQUESTS
type RequestMessageType uint8

const (
	RequestMessageTypeReadyForCommand     RequestMessageType = 0x01
	RequestMessageTypeCommandAcknowledged RequestMessageType = 0x02
	RequestMessageTypeCommandOutput       RequestMessageType = 0x03

	// 0xF0-0xFF ARE ERROR CODES
	RequestMessageTypeErrorWithDebug RequestMessageType = 0xF0
	RequestMessageTypeErrorGoBack    RequestMessageType = 0xF1
)

////////////////////////////////////////////////////////////////////
// RequestMessage type
////////////////////////////////////////////////////////////////////

// RequestMessage IS THE MESSAGE FOR A SILENCE REQUEST TO THE SERVER
// IMPLEMENTS SilenceMessage
type RequestMessage struct {
	Type           	RequestMessageType // PROTOCOL MESSAGE TYPE
	SequenceNumber 	uint8              // SEQUENCE NUMBER, GOES 0->1->...255->0->1...
	AckNumber		uint8				// THE SEQUENCE NUMBER FROM THE SERVER LAST RECEIVED
	Nonce          	uint32             // A RANDOM 32-BIT INTEGER TO XOR WITH THE MESSAGE BODY
	Body           	RequestMessageBody // A REQUEST MESSAGE BODY, DEPENDING ON WHAT THE BODY IS
}

// Marshall WILL BUILD OUT THE RequestMessage INTO A STRING OF BYTES, PERFORMING ENCODING AS APPROPRIATE
func (r *RequestMessage) Marshall() ([]byte, error) {

	var messageBytes []byte
	messageBytes = make([]byte, 3) // 6 BYTES FOR 48 BITS
	messageBytes[0] = uint8(r.Type)
	messageBytes[1] = r.SequenceNumber
	messageBytes[2] = r.AckNumber

	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, r.Nonce)
	messageBytes = append(messageBytes, nonceBytes...)

	bodyBytes, err := r.Body.Marshall()
	if err != nil {
		return messageBytes, err
	}
	encodedBytes := xorEncode(&bodyBytes, r.Nonce)
	messageBytes = append(messageBytes, encodedBytes...)

	return messageBytes, nil
}

// Unmarshall WILL ATTEMPT TO UNMARSHALL THE BYTES INTO THE REQUEST MESSAGE. IF THERE'S AN ERROR, THE MESSAGE WILL BE NIL
func (r *RequestMessage) Unmarshall(data []byte) (err interface{}) {

	err = nil
	// IF THE MESSAGE IS MALFORMED, RECOVER AND DON'T DO ANYTHING
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered from panic while unmarshalling: ", p)
			r = nil
			err = p
		}
	}()

	// PULL HEADER FIELDS OUT
	r.Type = RequestMessageType(data[0])
	r.SequenceNumber = data[1]
	r.AckNumber = data[2]
	r.Nonce = binary.LittleEndian.Uint32(data[3:7])

	// GET THE PAYLOAD SLICE AND DECODE IT - XOR BY NONCE
	payload := data[7:]
	decoded := xorDecode(&payload, r.Nonce)

	if r.Type == RequestMessageTypeReadyForCommand {
		r.Body = &RequestBodyReadyForCommand{}
		r.Body.Unmarshall(decoded)
	} else {
		// DEFAULT
		r.Body = &RequestBodyNull{}
	}

	return err
}

// NewRequestMessage WILL BUILD A NEW RequestMessage WITH A RANDOM NONCE
func NewRequestMessage(t RequestMessageType, seqNo uint8, ack uint8, body RequestMessageBody) *RequestMessage {
	return &RequestMessage{
		Type:           t,
		SequenceNumber: seqNo,
		AckNumber: 		ack,
		Nonce:          rand.Uint32(),
		Body:           body,
	}
}

////////////////////////////////////////////////////////////////////
// RequestMessageBody type
////////////////////////////////////////////////////////////////////

// RequestMessageBody DEFINES AN INTERFACE FOR DIFFERENT REQUEST MESSAGE BODY TYPES
type RequestMessageBody interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

// RequestBodyNull IS A NULL BODY THAT'S JUST 4 NULL BYTES - USED WHEN ONLY THE CODE IS IMPORTANT
type RequestBodyNull struct {
	Data []byte
}

// Marshall WILL SERIALIZE THE NULL BODY AND RETURN IT
func (nb *RequestBodyNull) Marshall() ([]byte, error) {
	return []byte{0, 0, 0, 0}, nil
}

// Unmarshall WILL UPDATE THE POINTER WITH THE PROPERTIES FROM THE BYTES
func (nb *RequestBodyNull) Unmarshall(b []byte) interface{}{
	nb.Data = []byte{0, 0, 0, 0}
	return nil
}

// RequestBodyReadyForCommand IS EMPTY SINCE NO DATA IS NEEDED
type RequestBodyReadyForCommand struct{}

// Marshall WILL SERIALIZE IT TO NOTHING :)
func (rcb *RequestBodyReadyForCommand) Marshall() ([]byte, error) {
	return []byte{}, nil
}

// Unmarshall WILL DESERIALIZE IT TO NOTHING :)
func (rcb *RequestBodyReadyForCommand) Unmarshall(b []byte) interface{} {
	return nil
}
