package fyneapp

import (
	"fmt"
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"
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
		[]string{"default", "light", "dark"},
		nil,
	)

	switch selectedTheme := config.AppConfig.IconTheme; selectedTheme {
	case "dark":
		iconTheme.SetSelected("dark")
	case "light":
		iconTheme.SetSelected("light")
	default:
		iconTheme.SetSelected("default")
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Icon Theme", Widget: iconTheme}},
		OnSubmit: func() { // optional, handle form submission
			config.AppConfig.IconTheme = iconTheme.Selected
			theme.SetIconTheme()
			config.AppConfig.Save()
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
	// preferencesWindow.SetFixedSize(true)
	preferencesWindow.Show()
}

func Display() {

}

func Quit() {
	appGUI.Quit()
}
