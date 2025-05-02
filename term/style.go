package term

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/ansi"
	"strings"
)

var (
	StyleOk    = Style{Foreground: &ColorOk}
	StyleInfo  = Style{Foreground: &ColorInfo}
	StyleWarn  = Style{Foreground: &ColorWarn}
	StyleError = Style{Foreground: &ColorError}
)

type Style struct {
	Foreground    *Color
	Background    *Color
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Uppercase     bool
}

func (style Style) Sprintf(text string, arguments ...any) string {
	return style.Sprint(fmt.Sprintf(text, arguments...))
}

func (style Style) Sprint(text ...any) string {
	formatedText := fmt.Sprint(text...)
	if len(text) == 0 {
		return formatedText
	}
	if style.Uppercase {
		formatedText = strings.ToUpper(formatedText)
	}
	return style.Begin() + formatedText + style.End()
}

func (style Style) Begin() string {
	styleString := ""
	if style.Bold {
		styleString += ansi.StyleBold()
	}
	if style.Italic {
		styleString += ansi.StyleItalic()
	}
	if style.Underline {
		styleString += ansi.StyleUnderline()
	}
	if style.Strikethrough {
		styleString += ansi.StyleStrikethrough()
	}
	if style.Foreground != nil {
		styleString += ansi.StyleForegroundColor(style.Foreground.Red(), style.Foreground.Green(), style.Foreground.Blue())
	}
	if style.Background != nil {
		styleString += ansi.StyleBackgroundColor(style.Background.Red(), style.Background.Green(), style.Background.Blue())
	}
	return styleString
}

func (style Style) End() string {
	return ansi.StyleReset()
}

func (style Style) And(otherStyle Style) Style {
	return Style{
		Foreground:    ColorAverage(style.Foreground, otherStyle.Foreground),
		Background:    ColorAverage(style.Background, otherStyle.Background),
		Bold:          style.Bold || otherStyle.Bold,
		Italic:        style.Italic || otherStyle.Italic,
		Underline:     style.Underline || otherStyle.Underline,
		Strikethrough: style.Strikethrough || otherStyle.Strikethrough,
		Uppercase:     style.Uppercase || otherStyle.Uppercase,
	}
}
