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
	"github.com/stuartthompson/dockterm/screens"
)

type DisplayScreen int

const (
	MainScreen = iota
	AboutScreen
)

func drawScreen(screen DisplayScreen) {
	switch screen {
	case MainScreen:
		screens.RenderMainScreen()
	case AboutScreen:
		screens.RenderAboutScreen()
	}
}

func mainLoop() {
	var displayScreen DisplayScreen
	displayScreen = MainScreen
	drawScreen(displayScreen)
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

	mainLoop()
}
