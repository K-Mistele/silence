package main

import (
	"encoding/binary"
	"math/rand"
)

var endpointID uint16

func init() {

	// GENERATE A RANDOM ENDPOINT ID TO IDENTIFY THIS LISTENER.
	// THE RANDOM PACKAGE ONLY SUPPORTS UINT32 SO GET THAT, THEN TAKE TWO RANDOM BYTES FROM IT
	random := rand.Uint32()
	randomBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(randomBytes, random)
	endpointID = binary.LittleEndian.Uint16(randomBytes[0:2])
}
