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
	"log"
	"os/exec"
	"strings"

	"github.com/stuartthompson/dockterm/io"
)

// MainScreen ...
type MainScreen struct {
}

// Render ...
// Renders the main screen.
func (s *MainScreen) Render() {
	io.ClearScreen(0)
	width, height := io.GetWindowSize()
	io.RenderPaneBorder(0, 0, width-1, height-1, 82, 0)
	io.RenderText("Main", 1, 1, 255, 0)

	// List running containers
	ctrs := s.getRunningContainers()
	ctrList := strings.Split(ctrs, "\n")
	for ix, ctr := range ctrList {
		io.RenderText(ctr, 1, 3+ix, 255, 0)
	}
	io.Flush()
}

func (s *MainScreen) getRunningContainers() string {
	output, err := exec.Command("/usr/bin/docker", "ps", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}
