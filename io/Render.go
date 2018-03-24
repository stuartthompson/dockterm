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
