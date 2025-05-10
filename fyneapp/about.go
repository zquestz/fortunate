package fyneapp

import (
	"fmt"
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var aboutWindow fyne.Window

// about creates/shows the about window.
func about() {
	if aboutWindow != nil {
		aboutWindow.Show()
		return
	}

	aboutWindow = appGUI.NewWindow(fmt.Sprintf("About %s", config.GUIAppName))

	ctrlW := &desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault}
	aboutWindow.Canvas().AddShortcut(ctrlW, func(shortcut fyne.Shortcut) {
		aboutWindow.Hide()
	})

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

	aboutWindow.SetCloseIntercept(func() {
		aboutWindow.Hide()
	})

	aboutWindow.SetFixedSize(true)
	aboutWindow.CenterOnScreen()
	aboutWindow.Show()
}
