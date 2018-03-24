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
