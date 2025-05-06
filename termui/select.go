package termui

import (
	"github.com/deverdeb/bvmgo-term/term"
	"strings"
)

type Select struct {
	uiHeaderText       Text
	options            []selectItem
	current            int
	ItemStyle          term.Style
	ItemBullet         string
	SelectedItemStyle  term.Style
	SelectedItemBullet string

	maxWidth  int
	maxHeight int
}

type selectItem struct {
	value  string
	uiText Text
}

func BuildSelect(options ...string) Select {
	uiSelect := Select{
		uiHeaderText:       BuildText("Choose an item:"),
		options:            make([]selectItem, len(options)),
		current:            0,
		ItemStyle:          term.Style{},
		ItemBullet:         "\u25CB",
		SelectedItemStyle:  term.Style{Bold: true, Foreground: &term.Green},
		SelectedItemBullet: "\u25CF",
	}
	for idx, option := range options {
		uiSelect.options[idx] = selectItem{
			value:  option,
			uiText: BuildText(option),
		}
	}
	return uiSelect
}

func (uiSelect *Select) NextItem() {
	uiSelect.current += 1
	if uiSelect.current >= len(uiSelect.options) {
		uiSelect.current = 0
	}
}

func (uiSelect *Select) PreviousItem() {
	uiSelect.current -= 1
	if uiSelect.current < 0 {
		uiSelect.current = len(uiSelect.options) - 1
	}
}

func (uiSelect *Select) SelectedItem() string {
	return uiSelect.options[uiSelect.current].value
}

func (uiSelect *Select) Hide() {
	var cursor term.Cursor
	width, height := uiSelect.Dimension()
	line := strings.Repeat(" ", width)
	for posRow := 0; posRow <= height; posRow++ {
		cursor.MoveTo(0, posRow)
		cursor.Print(line)
	}
	cursor.MoveTo(0, 0)
}

func (uiSelect *Select) Display() {
	var cursor term.Cursor
	uiSelect.uiHeaderText.Display()
	_, height := uiSelect.uiHeaderText.Dimension()
	cursor.MoveDown(height)
	uiSelect.displayOptions(&cursor)
	cursor.MoveTo(0, 0)
}

func (uiSelect *Select) displayOptions(cursor *term.Cursor) {
	for idx, option := range uiSelect.options {
		bullet := uiSelect.ItemBullet
		style := uiSelect.ItemStyle
		if idx == uiSelect.current {
			bullet = uiSelect.SelectedItemBullet
			style = uiSelect.SelectedItemStyle
		}
		cursor.Print(style.Sprintf(" %s ", bullet))
		option.uiText.Display()
		_, height := option.uiText.Dimension()
		cursor.MoveDown(height)
		cursor.MoveLeft(3)
	}
}

func (uiSelect *Select) Dimension() (width, height int) {
	width, height = uiSelect.uiHeaderText.Dimension()
	for _, option := range uiSelect.options {
		optWidth, optHeight := option.uiText.Dimension()
		width = max(width, optWidth+3)
		height += optHeight
	}
	return width, height
}

func (uiSelect *Select) SetText(header string) {
	uiSelect.uiHeaderText.SetText(header)
}

func (uiSelect *Select) SetMaxDimension(maxWidth, maxHeight int) {
	uiSelect.SetMaxWidth(maxWidth)
	uiSelect.SetMaxHeight(maxHeight)
}

func (uiSelect *Select) SetMaxWidth(maxWidth int) {
	uiSelect.maxWidth = maxWidth
	uiSelect.uiHeaderText.SetMaxWidth(maxWidth)
	for _, option := range uiSelect.options {
		option.uiText.SetMaxHeight(maxWidth - 3)
	}
}

func (uiSelect *Select) SetMaxHeight(maxHeight int) {
	uiSelect.maxHeight = maxHeight
	uiSelect.uiHeaderText.SetMaxHeight(maxHeight)
}
