package term

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/ansi"
	"strings"
)

// Cursor is used to move cursor relatively current cursor position (when cursor is created or reset).
// Warning: use only cursor methods to move cursor or print string, otherwise cursor position can be false.
type Cursor struct {
	currentColumn, currentRow int
	maxColumn, maxRow         int
	minColumn, minRow         int
}

func (cursor *Cursor) setColumn(column int) {
	cursor.currentColumn = column
	cursor.maxColumn = max(column, cursor.maxColumn)
	cursor.minColumn = min(column, cursor.minColumn)
}

func (cursor *Cursor) setRow(row int) {
	cursor.currentRow = row
	cursor.maxRow = max(row, cursor.maxRow)
	cursor.minRow = min(row, cursor.minRow)
}

func (cursor *Cursor) MoveTo(column, row int) {
	cursor.MoveToColumn(column)
	cursor.MoveToRow(row)
}

func (cursor *Cursor) MoveUp(row int) {
	Print(ansi.CursorMoveUp(row))
	cursor.setRow(cursor.currentRow - row)
}

func (cursor *Cursor) MoveDown(row int) {
	Print(ansi.CursorMoveDown(row))
	cursor.setRow(cursor.currentRow + row)
}

func (cursor *Cursor) MoveRight(column int) {
	Print(ansi.CursorMoveRight(column))
	cursor.setColumn(cursor.currentColumn + column)
}

func (cursor *Cursor) MoveLeft(column int) {
	Print(ansi.CursorMoveLeft(column))
	cursor.setColumn(cursor.currentColumn - column)
}

func (cursor *Cursor) MoveToColumn(column int) {
	move := column - cursor.currentColumn
	if move < 0 {
		cursor.MoveLeft(-move)
	} else if move > 0 {
		cursor.MoveRight(move)
	}
}

func (cursor *Cursor) MoveToRow(row int) {
	move := row - cursor.currentRow
	if move < 0 {
		cursor.MoveUp(-move)
	} else if move > 0 {
		cursor.MoveDown(move)
	}
}

func (cursor *Cursor) MoveToNextLine() {
	cursor.MoveDown(1)
	cursor.MoveLeft(cursor.currentColumn)
}

func (cursor *Cursor) Printf(format string, params ...any) {
	cursor.Print(fmt.Sprintf(format, params...))
}

func (cursor *Cursor) Print(text ...any) {
	formatedText := fmt.Sprint(text...)
	formatedText = strings.Replace(formatedText, "\t", "  ", -1)
	lines := strings.Split(formatedText, "\n")
	for idx, line := range lines {
		Print(line)
		lineLength := ansi.StringLen(line)
		cursor.setColumn(cursor.currentColumn + lineLength)
		if idx < len(lines)-1 {
			// Not last line - move to next line
			cursor.MoveDown(1)
			cursor.MoveLeft(lineLength)
		}
	}
}

func (cursor *Cursor) MaxPosition() (maxColumn, maxRow int) {
	return cursor.maxColumn, cursor.maxRow
}

func (cursor *Cursor) MinPosition() (minColumn, minRow int) {
	return cursor.minColumn, cursor.minRow
}
