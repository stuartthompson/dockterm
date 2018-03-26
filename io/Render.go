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

const (
	BorderRuneHorizontal = rune('\u2550')
	BorderRuneVertical = rune('\u2551')
	BorderRuneCornerTopLeft = rune('\u2554')
	BorderRuneCornerTopRight = rune('\u2557')
	BorderRuneCornerBottomLeft = rune('\u255A')
	BorderRuneCornerBottomRight = rune('\u255D')
)

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

func GetWindowSize() (width int, height int) {
	return termbox.Size()
}

// RenderPaneBorder ...
// Renders a border for a window pane.
func RenderPaneBorder(x int, y int, width int, height int, fgColor int, bgColor int) {
	// Render the corners
	termbox.SetCell(x, y, BorderRuneCornerTopLeft, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x + width, y, BorderRuneCornerTopRight, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x, y + height, BorderRuneCornerBottomLeft, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	termbox.SetCell(x + width, y + height, BorderRuneCornerBottomRight, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	// Render top border
	for ix := 1; ix < width; ix++ {
		termbox.SetCell(x + ix, y, BorderRuneHorizontal, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render bottom border
	for ix := 1; ix < width; ix++ {
		termbox.SetCell(x + ix, y + height, BorderRuneHorizontal, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render left border
	for iy := 1; iy < height; iy++ {
		termbox.SetCell(x, y + iy, BorderRuneVertical, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
	// Render right border
	for iy := 1; iy < height; iy++ {
		termbox.SetCell(x + width, y + iy, BorderRuneVertical, termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
}

// RenderText ...
// Renders a string at specific coordinates using the supplied colors.
func RenderText(text string, x int, y int, fgColor int, bgColor int) {
	for i := 0; i < len(text); i++ {
		termbox.SetCell(x+i, y, rune(text[i]), termbox.Attribute(fgColor), termbox.Attribute(bgColor))
	}
}
