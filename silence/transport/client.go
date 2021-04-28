package transport

import "net"

// TransportClient IS FOR THE IMPLANT
type TransportClient struct {
	EndpointID uint16
	PacketConn *net.PacketConn
}

// ClientConnect ESTABLISHES A CLIENT CONNECTION TO THE SERVER, WITH A VARIBALE PING RATE AND JITTER
func ClientConnect(serverIP string, pingRate int, jitter int)  (*TransportClient, error) {

}
