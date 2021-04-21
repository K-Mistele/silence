package main

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func process(packet gopacket.Packet) error {

	var ip4, icmp gopacket.Layer
	var ip4Layer *layers.IPv4
	var icmpLayer *layers.ICMPv4
	ip4 = packet.Layer(layers.LayerTypeIPv4)
	if ip4 == nil { return errors.New("packet did not contain an IPv4 layer")}

	icmp = packet.Layer(layers.LayerTypeICMPv4)
	if icmp == nil { return errors.New("packet did not contain an ICMP layer")}

	payload := packet.ApplicationLayer().(*gopacket.Payload)
	ip4Layer, _ = ip4.(*layers.IPv4)
	icmpLayer, _ = icmp.(*layers.ICMPv4)

	fmt.Println("Packet:\n----------------------------------")
	fmt.Println("IPv4 ", ip4Layer)
	fmt.Println("ICMP ", icmpLayer)
	fmt.Println("Payload", payload.LayerContents(), len(payload.LayerContents()))

	return nil
}