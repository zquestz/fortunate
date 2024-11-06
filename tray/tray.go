//go:build linux || windows || dragonfly || freebsd || netbsd || openbsd || darwin
// +build linux windows dragonfly freebsd netbsd openbsd darwin

package tray

import (
	"fyne.io/systray"
	"github.com/zquestz/fortunate/fyneapp"
	"github.com/zquestz/fortunate/notify"
	"github.com/zquestz/fortunate/theme"
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
	fyneapp.Display()
}

func sendNotification() {
	notify.NotifyFortune()
}

func about() {
	fyneapp.About()
}

func preferences() {
	fyneapp.Preferences()
}

func quit() {
	fyneapp.Quit()
}

func onExit() {
}
