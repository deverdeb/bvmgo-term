package ansi

import (
	"fmt"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

const (
	// Esc is the ANSI escape character.
	Esc = "\u001b" //"\033";
	// Csi is the Control Sequence Introducer
	Csi = Esc + "[" //"\033[";
)

func ClearScreen() string {
	return Csi + "2J"
}

func CursorAskPosition() string {
	return Csi + "6n"
}

func CursorMoveHome() string {
	return Csi + "H"
}

func CursorMoveTo(column, row int) string {
	return fmt.Sprintf(Csi+"%d;%dH", row, column)
}

func CursorMoveUp(row int) string {
	return fmt.Sprintf(Csi+"%dA", row)
}

func CursorMoveDown(row int) string {
	return fmt.Sprintf(Csi+"%dB", row)
}

func CursorMoveRight(column int) string {
	return fmt.Sprintf(Csi+"%dC", column)
}

func CursorMoveLeft(column int) string {
	return fmt.Sprintf(Csi+"%dD", column)
}

func CursorMoveToColumn(column int) string {
	return fmt.Sprintf(Csi+"%dG", column)
}

func StyleReset() string {
	return Csi + "0m"
}

func StyleBold() string {
	return Csi + "1m"
}

func StyleItalic() string {
	return Csi + "3m"
}

func StyleUnderline() string {
	return Csi + "4m"
}

func StyleStrikethrough() string {
	return Csi + "9m"
}

func StyleForegroundColor(red, green, blue uint8) string {
	return fmt.Sprintf(Csi+"38;2;%d;%d;%dm", red, green, blue)
}

func StyleBackgroundColor(red, green, blue uint8) string {
	return fmt.Sprintf(Csi+"48;2;%d;%d;%dm", red, green, blue)
}
