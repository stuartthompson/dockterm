package screens

import (
	"github.com/stuartthompson/dockterm/io"
)

// MainScreen ...
type MainScreen struct {
}

// Render ...
// Renders the main screen.
func (s *MainScreen) Render() {
	io.ClearScreen(0)
	io.RenderText("Main", 1, 1, 255, 0)
	io.Flush()
}
