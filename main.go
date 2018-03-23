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

package main

import (
	"github.com/nsf/termbox-go"
)

type DisplayScreen int

const (
	MainScreen = iota
	AboutScreen
)

func drawBackground(bgColor termbox.Attribute) {
	termbox.Clear(0, bgColor)
}

func drawScreen(screen DisplayScreen) {
	switch screen {
	case MainScreen:
		drawMainScreen()
	case AboutScreen:
		drawAboutScreen()
	}
}

func drawMainScreen() {
	drawBackground(termbox.Attribute(1))
	termbox.SetCell(10, 0, 'M', termbox.Attribute(80), termbox.Attribute(125))
	termbox.SetCell(11, 0, 'a', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(12, 0, 'i', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(13, 0, 'n', termbox.Attribute(80), termbox.Attribute(1))
	termbox.Flush()
}

func drawAboutScreen() {
	drawBackground(termbox.Attribute(1))
	termbox.SetCell(10, 0, 'A', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(11, 0, 'b', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(12, 0, 'o', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(13, 0, 'u', termbox.Attribute(80), termbox.Attribute(1))
	termbox.SetCell(14, 0, 't', termbox.Attribute(80), termbox.Attribute(1))
	termbox.Flush()
}

func mainLoop() {
	var displayScreen DisplayScreen
	displayScreen = MainScreen
	drawMainScreen()
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Ch == '?' {
				if displayScreen == AboutScreen {
					displayScreen = MainScreen
				} else {
					displayScreen = AboutScreen
				}
			}
			if event.Ch == 'q' {
				break
			}

			drawScreen(displayScreen)
		}
		if event.Type == termbox.EventResize {
			drawScreen(displayScreen)
		}

	}
}

func main() {
	var err error

	err = termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

	mainLoop()
}
