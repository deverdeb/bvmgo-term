package ansi

// TextSplitSize method returns lines of string. Lines length is maxWidth (maximum).
// Ignore and don't cut Ansi sequences.
// If possible, cut line at space character.
func TextSplitToSize(str string, maxWidth int) []string {
	return (&textSplitter{}).Process(str, maxWidth)
}

type textSplitItem struct {
	data   string
	length int
}

func (item *textSplitItem) AddString(str string, count bool) {
	for _, character := range str {
		item.AddRune(character, count)
	}
}

func (item *textSplitItem) AddRune(character rune, count bool) {
	item.data += string(character)
	if count {
		item.length++
	}
}

func (item *textSplitItem) AddItem(otherItem textSplitItem) {
	item.data += otherItem.data
	item.length += otherItem.length
}

func (item *textSplitItem) Reset() {
	item.data = ""
	item.length = 0
}

type textSplitter struct {
	lines         []string
	line          textSplitItem
	word          textSplitItem
	inEscSequence bool
	maxWidth      int
	lastSpaceChar string
}

func (splitter *textSplitter) Process(text string, maxWidth int) []string {
	splitter.lines = make([]string, 0)
	splitter.line = textSplitItem{}
	splitter.word = textSplitItem{}
	splitter.inEscSequence = false
	splitter.maxWidth = maxWidth
	splitter.lastSpaceChar = ""
	for _, character := range text {
		splitter.AddRune(character)
	}
	splitter.lines = append(splitter.lines, splitter.line.data+splitter.lastSpaceChar+splitter.word.data)
	return splitter.lines
}

func (splitter *textSplitter) AddRune(character rune) {
	if splitter.inEscSequence {
		// Start ANSI sequence
		splitter.word.AddRune(character, !splitter.inEscSequence)
		if isLetter(character) {
			// Stop ANSI sequence
			splitter.inEscSequence = false
		}
	} else if character == Esc {
		// Start ANSI sequence
		splitter.inEscSequence = true
		splitter.word.AddRune(character, !splitter.inEscSequence)
	} else if character == '\n' {
		splitter.processEndLine()
	} else if character == ' ' || character == '\t' {
		splitter.processSpace(character)
	} else {
		splitter.processCharacter(character)
	}
}

func (splitter *textSplitter) processEndLine() {
	// End line
	splitter.lines = append(splitter.lines, splitter.line.data+splitter.lastSpaceChar+splitter.word.data)
	splitter.line.Reset()
	splitter.word.Reset()
	splitter.lastSpaceChar = ""
}

func (splitter *textSplitter) processSpace(character rune) {
	// Space - End word
	splitter.line.AddString(splitter.lastSpaceChar+splitter.word.data, !splitter.inEscSequence)
	splitter.word.Reset()
	splitter.lastSpaceChar = string(character)
}

func (splitter *textSplitter) processCharacter(character rune) {
	splitter.word.AddRune(character, !splitter.inEscSequence)
	if splitter.maxWidth > 0 && splitter.line.length+len(splitter.lastSpaceChar)+splitter.word.length >= splitter.maxWidth {
		// Change line
		if splitter.line.length == 0 {
			// Too long word: cut word
			splitter.lines = append(splitter.lines, splitter.line.data+splitter.lastSpaceChar+splitter.word.data)
			splitter.line.Reset()
			splitter.word.Reset()
			splitter.lastSpaceChar = ""
		} else {
			splitter.lines = append(splitter.lines, splitter.line.data)
			splitter.line.Reset()
			splitter.lastSpaceChar = ""
		}
	}
}
