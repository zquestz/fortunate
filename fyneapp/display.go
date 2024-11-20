package fyneapp

import (
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
	"github.com/zquestz/fortunate/icon"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var displayWindow fyne.Window

// display creates/shows the display fortune window.
func display() {
	if displayWindow != nil && displayWindow.Content().Visible() {
		displayWindow.Show()
		return
	}

	var newDisplayWindow fyne.Window

	newDisplayWindow = appGUI.NewWindow(config.GUIAppName)

	ctrlW := &desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault}
	newDisplayWindow.Canvas().AddShortcut(ctrlW, func(shortcut fyne.Shortcut) {
		newDisplayWindow.Hide()
	})

	i := canvas.NewImageFromResource(&fyne.StaticResource{
		StaticName:    config.GUIAppName,
		StaticContent: icon.Data,
	})
	i.FillMode = canvas.ImageFillOriginal

	fortuneCookie, fortuneContent, err := fortune.Run()
	if err != nil {
		notifyError(err)
	}

	fortune := widget.NewTextGrid()
	fortune.TabWidth = 8
	fortune.SetText(fortuneContent)

	copyFunc := func() {
		newDisplayWindow.Clipboard().SetContent(fortuneContent)
	}

	closeFunc := func() {
		newDisplayWindow.Hide()
	}

	nextFunc := func() {
		newDisplayWindow.Hide()
		display()
	}

	var cookie fyne.CanvasObject
	if config.AppConfig.ShowCookie {
		cookie = widget.NewLabel(fortuneCookie)
	}

	copy := widget.NewButtonWithIcon("", theme.ContentCopyIcon(), copyFunc)
	close := widget.NewButton("Close", closeFunc)
	next := widget.NewButton("Next", nextFunc)
	next.Importance = widget.HighImportance
	buttons := container.NewHBox(copy, close, next)

	bottom := container.NewBorder(nil, nil, cookie, buttons, nil)

	newDisplayWindow.SetContent(container.NewVBox(
		i,
		container.NewPadded(container.NewPadded(fortune)),
		bottom,
	))

	newDisplayWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case "Escape":
			closeFunc()
		case "Return":
			nextFunc()
		}
	})

	newDisplayWindow.Canvas().AddShortcut(&fyne.ShortcutCopy{}, func(shortcut fyne.Shortcut) {
		copyFunc()
	})

	newDisplayWindow.SetCloseIntercept(func() {
		newDisplayWindow.Hide()
	})

	newDisplayWindow.CenterOnScreen()
	newDisplayWindow.SetFixedSize(true)
	newDisplayWindow.Show()

	if displayWindow != nil && !displayWindow.Content().Visible() {
		displayWindow.Close()
	}

	displayWindow = newDisplayWindow
}
