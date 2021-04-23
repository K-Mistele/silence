package silence

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// xorEncode WILL ENCODE THE BYTES IN A SLICE BY XORING THEM WITH THE KEY
func xorEncode(bytes *[]byte, key uint32) []byte {

	// SKIP THE ENCODING IF THE PACKAGE IS CONFIGURED NOT TO DO IT
	if !encodeMesages {
		return *bytes
	}

	encoded := make([]byte, len(*bytes))

	// MAKE THE KEY INTO A SLICE OF BYTES
	keySlice := make([]byte, 4)
	binary.LittleEndian.PutUint32(keySlice, key)

	// ITERATE ACROSS THE SLICE AND SET EACH BYTE TO ITSELF XORED WITH THE APPROPRIATE BYTE OF THE KEY
	for i := range *bytes {
		encoded[i]	= (*bytes)[i] ^ keySlice[i % 4]
	}
	return encoded
}

// xorDecode IS AN ALIAS FOR xorEncode - THAT'S HOW XOR ENCODING/DECODING WORK
func xorDecode(bytes *[]byte, key uint32) []byte {
	return xorEncode(bytes, key)
}

// compareRequestMessages RETURNS IF TWO MESSAGES MATCH
func compareRequestMessages (m *RequestMessage, m2 *RequestMessage) (bool, string){
	if m.Type != m2.Type {
		return false, "Message Types don't match"
	}
	if m.SequenceNumber != m2.SequenceNumber {
		return false, "message sequence numbers don't match"
	}
	if m.AckNumber != m2.AckNumber {
		return false, "message ack numbers don't match"
	}
	if m.Nonce != m2.Nonce {
		return false, "nonces don't match"
	}
	b1, err1 := m.Body.Marshall()
	b2, err2 := m2.Body.Marshall()
	if err1 != nil || err2 != nil {
		return false, fmt.Sprintf("Error: %v, %v", err1, err2)
	}
	if bytes.Compare(b1, b2) != 0 {
		return false, "bodies don't match"
	}

	return true, ""
}

// compareRequestMessages RETURNS IF TWO MESSAGES MATCH
func compareResponseMessages (m *ResponseMessage, m2 *ResponseMessage) (bool, string){
	if m.Type != m2.Type {
		return false, "Message Types don't match"
	}
	if m.SequenceNumber != m2.SequenceNumber {
		return false, "message sequence numbers don't match"
	}
	if m.AckNumber != m2.AckNumber {
		return false, "message ack numbers don't match"
	}
	if m.Nonce != m2.Nonce {
		return false, "nonces don't match"
	}
	b1, err1 := m.Body.Marshall()
	b2, err2 := m2.Body.Marshall()
	if err1 != nil || err2 != nil {
		return false, fmt.Sprintf("Error: %v, %v", err1, err2)
	}
	if bytes.Compare(b1, b2) != 0 {
		return false, "bodies don't match"
	}

	return true, ""
}

// sliceContains WILL DETERMINE IF THE SLICE CONTAINS THE VALUE
func sliceContains(s []interface{}, value interface{}) bool {
	for i := range s {
		if s[i] == value {
			return true
		}
	}
	return false
}