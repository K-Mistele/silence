package silence

import (
	"fmt"
)


func main() {
	fmt.Println("Hello, world")

	silenceMessage := RequestMessage{
		Type:  ReadyForCommand,
		SequenceNumber: 1,
		Nonce: 0x41424344,
		Body: NullBody{},
	}

	m := silenceMessage.Marshall()
	fmt.Println(m)


}
