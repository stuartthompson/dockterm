// Copyright 2018 Stuart Thompson.

// This file is part of Dockterm.

// Dockterm is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Foobar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

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
	width, height := io.GetWindowSize()
	io.RenderPaneBorder(0, 0, width - 1, height - 1, 212, 0)
	io.RenderText("About", 1, 1, 255, 0)
	io.RenderText("Dockterm is a docker management app.", 1, 3, 255, 0)
	io.Flush()
}
