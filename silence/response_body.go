package silence

import (
	"errors"
	"fmt"
)

////////////////////////////////////////////////////////////////////
// ResponseMessageBody type
////////////////////////////////////////////////////////////////////

// ResponseMessageBody DEFINES AN INTERFACE FOR DIFFERENT MESSAGE BODY TYPES
type ResponseMessageBody interface{
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

// ResponseBodyNull IS A NULL BODY THAT'S JUST 4 NULL BYTES - USED WHEN ONLY THE CODE IS IMPORTANT
type ResponseBodyNull struct {
	Data []byte
}

// Marshall WILL SERIALIZE THE NULL BODY AND RETURN IT
func (nb *ResponseBodyNull) Marshall() ([]byte, error) {
	return []byte{0, 0, 0, 0}, nil
}

// Unmarshall WILL DESERIALIZE BYTES AND RETURN THE BODY
func (nb *ResponseBodyNull) Unmarshall(b []byte) interface{}{
	nb.Data = []byte{0, 0, 0, 0}
	return nil
}

// ResponseBodyExecuteCommands DEFINES A LIST OF COMMANDS TO EXECUTE
type ResponseBodyExecuteCommands struct {
	Delimiter AsciiCharacter		// FOR EXTENSIBILITY IF USING UNIT SEPARATOR BECOMES A PROBLEM
	Commands  []string
}

// Marshall WILL SERIALIZE THE ResponseBodyExecuteCommands TO BYTES
func (rec *ResponseBodyExecuteCommands) Marshall() ([]byte, error) {

	var data []byte

	// MAKE SURE WE'RE NOT GOING OVER THE TOTAL LENGTH
	totalCommandLength := 1 // 1 FOR THE DELIMITER BYTE AT THE FRONT
	for i := range rec.Commands {
		totalCommandLength += len(rec.Commands[i]) + 1 // + 1 for the DELIMITER
	}

	if totalCommandLength > maxBodyLength {
		return data, errors.New("total command length is too long")
	}

	//  BUILD OUT THE SERIALIZATION
	data = append(data, uint8(rec.Delimiter)) // ADD THE DELIMETER
	for i := range rec.Commands {

		// ADD THE COMMAND AND DELIMITER AFTER
		data = append(data, []byte(rec.Commands[i])...)
		data = append(data, uint8(rec.Delimiter))
	}

	return data, nil


}

// Unmarshall WILL DESERALIZE BYTES TO A ResponseBodyExecuteCommands
func (rec *ResponseBodyExecuteCommands) Unmarshall(b []byte) (err interface{}) {

	err = nil
	// CATCH A PANIC IF WE FAIL TO DECODE PROPERLY
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic while unmarshalling: ", r)
			r = nil
			err = r
		}
	}()

	// PULL OUT THE DELIMITER FROM THE FIRST BYTE
	rec.Delimiter = AsciiCharacter(b[0])
	delimiter := uint8(rec.Delimiter)

	var commands []string
	var curCommand []byte
	for i := 1; i < len(b); i++ {

		// IF WE FIND A DELIMiTER, SAVE THE STRING OFF TO THE LIST OF COMMANDS
		if b[i] == delimiter {
			commands = append(commands, string(curCommand))
			curCommand = []byte{}
		} else {

			// OTHERWISE, ADD THE BYTE TO THE CURRENT COMMAND
			curCommand = append(curCommand, b[i])
		}
	}

	rec.Commands = commands
	return nil
}

// NewResponseBodyExecuteCommands WILL BUILD A NEW ResponseBodyExecuteCommands STRUCT
func NewResponseBodyExecuteCommands(commands []string) *ResponseBodyExecuteCommands {

	delimiter := defaultDelimiter
	// TODO - SWITCH DELIMITER TO A DIFFERENT ONE IF ITS' IN THE MESSAGE, BUT WILL ASSUME FOR NOT THAT ITS NOT

	body := &ResponseBodyExecuteCommands{
		Delimiter: 		delimiter,
		Commands: 		commands,
	}
	return body
}


