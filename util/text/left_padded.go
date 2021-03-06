package text

import (
	"bytes"
	"fmt"
	"strings"
)

// LeftPadMaxLine pads a string on the left by a specified amount and pads the string on the right to fill the maxLength
func LeftPadMaxLine(text string, length, leftPad int) string {
	runes := []rune(text)

	// truncate and ellipse if needed
	if len(runes)+leftPad > length {
		runes = append(runes[:(length-leftPad-1)], '…')
	}

	if len(runes)+leftPad < length {
		runes = append(runes, []rune(strings.Repeat(" ", length-len(runes)-leftPad))...)
	}

	return fmt.Sprintf("%s%s",
		strings.Repeat(" ", leftPad),
		string(runes),
	)
}

// LeftPad left pad each line of the given text
func LeftPad(text string, leftPad int) string {
	var result bytes.Buffer

	pad := strings.Repeat(" ", leftPad)

	for _, line := range strings.Split(text, "\n") {
		result.WriteString(pad)
		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}
