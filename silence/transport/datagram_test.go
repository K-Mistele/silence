package transport

import (
	"testing"
)

func TestDatagramEncoding(t *testing.T) {
	var d1, d2 *Datagram

	d1 = &Datagram{
		EndpointID:     1,
		SequenceNumber: 2,
		AckNumber:      3,
		Flags:          struct {
			SYN 			bool // 10000000
			ACK 			bool // 01000000
			FIN 			bool // 00100000
			RST 			bool // 00010000
			IsFragment		bool // 00001000
			Retransmit		bool // 00000100
			Reserved1		bool // 00000010
			Reserved2 		bool // 00000001
		}{
			true,
			true,
			false,
			false,
			false,
			false,
			true,
			false,
		},
		Checksum:       5,
		FragmentNumber: 6,
		TotalFragments: 7,
		Data:           []byte{0x41, 0x42, 0x43, 0x44},
	}

	d2 = &Datagram{}

	icmpPayload, err := d1.Marshall()
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	e := d2.Unmarshall(icmpPayload)
	if e != nil {
		t.Fatalf("%v\n", e)
	}

	if !datagramsAreEqual(d1, d2) {
		t.Fatalf("%+v\n%+v\n", *d1, *d2)
	}
}