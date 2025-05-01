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
	selection := ""
	selectui := termui.BuildSelect(options...)
	selectui.HeaderText = title

	term.Clear()
	selectui.Display()

	err := term.Read(func(event keyboard.KeyEvent, processor *term.EventProcessor) {
		if event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q' {
			processor.Stop()
			selectui.Hide()
		} else if event.Key == keyboard.KeyArrowUp {
			selectui.PreviousItem()
			selectui.Display()
		} else if event.Key == keyboard.KeyArrowDown {
			selectui.NextItem()
			selectui.Display()
		} else if event.Key == keyboard.KeyEnter {
			processor.Stop()
			selectui.Hide()
			selection = selectui.SelectedItem()
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
