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

var output = os.Stderr

var termInformation *terminalInformation = nil

func Initialize() error {
	if termInformation != nil {
		return nil
	}
	var err error
	termInformation, err = initializeWithHandler(output.Fd())
	if termInformation.isValid {
		return err
	}
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
		err = fmt.Errorf("terminal not found")
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
		return 0, 0, fmt.Errorf("terminal not found")
	}
	return termInformation.width, termInformation.height, nil
}

func Println(text ...any) {
	_, _ = fmt.Fprintln(output, text...)
}

func Print(text ...any) {
	_, _ = fmt.Fprint(output, text...)
}

func Printf(format string, params ...any) {
	Print(fmt.Sprintf(format, params...))
}

func Clear() {
	Print(ansi.ClearScreen())
}
