package main

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/k-mistele/silence/silence/transport"

)

// process EACH PACKET AS IT HITS THE WIRE
func process(packet gopacket.Packet) error {

	// GET THE IP AND ICMP LAYERS
	var ip4, icmp gopacket.Layer
	var ip4Layer *layers.IPv4
	var icmpLayer *layers.ICMPv4

	ip4 = packet.Layer(layers.LayerTypeIPv4)
	if ip4 == nil { return errors.New("packet did not contain an IPv4 layer")}

	icmp = packet.Layer(layers.LayerTypeICMPv4)
	if icmp == nil { return errors.New("packet did not contain an ICMP layer")}

	// IF THE PACKET IS NOT AN ECHO REQUEST WE SHOULD DISCARD IT
	// GRAB THE PAYLOAD, GOPACKET CONSIDERS THIS APPLICATION LAYER. THIS IS ALSO AVAILABLE IN THE ICMP LAYER.
	payload := packet.ApplicationLayer().(*gopacket.Payload)
	ip4Layer, _ = ip4.(*layers.IPv4)
	icmpLayer, _ = icmp.(*layers.ICMPv4)

	// MAKE SURE WE'RE ONLY READING ECHO REQUESTS TO OUR INTERFACE
	if ip4Layer.DstIP.String() != interfaceAddress.String() || icmpLayer.TypeCode.Type() != layers.ICMPv4TypeEchoRequest {
		return nil
	}

	fmt.Println("Packet:\n----------------------------------")
	fmt.Printf("IPv4: %+v\n", ip4Layer)
	fmt.Printf("ICMP: %+v\n", icmpLayer)
	fmt.Println("Payload", payload.LayerContents(), len(payload.LayerContents()))

	// DECODE THE DATAGRAM
	var datagram transport.Datagram
	err := datagram.Unmarshall(payload.LayerContents())
	if err != nil {
		fmt.Printf("Error decoding datagram: %v\n", err)
		return errors.New(fmt.Sprintf("%v", err))
	}
	fmt.Printf("Decoded datagram: %+v", datagram)

	return nil
}