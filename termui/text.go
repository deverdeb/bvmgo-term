package termui

import (
	"github.com/deverdeb/bvmgo-term/ansi"
	"github.com/deverdeb/bvmgo-term/term"
	"strings"
)

type Text struct {
	content string
	lines   []string

	maxWidth  int
	maxHeight int
}

func BuildText(textContent string) Text {
	uiText := Text{
		content:   textContent,
		lines:     nil,
		maxWidth:  0,
		maxHeight: 0,
	}
	uiText.computeLines()
	return uiText
}

func (text *Text) SetText(content string) {
	text.content = content
	text.computeLines()
}

func (text *Text) SetMaxWidth(maxWidth int) {
	text.maxWidth = maxWidth
	text.computeLines()
}

func (text *Text) SetMaxHeight(maxHeight int) {
	text.maxHeight = maxHeight
}

func (text *Text) Display() {
	cursor := term.Cursor{}
	for row, line := range text.lines {
		cursor.MoveTo(0, row)
		cursor.Print(line)
	}
	cursor.MoveTo(0, 0)
}

func (text *Text) Hide() {
	cursor := term.Cursor{}
	width, height := text.Dimension()
	line := strings.Repeat(" ", width)
	for posRow := 0; posRow < height; posRow++ {
		cursor.MoveTo(0, posRow)
		cursor.Print(line)
	}
	cursor.MoveTo(0, 0)
}

func (text *Text) Dimension() (width, height int) {
	height = len(text.lines)
	width = 0
	for _, line := range text.lines {
		width = max(width, ansi.StringLen(line))
	}
	return width, height
}

func (text *Text) computeLines() {
	cleanText := strings.Replace(text.content, "\t", "  ", -1)
	text.lines = ansi.TextSplitToSize(cleanText, text.maxWidth)
}
