package main

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/term"
	"github.com/deverdeb/bvmgo-term/termui"
	"github.com/eiannone/keyboard"
)

func main() {
	uiInProgress := termui.BuildInProgress()
	uiInProgress.SetText("Press 'q' to quit")

	term.Clear()
	uiInProgress.Start()

	err := term.Read(func(event keyboard.KeyEvent, processor *term.EventProcessor) {
		if event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q' {
			processor.Stop()
		} else if event.Rune >= 'a' && event.Rune <= 'z' {
			uiInProgress.SetText(fmt.Sprintf("You press '%s' key. Press 'q' to quit", string(event.Rune)))
		}
	})
	uiInProgress.Stop()
	if err != nil {
		fmt.Println("execution error:", err)
	}
	fmt.Println("Good bye")
}
