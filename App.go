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
	"github.com/stuartthompson/dockterm/io"
	"github.com/stuartthompson/dockterm/screens"
)

type Screen int

const (
	MainScreen = iota
	AboutScreen
)

type App struct {
	isRunning     bool
	eventListener *io.EventListener
	currentScreen Screen
	mainScreen    *screens.MainScreen
	aboutScreen   *screens.AboutScreen
}

func InitApp() *App {
	app := &App{
		isRunning:   true,
		mainScreen:  &screens.MainScreen{},
		aboutScreen: &screens.AboutScreen{},
	}
	app.eventListener = io.NewEventListener(app.Render)

	return app
}

func (a *App) Run() {
	io.Init()
	a.registerKeypressHandlers()
	a.Render()
	for a.isRunning {
		a.eventListener.WaitForEvent()
		a.Render()
	}
}

func (a *App) Render() {
	switch a.currentScreen {
	case MainScreen:
		a.mainScreen.Render()
	case AboutScreen:
		a.aboutScreen.Render()
	}
}

func (a *App) registerKeypressHandlers() {
	a.eventListener.RegisterKeypressHandler('?', a.toggleAboutScreen)
	a.eventListener.RegisterKeypressHandler('q', a.onQuit)
}

func (a *App) toggleAboutScreen() {
	if a.currentScreen == AboutScreen {
		a.currentScreen = MainScreen
	} else {
		a.currentScreen = AboutScreen
	}
}

func (a *App) onQuit() {
	a.isRunning = false
}
