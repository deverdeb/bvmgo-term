package main

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/term"
	"github.com/deverdeb/bvmgo-term/termui"
	"github.com/eiannone/keyboard"
	"time"
)

func main() {
	uiInProgress := termui.BuildInProgress()
	uiInProgress.SetText("Press 'q' to quit")

	term.Clear()
	uiInProgress.Start()
	quit := false

	err := term.Read(func(event keyboard.KeyEvent, processor *term.EventProcessor) {
		if event.Key == keyboard.KeyEsc || event.Rune == 'q' || event.Rune == 'Q' {
			uiInProgress.Stop()
			quit = true
		} else if event.Rune >= 'a' && event.Rune <= 'z' {
			uiInProgress.SetText(fmt.Sprintf("You press '%s' key. Press 'q' to quit", string(event.Rune)))
		}
	})
	if err != nil {
		fmt.Println("execution error:", err)
	}

	for !quit {
		time.Sleep(100 * time.Millisecond)
	}
}
