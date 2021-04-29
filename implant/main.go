package main

import (
	"fmt"
	"os"
)

func main() {
	listenerAddress := os.Args[1]
	transportConnection, err := newTransportConnection(listenerAddress)
	if err != nil {
		fmt.Printf("Couldn't establish transport connection! %v", err)
		os.Exit(1)
	}
}