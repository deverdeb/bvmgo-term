package termui

import (
	"github.com/deverdeb/bvmgo-term/term"
	"time"
)

type InProgress struct {
	uiText  Text
	States  []string
	current int
}

func BuildInProgress() InProgress {
	uiInProgress := InProgress{
		uiText: BuildText("In progress"),
		States: []string{
			term.Style{Foreground: &term.Color{110, 110, 250}}.Sprint("⠋"),
			term.Style{Foreground: &term.Color{115, 115, 240}}.Sprint("⠙"),
			term.Style{Foreground: &term.Color{120, 120, 230}}.Sprint("⠹"),
			term.Style{Foreground: &term.Color{125, 125, 220}}.Sprint("⠸"),
			term.Style{Foreground: &term.Color{130, 130, 210}}.Sprint("⠼"),
			term.Style{Foreground: &term.Color{135, 135, 200}}.Sprint("⠴"),
			term.Style{Foreground: &term.Color{130, 130, 210}}.Sprint("⠦"),
			term.Style{Foreground: &term.Color{125, 125, 220}}.Sprint("⠧"),
			term.Style{Foreground: &term.Color{120, 120, 230}}.Sprint("⠇"),
			term.Style{Foreground: &term.Color{115, 115, 240}}.Sprint("⠏"),
		},
		current: -1,
	}
	return uiInProgress
}

func (uiInProgress *InProgress) Start() {
	if uiInProgress.current < 0 {
		uiInProgress.current = 0
		go uiInProgress.launch()
	}
}

func (uiInProgress *InProgress) launch() {
	for uiInProgress.current >= 0 {
		uiInProgress.Display()
		uiInProgress.current = (uiInProgress.current + 1) % len(uiInProgress.States)
		time.Sleep(200 * time.Millisecond)
	}
}

func (uiInProgress *InProgress) Stop() {
	uiInProgress.current = -1
	uiInProgress.Hide()
}

func (uiInProgress *InProgress) Hide() {
	var cursor term.Cursor
	cursor.Print("  ")
	uiInProgress.uiText.Hide()
	cursor.MoveTo(0, 0)
}

func (uiInProgress *InProgress) Display() {
	var cursor term.Cursor
	if uiInProgress.current >= 0 {
		waitChar := uiInProgress.States[uiInProgress.current]
		cursor.Print(waitChar, " ")
		uiInProgress.uiText.Display()
	}
	cursor.MoveTo(0, 0)
}

func (uiInProgress *InProgress) SetText(text string) {
	uiInProgress.uiText.SetText(text)
}
