package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var err error
	// PARSE ARGS
	// GET LISTENER ADDRESS FROM ARGS, IF NOT THEN USE A HARD CODED ONE
	listenerAddrString := flag.String("listener", "1.1.1.1", "Listener IPv4 Address")
	flag.Parse()

	fmt.Printf("Address string: %s\n", *listenerAddrString)
	listenerAddress, err = net.ResolveIPAddr("ip4", *listenerAddrString)
	if err != nil {
		fmt.Printf("Error: couldn't resolve address %s", os.Args[1])
	}

	_, err = connect(listenerAddress)
	if err != nil {
		fmt.Printf("Couldn't establish transport connection! %v", err)
		os.Exit(1)
	}

	fmt.Println("Connection established!")

}