package term

import "github.com/deverdeb/bvmgo-term/ansi"

func CursorMoveTo(column, row int) {
	Print(ansi.CursorMoveTo(column, row))
}

func CursorMoveUp(row int) {
	Print(ansi.CursorMoveUp(row))
}

func CursorMoveDown(row int) {
	Print(ansi.CursorMoveDown(row))
}

func CursorMoveRight(column int) {
	Print(ansi.CursorMoveRight(column))
}

func CursorMoveLeft(column int) {
	Print(ansi.CursorMoveLeft(column))
}

func CursorMoveToColumn(column int) {
	Print(ansi.CursorMoveToColumn(column))
}
