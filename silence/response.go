package silence

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

// THE ResponseMessageType DEFINES TYPE CODES FOR SILENCE REPLIES
type ResponseMessageType uint8
const (
	ResponseMessageTypeNoop 			ResponseMessageType = 0x00
	ResponseMessageTypeExecuteCommands 	ResponseMessageType = 0x01

	// OxFO-0xFF ARE ERROR CODES
	ResponseMessageTypeErrorWithDebug 	ResponseMessageType = 0xF0
	ResponseMessageTypeErrorGoBack		ResponseMessageType = 0xF1

)
////////////////////////////////////////////////////////////////////
// ResponseMessage type
////////////////////////////////////////////////////////////////////

// THE ResponseMessage IS THE MESSAGE FOR A REPLY FROM THE SERVER
// IMPLEMENTS SilenceMessage
type ResponseMessage struct {
	Type 			SilenceMessageType		// DEFINES THE MESSAGE TYPE
	Code           	ResponseMessageType 	// PROTOCOL MESSAGE CODE FOR THE TYPE
	SequenceNumber 	uint8               	// SEQUENCE NUMBER, 0-255 WITH WRAPAROUND
	AckNumber      	uint8               	// THE SEQUENCE NUMBER FROM THE LAST CLIENT MESSAGE RECEIVED
	Nonce          	uint32              	// A RANDOM 32-BIT INTEGER TO XOR WITH THE MESSAGE
	Body           	ResponseMessageBody 	// A RESPONSE MESSAGE BODY DEPENDING ON THE MESSAGE TYPE
}

// Marshall WILL BUILD OUT THE ResponseMessage INTO A STRING OF BYTES, PERFORM ENCODING OF THE PAYLOAD
func (r *ResponseMessage) Marshall() ([]byte, error) {

	var messageBytes []byte
	messageBytes = make([]byte, 4)
	messageBytes[0] = uint8(r.Type)
	messageBytes[1] = uint8(r.Code)
	messageBytes[2] = r.SequenceNumber
	messageBytes[3] = r.AckNumber

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

// Unmarshall WILL ATTEMPT TO UNMARSHALL THE BYTES INTO THE REQUEST MESSAGE. IF THERE'S AN ERROR, THE MESSAGE WILL BE NILL
func (r *ResponseMessage) Unmarshall(data []byte) (err interface{} ) {
	err = nil
	// CATCH A PANIC IF WE FAIL TO DECODE PROPERLY
	defer func(){
		if p := recover(); p != nil {
			fmt.Println("Recovered from panic while unmarshalling: ", p)
			r = nil
			err = p
		}
	}()

	r.Type = SilenceMessageType(data[0])
	r.Code = ResponseMessageType(data[1])
	r.SequenceNumber = data[2]
	r.AckNumber = data[3]
	r.Nonce = binary.LittleEndian.Uint32(data[4:8])

	payload := data[8:]
	decoded := xorDecode(&payload, r.Nonce)

	if r.Code == ResponseMessageTypeExecuteCommands {
		r.Body = &ResponseBodyExecuteCommands{}
		e := r.Body.Unmarshall(decoded)
		if e != nil {
			r = nil
			return e
		}
	}

	return nil
}

// NewResponseMessage BUILDS A NEW RESPONSE MESSAGE
func NewResponseMessage(t ResponseMessageType, seqNo uint8, ack uint8, body ResponseMessageBody) *ResponseMessage {

	return &ResponseMessage{
		Type: 			SilenceMessageResponse,
		Code:           t,
		SequenceNumber: seqNo,
		AckNumber:      ack,
		Nonce:          rand.Uint32(),
		Body:           body,
	}
}

