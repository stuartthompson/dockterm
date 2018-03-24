package screens

import (
	"github.com/stuartthompson/dockterm/io"
)

// AboutScreen ...
type AboutScreen struct {
}

// Render ...
// Renders the about screen.
func (s *AboutScreen) Render() {
	io.ClearScreen(0)
	io.RenderText("About", 1, 1, 255, 0)
	io.Flush()
}
