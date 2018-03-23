package screens

import (
	"github.com/stuartthompson/dockterm/io"
)

func RenderMainScreen() {
	io.ClearScreen(0)
	io.RenderText("Main", 1, 1, 255, 0)
	io.Flush()
}
