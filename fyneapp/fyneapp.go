package fyneapp

import (
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/systray"
)

const appID = "at.greyh.fortunate"

var appGUI fyne.App
var desktopApp desktop.App

// Run starts the Fyne application.
func Run() {
	appGUI = app.New()
	app.SetMetadata(fyne.AppMetadata{
		ID:      appID,
		Name:    config.GUIAppName,
		Version: config.Version,
		Icon: &fyne.StaticResource{
			StaticName:    config.GUIAppName,
			StaticContent: icon.Data,
		},
	})

	desktopApp = appGUI.(desktop.App)
	setSystrayMenu()
	setSystrayIcon()
	setActivate()

	go startFortuneTicker()

	appGUI.Run()
}

func setSystrayIcon() {
	if desktopApp == nil {
		return
	}

	resource := &fyne.StaticResource{
		StaticName: config.GUIAppName,
	}

	switch iconTheme := config.AppConfig.IconTheme; iconTheme {
	case "dark":
		resource.StaticContent = icon.DataDark
	case "light":
		resource.StaticContent = icon.DataLight
	default:
		resource.StaticContent = icon.Data
	}

	desktopApp.SetSystemTrayIcon(resource)
}

func setSystrayMenu() {
	if desktopApp == nil {
		return
	}

	m := fyne.NewMenu(config.GUIAppName,
		fyne.NewMenuItem("Display Fortune", func() {
			display()
		}), fyne.NewMenuItem("Notify Fortune", func() {
			notifyFortune()
		}), fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Settings", func() {
			settings()
		}), fyne.NewMenuItem("About", func() {
			about()
		}))
	desktopApp.SetSystemTrayMenu(m)
}

func setActivate() {
	if desktopApp == nil {
		return
	}

	systray.SetActivateFunc(display)
}
