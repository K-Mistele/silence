package main

import (
	"github.com/k-mistele/silence/silence/transport"
	"hash/crc32"
	"net"
	"testing"
)



func TestDatagramSend(t *testing.T) {

	var err error
	// SETUP ADDRESS INFORMATION
	listenerAddress, err = net.ResolveIPAddr("ip4", "10.0.2.23")

	tc, err := connect(listenerAddress)
	if err != nil {
		t.Fatalf("%v", err)
	}
	data := []byte("abcdefghijklmnopqrstuvwxyz")

	datagram := transport.Datagram{
		EndpointID:     0x4141,
		SequenceNumber: 0x42424242,
		AckNumber:      0x43434343,
		Flags: transport.DatagramFlags{
			SYN:        true,
			ACK:        false,
			FIN:        false,
			ERR_RST:    false,
			IsFragment: false,
			Retransmit: false,
			Reserved1:  false,
			Reserved2:  false,
		},
		Checksum:       crc32.ChecksumIEEE(data),
		FragmentNumber: 0,
		TotalFragments: 0,
		Data:           data,
	}

	err = tc.sendToICMPLayer(datagram)
	if err != nil {
		t.Fatalf("%v", err)
	}

}
