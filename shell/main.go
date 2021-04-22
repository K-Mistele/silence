package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"time"
)

const (
	protoICMP = 1
	listenAddress = "0.0.0.0"
	networkICMP = "ip4:icmp"
	networkIPv4 = "ip4"
)
func main() {

	// COMMAND LINE ARGUMENTS
	// COMMAND LINE ARGUMENT PARSING
	var listenerIP4 string
	var listenerAddr net.Addr
	var err error

	flag.StringVar(&listenerIP4, "listener", "", "The IPv4 address or hostname of the listener server")
	flag.Parse()

	if listenerIP4 == "" {
		panic(errors.New("you must specify the IPv4 address of the listening server"))
	} else {
		listenerAddr, err = net.ResolveIPAddr(networkIPv4, listenAddress)
		if err != nil { panic(err)}
	}

	// SET UP A CONNECTION
	packetConn, err := icmp.ListenPacket(networkICMP, listenAddress)
	if err != nil { panic(err) }
	defer packetConn.Close()

	// BUILD THE PACKET
	icmpMessage := icmp.Message {
		Type: ipv4.ICMPTypeEcho,
		Body: &icmp.Echo{
			ID: 0x4242,
			Seq: 255,
			Data: []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48}, // ABCDEFGH
		},
	}

	// MARSHALL AND SEND THE MESSAGE
	messageBytes, err := icmpMessage.Marshal(nil)
	if err != nil { panic(err) }
	length, err := packetConn.WriteTo(messageBytes, listenerAddr)
	if err != nil { panic(err) }

	fmt.Printf("length: %d\n", length)

	// BUILD A REPLY
	reply := make([]byte, 1500)
	err = packetConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil { panic(err) }
	length, peer, err := packetConn.ReadFrom(reply)
	if err != nil { panic(err) }
	fmt.Printf("Received reply from %s:%s\n", peer.Network(), peer.String())

	// PARSE THE REPLY
	replyMessage, err := icmp.ParseMessage(protoICMP, reply[:length])
	if err != nil { panic(err)}
	switch replyMessage.Code {
	case int(ipv4.ICMPTypeEchoReply):
		fmt.Println("Got message reply!")
	case int(ipv4.ICMPTypeEcho):
		fmt.Println("Got Message reply (2)!")
	case int(ipv4.ICMPTypeExtendedEchoReply):
		fmt.Println("Got Message reply (3)!")
	default:
		fmt.Printf("Got message reply %v but wanted %v (%d, %d)\n", replyMessage.Type, ipv4.ICMPTypeEchoReply, replyMessage.Type, ipv4.ICMPTypeEchoReply)
	}
	fmt.Printf("Got message code %v but wanted %v (%d, %d)\n", replyMessage.Code, ipv4.ICMPTypeEchoReply, replyMessage.Code, ipv4.ICMPTypeEchoReply)


}