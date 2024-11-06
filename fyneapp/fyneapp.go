package fyneapp

import (
	"fmt"
	"strings"
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"
	"github.com/zquestz/fortunate/notify"
	"github.com/zquestz/fortunate/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var appGUI fyne.App
var aboutWindow fyne.Window
var preferencesWindow fyne.Window

func Run(trayReady func(), trayExit func()) {
	appGUI = app.New()

	trayReady()
	appGUI.Run()
	trayExit()
}

func About() {
	// Window was already hidden, just show it.
	if aboutWindow != nil {
		aboutWindow.Show()
		return
	}

	aboutWindow = appGUI.NewWindow(fmt.Sprintf("About %s", config.GUIAppName))

	i := canvas.NewImageFromResource(&fyne.StaticResource{
		StaticName:    config.GUIAppName,
		StaticContent: icon.Data,
	})
	i.FillMode = canvas.ImageFillOriginal
	name := widget.NewLabel(config.GUIAppName)
	name.Alignment = fyne.TextAlignCenter
	name.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	version := widget.NewLabel(fmt.Sprintf("Version %s", config.Version))
	version.Alignment = fyne.TextAlignCenter
	copy := widget.NewLabel(fmt.Sprintf("Copyright © %d Josh Ellithorpe", time.Now().Year()))
	copy.Alignment = fyne.TextAlignCenter

	aboutWindow.SetContent(container.NewVBox(
		i,
		name,
		version,
		copy,
	))

	// Hack to make sure we don't accidently quit.
	aboutWindow.SetCloseIntercept(func() {
		aboutWindow.Hide()
	})

	aboutWindow.CenterOnScreen()
	aboutWindow.SetFixedSize(true)
	aboutWindow.Show()
}

func Preferences() {
	// Window was already hidden, just show it.
	if preferencesWindow != nil {
		preferencesWindow.Show()
		return
	}

	preferencesWindow = appGUI.NewWindow("Preferences")

	iconTheme := widget.NewSelect(
		[]string{"Default", "Light", "Dark"},
		nil,
	)

	switch selectedTheme := strings.ToLower(config.AppConfig.IconTheme); selectedTheme {
	case "dark":
		iconTheme.SetSelected("Dark")
	case "light":
		iconTheme.SetSelected("Light")
	default:
		iconTheme.SetSelected("Default")
	}

	notifyFortuneTimes := widget.NewSelect(
		[]string{"Disabled", "1 hour", "3 hours", "6 hours", "12 hours", "24 hours"},
		nil,
	)

	switch config.AppConfig.FortuneTimer {
	case 1:
		notifyFortuneTimes.SetSelected("1 hour")
	case 3:
		notifyFortuneTimes.SetSelected("3 hours")
	case 6:
		notifyFortuneTimes.SetSelected("6 hours")
	case 12:
		notifyFortuneTimes.SetSelected("12 hours")
	case 24:
		notifyFortuneTimes.SetSelected("24 hours")
	default:
		notifyFortuneTimes.SetSelected("Disabled")
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Icon Theme", Widget: iconTheme},
			{Text: "Fortune Timer", Widget: notifyFortuneTimes},
		},
		OnSubmit: func() { // optional, handle form submission
			config.AppConfig.IconTheme = strings.ToLower(iconTheme.Selected)
			switch notifyFortuneTimes.Selected {
			case "1 hour":
				config.AppConfig.FortuneTimer = 1
			case "3 hours":
				config.AppConfig.FortuneTimer = 3
			case "6 hours":
				config.AppConfig.FortuneTimer = 6
			case "12 hours":
				config.AppConfig.FortuneTimer = 12
			case "24 hours":
				config.AppConfig.FortuneTimer = 24
			default:
				config.AppConfig.FortuneTimer = 0
			}
			err := config.AppConfig.Save()
			if err != nil {
				notify.NotifyError(err)
			}
			theme.SetIconTheme()
			notify.FortuneTickerReset()
			preferencesWindow.Hide()
		},
		OnCancel: func() {
			preferencesWindow.Hide()
		},
	}

	content := container.NewVBox(
		form,
	)

	// Hack to make sure we don't accidently quit.
	preferencesWindow.SetCloseIntercept(func() {
		preferencesWindow.Hide()
	})

	preferencesWindow.SetContent(content)
	preferencesWindow.CenterOnScreen()
	preferencesWindow.SetFixedSize(true)
	preferencesWindow.Show()
}

func Display() {

}

func Quit() {
	appGUI.Quit()
}
