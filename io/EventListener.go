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

import termbox "github.com/nsf/termbox-go"

// EventListener ...
type EventListener struct {
	keypressHandlers map[rune]func()
	resizeHandler    func()
}

// NewEventListener ...
// Constructs a new event listener.
func NewEventListener(resizeHandler func()) *EventListener {
	eventListener := &EventListener{resizeHandler: resizeHandler}
	eventListener.keypressHandlers = make(map[rune]func())

	return eventListener
}

// RegisterKeypressHandler ...
// Registers a new keypress handler.
func (e *EventListener) RegisterKeypressHandler(key rune, handler func()) {
	e.keypressHandlers[key] = handler
}

// WaitForEvent ...
// Waits for user input.
func (e *EventListener) WaitForEvent() {
	// Block and wait for input
	event := termbox.PollEvent()

	// Handle keypress
	if event.Type == termbox.EventKey {
		for key, value := range e.keypressHandlers {
			if event.Ch == key {
				value()
			}
		}
	}

	// Handle resize
	if event.Type == termbox.EventResize {
		e.resizeHandler()
	}
}
