package term

import "github.com/deverdeb/bvmgo-term/ansi"

func CursorMoveTo(column, row int) {
	Printf(ansi.CursorMoveTo(column, row))
}

func CursorMoveUp(row int) {
	Printf(ansi.CursorMoveUp(row))
}

func CursorMoveDown(row int) {
	Printf(ansi.CursorMoveDown(row))
}

func CursorMoveRight(column int) {
	Printf(ansi.CursorMoveRight(column))
}

func CursorMoveLeft(column int) {
	Printf(ansi.CursorMoveLeft(column))
}

func CursorMoveToColumn(column int) {
	Printf(ansi.CursorMoveToColumn(column))
}
