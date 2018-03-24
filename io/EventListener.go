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
