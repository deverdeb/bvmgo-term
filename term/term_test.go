package term

func ExamplePrint() {
	Print("hello world")
	// Output:
	// hello world
}

func ExamplePrintf() {
	Printf("hello %s", "world")
	// Output:
	// hello world
}

func ExampleCursorMoveTo() {
	CursorMoveTo(2, 3)
	Print("hello world")
	// Output:
	// [3;2Hhello world
}
