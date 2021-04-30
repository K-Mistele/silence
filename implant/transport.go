package main

import (
	"errors"
	"fmt"
	"github.com/k-mistele/silence/silence/transport"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"math/rand"
	"net"
)

// transportConnection REPRESENTS A TRANSPORT LAYER CONNECTION WITH THE SERVER THAT PROVIDES
// THE INTERFACE TO INTERACT WITH IT AT THE TRANSPORT LEVEL TO HANDLE FLOW AND ERROR CONTROL
type transportConnection struct {
	ListenerAddress    *net.IPAddr
	EndpointID         uint16
	Conn               *icmp.PacketConn
	ErrorState         error
	curSequenceNumber uint32
	curAck             uint32
	packetTimeout      uint8		// SECONDS FOR EACH PACKET TO WAIT W/O ACK BEFORE RESEND

}


func connect(addr *net.IPAddr) (*transportConnection, error) {

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

func (tc *transportConnection) sendFromApplicationLayer(message []byte) (response []byte, err error) {

	// TODO START LISTENER HERE AND PUSH MESSAGES INTO A CHANNEL WHERE THEY CAN BE SORTED AND THEN RETURNED

	// TODO: FRAGMENT AS NECESSARY, START WITH A SEQUENCE NUMBER
	return []byte{}, nil
}

func (tc *transportConnection) sendToICMPLayer(datagram transport.Datagram) error {

	// MARSHALL THE TRANSPORT LAYER DATAGRAM TO BYTES
	transportBytes, err := datagram.Marshall()
	if err != nil {
		return err
	}

	// DOUBLE CHECK IT'S SMALL ENOUGH TO TUNNEL. IT SHOULD BE FINE, BUT DOUBLE CHECK
	if len(transportBytes) > 548 {
		return errors.New(fmt.Sprintf("datagram %+v is too long for icmp encapsulation", datagram))
	}

	// BUILD ECHO REPLY
	icmpMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo {
			ID: 	rand.Int(),
			Seq: 	int(tc.curSequenceNumber % 65535),
			Data: 	transportBytes,
		},
	}


	// MARSHALL THE ECHO REQUEST AND PUT IT ON THE WIRE
	networkBytes, err := icmpMessage.Marshal(nil )
	if err != nil {
		return nil
	}

	_, err = tc.Conn.WriteTo(networkBytes, tc.ListenerAddress)
	if err != nil {
		return err
	}
	return nil
}