//go:build linux || windows || dragonfly || freebsd || netbsd || openbsd || darwin

package tray

import (
	"github.com/zquestz/fortunate/fyneapp"
	"github.com/zquestz/fortunate/notify"
	"github.com/zquestz/fortunate/theme"

	"fyne.io/systray"
)

const (
	appName = "Fortunate"
)

func Run() (func(), func()) {
	return systray.RunWithExternalLoop(onReady, onExit)
}

func onReady() {
	theme.SetIconTheme()

	systray.SetTitle(appName)
	systray.SetTooltip(appName)

	mDisplay := systray.AddMenuItem("Display Fortune", "Display Fortune")
	mNotify := systray.AddMenuItem("Notify Fortune", "Notify Fortune")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "About")
	mSettings := systray.AddMenuItem("Settings", "Settings")
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
			case <-mSettings.ClickedCh:
				settings()
			case <-mQuit.ClickedCh:
				quit()
				return
			}
		}
	}()
}

func display() {
	fyneapp.Display()
}

func sendNotification() {
	err := notify.NotifyFortune()
	if err != nil {
		notify.NotifyError(err)
	}
}

func about() {
	fyneapp.About()
}

func settings() {
	fyneapp.Settings()
}

func quit() {
	fyneapp.Quit()
}

func onExit() {
}
