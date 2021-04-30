package main

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"net"
)

var interfaceAddress net.IP

func getLocalAddr() (net.IP, string, error){

	var interfaceName string
	var err error
	// COMMAND LINE ARGUMENT PARSING
	flag.StringVar(&interfaceName, "iface", "eth0", "The name of the network interface to listen on")
	flag.Parse()

	// GET INTERFACE IP ADDRESS
	fmt.Printf("Capturing packets on interface %s\n", interfaceName)
	iFace, err := net.InterfaceByName(interfaceName)
	if err != nil {panic (err)}
	addrs, err := iFace.Addrs()
	if err != nil {return interfaceAddress, interfaceName, nil }

	for _, addr := range addrs {
		interfaceAddress = addr.(*net.IPNet).IP.To4()
		break
	}
	fmt.Printf("Interface IP address is %s\n", interfaceAddress.String())
	return interfaceAddress, interfaceName, err
}

func main() {

	_, interfaceName, err := getLocalAddr()
	if err != nil {
		panic(err)
	}

	// CREATE THE PACKET CAPTURE
	handle, err := pcap.OpenLive(interfaceName, 3600, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	// FILTER FOR ICMP ECHO REQUESTS WITH DESTINATION HOST OF ME, THE LISTENER. THESE ARE THE ONLY ONES WE CARE ABOUT.
	filter := fmt.Sprintf("icmp and icmp[icmptype] == icmp-echo and dst host %s", interfaceAddress.String())
	err = handle.SetBPFFilter(filter)
	if err != nil {
		fmt.Printf("Failed to set packet filter: %v; filter string: %s\n", err, filter)
	} else {
		fmt.Printf("Applied filter '%s'\n", filter)
	}

	fmt.Println("Ready to start processing packets!")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// PROCESS PACKETS
	for packet := range packetSource.Packets() {

		process(packet)

	}

}
