package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var appGUI fyne.App

func Run(trayReady func(), trayExit func()) {
	appGUI = app.New()

	trayReady()
	appGUI.Run()
	trayExit()
}

func Quit() {
	appGUI.Quit()
}
