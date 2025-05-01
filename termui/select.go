package termui

import (
	"github.com/deverdeb/bvmgo-term/term"
	"strings"
)

type Select struct {
	HeaderText         string
	options            []string
	current            int
	ItemStyle          term.Style
	ItemBullet         string
	SelectedItemStyle  term.Style
	SelectedItemBullet string

	minRow, maxRow, minColomn, maxColumn int
}

func BuildSelect(options ...string) Select {
	uiSelect := Select{
		HeaderText:         "Choose an item",
		options:            make([]string, 0),
		current:            0,
		ItemStyle:          term.Style{},
		ItemBullet:         "\u25CB",
		SelectedItemStyle:  term.Style{Bold: true, Foreground: &term.Green},
		SelectedItemBullet: "\u25CF",
	}
	uiSelect.options = append(uiSelect.options, options...)
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
	return uiSelect.options[uiSelect.current]
}

func (uiSelect *Select) Hide() {
	var cursor term.Cursor
	line := strings.Repeat(" ", uiSelect.maxColumn-uiSelect.minColomn)
	for posRow := uiSelect.minRow; posRow <= uiSelect.maxRow; posRow++ {
		cursor.MoveTo(uiSelect.minColomn, posRow)
		cursor.Print(line)
	}
	cursor.MoveTo(0, 0)
	uiSelect.minColomn, uiSelect.minRow = 0, 0
	uiSelect.maxColumn, uiSelect.maxRow = 0, 0
}

func (uiSelect *Select) Display() {
	uiSelect.minColomn, uiSelect.minRow = 0, 0
	uiSelect.maxColumn, uiSelect.maxRow = 0, 0
	var cursor term.Cursor
	cursor.Printf("%s:\n", uiSelect.HeaderText)
	uiSelect.displayOptions(&cursor)
	uiSelect.minColomn, uiSelect.minRow = cursor.MinPosition()
	uiSelect.maxColumn, uiSelect.maxRow = cursor.MaxPosition()
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
		cursor.Print(style.Sprintf(" %s %s\n", bullet, option))
	}
}
