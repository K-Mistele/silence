package silence

// ASCII CHARACTER MAPPIGNS
type AsciiCharacter uint8
const (
	escape			AsciiCharacter = 0x1B
	fileSeparator	AsciiCharacter = 0x1C
	groupSeparator	AsciiCharacter = 0x1D
	recordSeparator	AsciiCharacter = 0x1E
	unitSeperator 	AsciiCharacter = 0x1F
	newLine       	AsciiCharacter = 0x0A

)

const maxBodyLength int = 576

const defaultDelimiter = unitSeperator
var alternativeDelimiters = []AsciiCharacter{recordSeparator, groupSeparator, fileSeparator}

// SET WHETHER TO ENCODE MESSAGES
var encodeMesages = true

// encodeMessages SETS WHETHER MESSAGES SHOULD BE ENCODED ON BEING BUILT. TRUE BY DEFAULT
func encodeMessages (b bool) {
	encodeMesages = b
}