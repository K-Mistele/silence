package main

import (
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"net"
	"flag"
	"fmt"
)


func main() {

	// COMMAND LINE ARGUMENT PARSING
	var iFaceStr string

	flag.StringVar(&iFaceStr, "interface", "eth0", "The name of the network interface to listen on")
	flag.Parse()

	// GET INTERFACE IP ADDRESS
	fmt.Printf("Capturing packets on interface %s\n", iFaceStr)
	iFace, err := net.InterfaceByName(iFaceStr)
	if err != nil {panic (err)}
	addrs, err := iFace.Addrs()
	if err != nil {panic (err)}

	var interfaceAddress net.IP
	for _, addr := range addrs {
		interfaceAddress = addr.(*net.IPNet).IP.To4()
		break
	}
	fmt.Printf("Interface IP address is %s\n", interfaceAddress.String())

	// CREATE THE PACKET CAPTURE
	handle, err := pcap.OpenLive(iFaceStr, 3600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	filter := fmt.Sprintf("icmp and ip.dst == %s", interfaceAddress)
	handle.SetBPFFilter(filter)

	// FILTER FOR ICMP MESSAGES DESTINED TO MY INTERFACE
	//filterString := fmt.Sprintf("icmp")
	//fmt.Println(filterString)
	//if err = handle.SetBPFFilter(filterString); err != nil {panic (err)}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// PROCESS PACKETS
	for packet := range packetSource.Packets() {

		process(packet)

	}

}
