package main

import "github.com/deverdeb/bvmgo-term/term"

func main() {
	width, height, err := term.GetSize()
	if err != nil {
		panic(err)
	}

	term.Clear()

	term.CursorMoveTo(1, 3)
	term.Printf("Window size: column=%d x row=%d", width, height)

	term.CursorMoveTo(1, 1)
	term.Print(term.StyleOk.Sprintf("X"))
	term.CursorMoveTo(width, 1)
	term.Print(term.StyleInfo.Sprintf("X"))
	term.CursorMoveTo(1, height)
	term.Print(term.StyleWarn.Sprintf("X"))
	term.CursorMoveTo(width, height)
	term.Print(term.StyleError.Sprintf("X"))

	term.CursorMoveTo(1, 5)
	term.Print(term.StyleOk.And(term.Style{Bold: true}).Sprintf("Hello"))
	term.CursorMoveTo(3, 6)
	term.Print(term.StyleInfo.And(term.Style{Underline: true}).Sprintf("World"))
	term.CursorMoveTo(5, 7)
	term.Print(term.StyleWarn.And(term.Style{Strikethrough: true}).Sprintf("Hello"))
	term.CursorMoveTo(7, 8)
	term.Print(term.StyleError.And(term.Style{Italic: true}).Sprintf("World"))

	term.CursorMoveTo(1, 9)
}
