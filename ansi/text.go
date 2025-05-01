package ansi

// StringLen method returns current size without ANSI escape sequence
func StringLen(str string) int {
	length := 0
	inEscSequence := false
	for _, character := range str {
		if inEscSequence {
			if isLetter(character) {
				inEscSequence = false
			}
		} else if character == Esc {
			inEscSequence = true
		} else {
			length++
		}
	}
	return length
}

// StringRemoveAnsi method returns string without ANSI sequences.
func StringRemoveAnsi(str string) string {
	cleanStr := ""
	inEscSequence := false
	for _, character := range str {
		if inEscSequence {
			if isLetter(character) {
				inEscSequence = false
			}
		} else if character == Esc {
			inEscSequence = true
		} else {
			cleanStr += string(character)
		}
	}
	return cleanStr
}

func isLetter(character rune) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
}
