package silence

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

// TESTS THE MARSHALLING OF THE MESSAGE
func TestRequestMarshall(t *testing.T) {

	// BUILD A RequestMessage AND MARSHALL IT TO BYTES
	m := RequestMessage{
		Type:           ReadyForCommand,
		SequenceNumber: 1,
		Nonce:          0x41424344,
		Body:           &NullRequestBody{},
	}
	b := m.Marshall()

	// BUILD WHAT WE THINK RequestMessage SHOULD LOOK LIKE AND MAKE SURE THEY MATCH
	c := []byte {0x01, 0x01, 0x44, 0x43, 0x42, 0x41, 0x00 ^ 0x44, 0x00 ^ 0x43, 0x00 ^ 0x42, 0x00 ^ 0x41}
	if bytes.Compare(b, c) != 0 {
		t.Fatalf("Serializing request failed: expected %v amd got %v\n", c, b)
	}

}

func TestRequestUnmarshall(t *testing.T) {
	// BUILD A RequestMessage AND MARSHALL IT TO BYTES
	m := RequestMessage{
		Type:           ReadyForCommand,
		SequenceNumber: 1,
		Nonce:          0x41424344,
		Body:           &ReadyForCommandBody{},
	}
	b := m.Marshall()

	// BUILD A SECOND RequestMessage AND UNMARSHALL THE FIRST ONE INTO IT
	m2 := RequestMessage{}
	m2.Unmarshall(b)
	messagesMatch, errString := compareRequestMessages(&m, &m2)
	if !messagesMatch {
		t.Fatalf(errString)
	}


}

// TEST THE READY FOR COMMAND MESSAGE
func TestReadyForCommandMessage(t *testing.T) {
	body := &ReadyForCommandBody{}
	nonce := rand.Uint32()
	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, nonce)

	message := RequestMessage{
		Type: ReadyForCommand,
		SequenceNumber: 2,
		Nonce: nonce,
		Body: body,
	}

	b := message.Marshall()
	expected := []byte {uint8(ReadyForCommand), 2, nonceBytes[0], nonceBytes[1], nonceBytes[2], nonceBytes[3]}

	if bytes.Compare(b, expected) != 0 {
		t.Fatalf("Failed to construct a Ready for Command Message. Expected %v but got %v\n", expected, b)
	}

	message2 := RequestMessage{}
	message2.Unmarshall(b)
	messagesMatch, errString := compareRequestMessages(&message, &message2)
	if !messagesMatch {
		fmt.Printf("%+v\n%+v\n", message.Body, message2.Body)
		t.Fatalf(errString)

	}



}