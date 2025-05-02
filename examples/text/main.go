package main

import (
	"github.com/deverdeb/bvmgo-term/term"
	"github.com/deverdeb/bvmgo-term/termui"
)

func main() {
	text := term.StyleOk.Sprint("Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium ") +
		term.StyleInfo.Sprint("doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore ") +
		term.StyleError.Sprint("veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ") +
		term.StyleWarn.Sprint("ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia ") +
		"consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui " +
		"dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora " +
		term.StyleInfo.Sprint("incidunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad ") +
		"minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea " +
		"commodi consequatur?"

	width, _, _ := term.GetSize()

	textui := termui.BuildText(text)
	textui.SetMaxWidth(width / 2)
	_, height := textui.Dimension()

	term.Clear()
	term.CursorMoveRight(width / 4)
	textui.Display()
	term.CursorMoveDown(height + 1)
	term.CursorMoveLeft(width / 4)
}
