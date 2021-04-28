package transport

import (
	"testing"
)

func TestDatagramEncoding(t *testing.T) {
	var d1, d2 *Datagram

	d1 = &Datagram{
		EndpointID: 1,
		SequenceNumber: 2,
		AckNumber: 3,
		Flags: 4,
		Checksum: 5,
		FragmentNumber: 6,
		Reserved: 7,
		Data: []byte{0x41, 0x42, 0x43, 0x44},
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