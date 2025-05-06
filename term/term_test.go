package term

import "os"

func ExamplePrintln() {
	output = os.Stdout
	Println("hello world")
	// Output:
	// hello world
	//
}

func ExamplePrint() {
	output = os.Stdout
	Print("hello world")
	// Output:
	// hello world
}

func ExamplePrintf() {
	output = os.Stdout
	Printf("hello %s", "world")
	// Output:
	// hello world
}

func ExampleCursorMoveTo() {
	output = os.Stdout
	CursorMoveTo(2, 3)
	Print("hello world")
	// Output:
	// [3;2Hhello world
}
