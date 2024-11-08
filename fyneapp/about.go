package fyneapp

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"
)

var aboutWindow fyne.Window

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

	aboutWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == "Escape" {
			aboutWindow.Hide()
		}
	})

	// Hack to make sure we don't accidently quit.
	aboutWindow.SetCloseIntercept(func() {
		aboutWindow.Hide()
	})

	aboutWindow.CenterOnScreen()
	aboutWindow.SetFixedSize(true)
	aboutWindow.Show()
}
