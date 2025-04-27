package termui

import "github.com/deverdeb/bvmgo-term/term"

type Select struct {
	HeaderText string
	options    []string
	current    int
	Standard   term.Style
	Selected   term.Style
}

func BuildSelect(options ...string) Select {
	uiSelect := Select{
		HeaderText: "Choose an item:",
		options:    make([]string, 0),
		current:    0,
		Standard:   term.Style{},
		Selected:   term.Style{Bold: true, Foreground: &term.Green},
	}
	uiSelect.options = append(uiSelect.options, options...)
	return uiSelect
}

func (uiSelect *Select) Display() {
	term.Print(uiSelect.HeaderText)
	uiSelect.displayOptions()
	term.CursorMoveUp(len(uiSelect.options) + 1)
	term.CursorMoveRight(len(uiSelect.HeaderText) + 1)
}

func (uiSelect *Select) displayOptions() {
	for idx, option := range uiSelect.options {
		cursor := "*"
		style := uiSelect.Standard
		if idx == uiSelect.current {
			cursor = ">"
			style = uiSelect.Selected
		}
		term.Print(style.Sprintf(" %s %s\n", cursor, option))
	}
}
