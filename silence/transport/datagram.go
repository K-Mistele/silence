package transport

import (
	"encoding/binary"
	"errors"
	"fmt"
)
// Datagram IS THE TRANSPORT-LAYER DATAGRAM OF THE SILENCE PROTOCOL
// SINCE IT'S ENCAPSULATED IN ICMP, IT MAY BE NO MORE THAN 548 BYTES MEANING THAT DATA MAY BE ONLY 525 BYTES
type Datagram struct {
	EndpointID		uint16		// 2 BYTES
	SequenceNumber	uint32		// 4 BYTES
	AckNumber		uint32		// 4 BYTES
	Flags 			uint8		// 1 BYTE
	Checksum		uint32		// 4 BYTES
	FragmentNumber	uint32 		// 4 BYTES
	Reserved		uint32		// 4 BYTES
	Data 			[]byte		// 525 bytes
}

// Marshall WILL CONVERT THE DATAGRAM TO A BYTE SLICE
func (d *Datagram) Marshall() ([]byte, error){

	var data []byte
	if len(d.Data) > 525 {
		return data, errors.New("Error: data is too long!")
	}

	// ADD THE ENDPOINT ID
	endpointIDBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(endpointIDBytes, d.EndpointID)
	data = append(data, endpointIDBytes...)

	// ADD THE SEQUENCE NUMBER
	seqNoBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqNoBytes, d.SequenceNumber)
	data = append(data, seqNoBytes...)

	// ADD THE ACK NUMBER
	ackNoBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(ackNoBytes, d.AckNumber)
	data = append(data, ackNoBytes...)

	// ADD THE FLAGS
	data = append(data, d.Flags)

	// ADD THE CHECKSUM
	checkSumBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(checkSumBytes, d.Checksum)
	data = append(data, checkSumBytes...)

	// ADD THE FRAGMENT NUMBER
	fragNoBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(fragNoBytes, d.FragmentNumber)
	data = append(data, fragNoBytes...)

	// ADD RESERVED BYTES
	reservedBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(reservedBytes, d.Reserved)
	data = append(data, reservedBytes...)

	// ADD DATA
	data = append(data, d.Data...)

	return data, nil
}

// Unmarshall WILL CONVERT A BYTE SLICE INTO A Datagram
func (d *Datagram) Unmarshall(data []byte) (err interface{}) {

	err = nil
	// IF THE MESSAGE IS MALFORMED, RECOVER AND DON'T DO ANYTHING
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered from panic while unmarshalling: ", p)
			d = nil
			err = p
		}
	}()

	d.EndpointID = binary.LittleEndian.Uint16(data[0:2])
	d.SequenceNumber = binary.LittleEndian.Uint32(data[2:6])
	d.AckNumber = binary.LittleEndian.Uint32(data[6:10])
	d.Flags = data[10]
	d.Checksum = binary.LittleEndian.Uint32(data[11:15])
	d.FragmentNumber = binary.LittleEndian.Uint32(data[15:19])
	d.Reserved = binary.LittleEndian.Uint32(data[19:23])
	d.Data = data[23:]

	return nil
}
