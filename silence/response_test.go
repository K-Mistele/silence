package silence


import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

// TEST THE ResponseBodyExecuteCommands MESSAGE
func TestResponseBodyExecuteCommands(t *testing.T) {

	commands := []string{"id", "whoami", "ls -al"}

	nonce := rand.Uint32()
	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, nonce)

	body := NewResponseBodyExecuteCommands(commands)
	message := ResponseMessage{
	Code:           ResponseMessageTypeExecuteCommands,
	SequenceNumber: 0,
	AckNumber:      0,
	Nonce:          nonce,
	Body:           body,

	}

	b, err := message.Marshall()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}

	message2 := ResponseMessage{}
	message2.Unmarshall(b)
	messagesMatch, errString := compareResponseMessages(&message, &message2)
	if !messagesMatch {
		fmt.Printf("%+v\n%+v\n", message.Body, message2.Body)
		t.Fatalf(errString)
	}


}

// TEST THE ResponseBodyNoop MESSAGE
func TestResponseBodyNoop(t *testing.T) {

	body := NewResponseBodyNoop()
	message := NewResponseMessage(ResponseMessageTypeNoop, 0, 0, body)

	b, err := message.Marshall()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}

	message2 := ResponseMessage{}
	message2.Unmarshall(b)
	messagesMatch, errString := compareResponseMessages(message, &message2)
	if !messagesMatch {
		t.Fatalf(errString)
	}
}