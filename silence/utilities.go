package silence

import (
	"encoding/binary"
)

// xorEncode WILL ENCODE THE BYTES IN A SLICE BY XORING THEM WITH THE KEY
func xorEncode(bytes *[]byte, key uint32) []byte {

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
