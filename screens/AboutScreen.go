package screens

import (
	"github.com/stuartthompson/dockterm/io"
)

func RenderAboutScreen() {
	io.ClearScreen(0)
	io.RenderText("About", 1, 1, 255, 0)
	io.Flush()
}
