//go:build linux || windows || dragonfly || freebsd || netbsd || openbsd || darwin
// +build linux windows dragonfly freebsd netbsd openbsd darwin

package tray

import (
	"github.com/getlantern/systray"
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
	"github.com/zquestz/fortunate/icon"
	"github.com/zquestz/fortunate/notify"
)

const (
	appName = "Fortunate"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	switch iconTheme := config.AppConfig.IconTheme; iconTheme {
	case "dark":
		systray.SetTemplateIcon(icon.DataDark, icon.DataDark)
	case "light":
		systray.SetTemplateIcon(icon.DataLight, icon.DataLight)
	default:
		systray.SetTemplateIcon(icon.Data, icon.Data)
	}

	systray.SetTitle(appName)
	systray.SetTooltip(appName)

	mDisplay := systray.AddMenuItem("Display Fortune", "Display Fortune")
	mNotify := systray.AddMenuItem("Notify Fortune", "Notify Fortune")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "About")
	mPreferences := systray.AddMenuItem("Preferences", "Preferences")
	mQuit := systray.AddMenuItem("Quit", "Quit")

	go func() {
		for {
			select {
			case <-mDisplay.ClickedCh:
				display()
			case <-mNotify.ClickedCh:
				sendNotification()
			case <-mAbout.ClickedCh:
				about()
			case <-mPreferences.ClickedCh:
				preferences()
			case <-mQuit.ClickedCh:
				quit()
				return
			}
		}
	}()
}

func display() {
}

func sendNotification() {
	output, _ := fortune.Run()

	notify.Notify(config.GUIAppName, config.GUIAppName, output, "")
}

func about() {
}

func preferences() {
}

func quit() {
	systray.Quit()
}

func onExit() {
}
