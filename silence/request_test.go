package silence

import (
	"bytes"
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
		Body:           &NullRequestBody{},
	}
	b := m.Marshall()

	// BUILD A SECOND RequestMessage AND UNMARSHALL THE FIRST ONE INTO IT
	m2 := RequestMessage{}
	m2.Unmarshall(b)
	if m.Type != m2.Type {
		t.Fatalf("Message Types don't match")
	}
	if m.SequenceNumber != m2.SequenceNumber {
		t.Fatalf("Message sequence numbers don't match")
	}
	if m.Nonce != m2.Nonce {
		t.Fatalf("Message nonces don't match!")
	}
	b1 := m.Body.Marshall()
	b2 := m2.Body.Marshall()
	if bytes.Compare(b1, b2) != 0 {
		t.Fatalf("%v should equal %v\n", b1, b2)
	}

}