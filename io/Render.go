package io

import (
	"github.com/nsf/termbox-go"
)

// Init ...
// Initializes output.
func Init() {
	termbox.SetOutputMode(termbox.Output256)
}

// ClearScreen ...
// Clears the screen using a specified color.
func ClearScreen(bgColor int) {
	termbox.Clear(0, termbox.Attribute(bgColor))
}

// Flush ...
// Flush render commands.
func Flush() {
	termbox.Flush()
}

// RenderText ...
// Renders a string at specific coordinates using the supplied colors.
func RenderText(text string, x int, y int, fgColor int, bgColor int) {
	for i := 0; i < len(text); i++ {
		termbox.SetCell(x+i, y, rune(text[i]), termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
}
