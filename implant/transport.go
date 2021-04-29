package main

import (
	transport "github.com/k-mistele/silence/silence/transport"
	icmp "golang.org/x/net/icmp"
)

// transportConnection REPRESENTS A TRANSPORT LAYER CONNECTION WITH THE SERVER THAT PROVIDES
// THE INTERFACE TO INTERACT WITH IT AT THE TRANSPORT LEVEL TO HANDLE FLOW AND ERROR CONTROL
type transportConnection struct {
	ListenerAddress    string
	EndpointID         uint16
	Conn               *icmp.PacketConn
	ErrorState         error
	curSequenceNumber uint32
	curAck             uint32
	packetTimeout      uint8		// SECONDS FOR EACH PACKET TO WAIT W/O ACK BEFORE RESEND

}


func connect(listenerAddress string) (*transportConnection, error) {

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil { return nil, err}

	// BUILD A TRANSPORT CONNECTION, AND TRY TO DO THE HANDSHAKE TO SETUP
	tc :=  &transportConnection{
		ListenerAddress: 	listenerAddress,
		EndpointID: 		endpointID,		// GENERATED RANDOMLY AT INIT
		Conn:				conn,
		ErrorState: 		nil,
		curSequenceNumber: 	0,
		curAck: 			0,
		packetTimeout: 		4,				// FOUR SECONDS FOR RIGHT NOW

	}

	return tc, nil
}

func (tc *transportConnection) sendMessage(message []byte) (response []byte, err error) {

	// TODO START LISTENER HERE AND PUSH MESSAGES INTO A CHANNEL WHERE THEY CAN BE SORTED AND THEN RETURNED

	// TODO: FRAGMENT AS NECESSARY, START WITH A SEQUENCE NUMBER
	return []byte{}, nil
}

func (tc *transportConnection) sendPacket(flags transport.DatagramFlags) {

	 tc.curSequenceNumber += 1
}

