package silence

////////////////////////////////////////////////////////////////////
// RequestMessageBody type
////////////////////////////////////////////////////////////////////

// RequestMessageBody DEFINES AN INTERFACE FOR DIFFERENT REQUEST MESSAGE BODY TYPES
type RequestMessageBody interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) interface{}
}

////////////////////////////////////////////////////////////////////
// RequestBodyBull type
////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////
// RequestBodyReadyForCommand type
////////////////////////////////////////////////////////////////////
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

// NewRequestBodyReadyForCommand WILL RETURN A NEW ONE
func NewRequestBodyReadyForCommand() *RequestBodyReadyForCommand {
	return &RequestBodyReadyForCommand{}
}

////////////////////////////////////////////////////////////////////
// RequestBodyPrintCommandData type
////////////////////////////////////////////////////////////////////
type RequestBodyPrintCommandData struct {
	CommandData 		string
}

// Marshall WILL SERIALIZE IT TO BYTES
func (pcd *RequestBodyPrintCommandData) Marshall() ([]byte, error) {
	return []byte(pcd.CommandData), nil
}

// Unmarshall WILL DESERIALIZE IT TO THE STRUCT
func (pcd *RequestBodyPrintCommandData) Unmarshall(b []byte) interface{} {
	pcd.CommandData = string(b)
	return nil
}

// NewRequestBodyPrintCommandData WILL BUILD A NEW RequestBodyPrintCommandData MESSAGE WITH THE SPECIFIED DATA
func NewRequestBodyPrintCommandData(s string) *RequestBodyPrintCommandData {
	return &RequestBodyPrintCommandData{
		CommandData: s,
	}
}