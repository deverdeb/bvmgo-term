package main

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/term"
	"github.com/deverdeb/bvmgo-term/termui"
	"github.com/eiannone/keyboard"
	"os"
)

func main() {
	title, options := extractParameter()
	selection := executeSelect(title, options...)
	fmt.Print(selection)
}

func extractParameter() (title string, options []string) {
	commandLineArguments := os.Args
	if len(commandLineArguments) < 3 {
		displayHelp()
		os.Exit(1)
	}
	return os.Args[1], os.Args[2:]
}

func executeSelect(title string, options ...string) string {

	width, height, err := term.GetSize()
	if err != nil {
		fmt.Println("terminal access error:", err)
		return ""
	}

	term.CursorHide()
	defer term.CursorShow()

	term.Clear()

	uiSelect := termui.BuildSelect(options...)
	uiSelect.SetText(title)
	uiSelect.SetMaxWidth(width)
	uiSelect.SetMaxHeight(height)

	uiSelect.Display()

	selection := ""
	err = term.Read(func(event keyboard.KeyEvent, processor *term.EventProcessor) {
		if event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q' {
			processor.Stop()
			uiSelect.Hide()
		} else if event.Key == keyboard.KeyArrowUp {
			uiSelect.PreviousItem()
			uiSelect.Display()
		} else if event.Key == keyboard.KeyArrowDown {
			uiSelect.NextItem()
			uiSelect.Display()
		} else if event.Key == keyboard.KeyEnter {
			processor.Stop()
			uiSelect.Hide()
			selection = uiSelect.SelectedItem()
		}
	})
	if err != nil {
		fmt.Println("execution error:", err)
		return ""
	}
	return selection
}

func displayHelp() {
	fmt.Println(`Usage: select TITLE OPTIONS...
Display a list of items.
User can:
    Choose item with UP and DOWN arrow keys.
    Validate with ENTER key, selected value is return.
    Quit with ESC or Q keys, empty value is return.

Arguments:
    TITLE      List title
    OPTIONS... Items list (required 2 or more values)

Example: select "Select an item:" "first item" "seconde item" "other item"`)
}
