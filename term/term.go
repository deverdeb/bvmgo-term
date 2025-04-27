package term

import (
	"fmt"
	"github.com/deverdeb/bvmgo-term/ansi"
	"golang.org/x/term"
	"os"
)

type terminalInformation struct {
	windowHandle int
	isValid      bool
	width        int
	height       int
}

var termInformation *terminalInformation = nil

func Initialize() error {
	if termInformation != nil {
		return nil
	}
	var err error
	termInformation, err = initializeWithHandler(os.Stdout.Fd())
	if termInformation.isValid {
		return err
	}
	termInformation, err = initializeWithHandler(os.Stderr.Fd())
	if termInformation.isValid {
		return err
	}
	termInformation, err = initializeWithHandler(os.Stdin.Fd())
	return err
}

func initializeWithHandler(windowHandle uintptr) (*terminalInformation, error) {
	windowInfo := &terminalInformation{
		windowHandle: int(windowHandle),
		isValid:      false,
		width:        0,
		height:       0,
	}
	var err error
	if term.IsTerminal(windowInfo.windowHandle) {
		windowInfo.width, windowInfo.height, err = term.GetSize(windowInfo.windowHandle)
		if err == nil && windowInfo.width > 0 {
			windowInfo.isValid = true
			return windowInfo, err
		}
	} else {
		err = fmt.Errorf("bvm-term not found")
	}
	return windowInfo, err
}

func GetSize() (width, height int, err error) {
	if termInformation == nil {
		err = Initialize()
		if err != nil {
			return 0, 0, err
		}
	}
	if !termInformation.isValid {
		return 0, 0, fmt.Errorf("bvm-term not found")
	}
	return termInformation.width, termInformation.height, nil
}

func Printf(format string, params ...any) {
	fmt.Printf(format, params...)
}

func Clear() {
	Printf(ansi.ClearScreen())
}
