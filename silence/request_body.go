package silence

////////////////////////////////////////////////////////////////////
// RequestMessageBody type
////////////////////////////////////////////////////////////////////

// RequestMessageBody DEFINES AN INTERFACE FOR DIFFERENT REQUEST MESSAGE BODY TYPES
type RequestMessageBody interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

// RequestBodyNull IS A NULL BODY THAT'S JUST 4 NULL BYTES - USED WHEN ONLY THE CODE IS IMPORTANT
type RequestBodyNull struct {
	Data []byte
}

// Marshall WILL SERIALIZE THE NULL BODY AND RETURN IT
func (nb *RequestBodyNull) Marshall() ([]byte, error) {
	return []byte{0, 0, 0, 0}, nil
}

// Unmarshall WILL UPDATE THE POINTER WITH THE PROPERTIES FROM THE BYTES
func (nb *RequestBodyNull) Unmarshall(b []byte) interface{}{
	nb.Data = []byte{0, 0, 0, 0}
	return nil
}

// RequestBodyReadyForCommand IS EMPTY SINCE NO DATA IS NEEDED
type RequestBodyReadyForCommand struct{}

// Marshall WILL SERIALIZE IT TO NOTHING :)
func (rcb *RequestBodyReadyForCommand) Marshall() ([]byte, error) {
	return []byte{}, nil
}

// Unmarshall WILL DESERIALIZE IT TO NOTHING :)
func (rcb *RequestBodyReadyForCommand) Unmarshall(b []byte) interface{} {
	return nil
}
