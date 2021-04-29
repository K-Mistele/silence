package main

import (
	"net"
)

type transportConnection struct {
	ListenerIPv4 		string
	EndpointID			uint16
	Conn 				*net.PacketConn
	ErrorState			error
	currSequenceNumber	uint32
	curAck				uint32
	packetTimeout		uint8		// SECONDS FOR EACH PACKET TO WAIT W/O ACK BEFORE RESEND

}

func connect(listenerIPv4 string) *transportConnection {

	return &transportConnection{}
}

func (tc *transportConnection) sendMessage(message []byte) (response []byte, err error) {

	// TODO START LISTENER HERE AND PUSH MESSAGES INTO A CHANNEL WHERE THEY CAN BE SORTED AND THEN RETURNED

	// TODO: FRAGMENT AS NECESSARY, START WITH A SEQUENCE NUMBER
	return []byte{}, nil
}

func (tc *transportConnection) sendPacket(/* TODO ALL PACKET DATA HERE*/) {

}

