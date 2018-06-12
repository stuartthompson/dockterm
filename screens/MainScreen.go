// Copyright 2018 Stuart Thompson.

// This file is part of Dockterm.

// Dockterm is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Dockterm is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Dockterm. If not, see <http://www.gnu.org/licenses/>.

package screens

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/stuartthompson/dockterm/entities"
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

	// Render headers
	io.RenderText("Container Id", 1, 3, 255, 0)
	io.RenderText("Running", 14, 3, 255, 0)

	// Render container details
	containers := s.getRunningContainers()
	for ix, container := range containers {
		shortID := container.ID[0:12]
		running := "No"
		if container.State.Running {
			running = "Yes"
		}
		io.RenderText(shortID, 1, 4+ix, 255, 0)
		io.RenderText(running, 14, 4+ix, 255, 0)
	}
	io.Flush()
}

// getRunningContainers ...
// Gets the list of containers running on this machine.
func (s *MainScreen) getRunningContainers() []*entities.Container {
	output, err := exec.Command("docker", "ps", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}

	ctrs := string(output)

	// Strip trailing LF (to avoid parsing an empty line at the end)
	ctrs = ctrs[0 : len(ctrs)-1]
	// Split on LF
	ctrList := strings.Split(ctrs, "\n")
	// Iterate through the lines and build container descriptions
	var containers []*entities.Container
	for ix, ctr := range ctrList {
		if ix == 0 {
			continue
		}
		// Split on tab
		ctrID := ctr[0:9]
		container := s.getContainerDetails(ctrID)
		containers = append(containers, container)
	}
	// Parse container id
	return containers
}

// getContainerDetails ...
// Gets the details for a running Docker container.
func (s *MainScreen) getContainerDetails(ctrID string) *entities.Container {
	output, err := exec.Command("docker", "inspect", ctrID).Output()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the output
	var container []entities.Container
	err = json.Unmarshal([]byte(output), &container)
	if err != nil {
		log.Fatal(err)
	}

	return &container[0]
}
