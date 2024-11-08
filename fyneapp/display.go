package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
	"github.com/zquestz/fortunate/icon"
	"github.com/zquestz/fortunate/notify"
)

var displayWindow fyne.Window

func Display() {
	// Window was is showing, bring to foreground.
	if displayWindow != nil && displayWindow.Content().Visible() {
		displayWindow.Show()
		return
	}

	var newDisplayWindow fyne.Window

	newDisplayWindow = appGUI.NewWindow(config.GUIAppName)

	i := canvas.NewImageFromResource(&fyne.StaticResource{
		StaticName:    config.GUIAppName,
		StaticContent: icon.Data,
	})
	i.FillMode = canvas.ImageFillOriginal
	newFortune, err := fortune.Run()
	if err != nil {
		notify.NotifyError(err)
	}
	fortune := widget.NewLabel(newFortune)
	fortune.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	closeFunc := func() {
		newDisplayWindow.Hide()
	}

	nextFunc := func() {
		newDisplayWindow.Hide()
		Display()
	}

	close := widget.NewButton("Close", closeFunc)
	next := widget.NewButton("Next", nextFunc)
	next.Importance = widget.HighImportance
	buttons := container.NewHBox(close, next)
	rightBorder := container.NewBorder(nil, nil, nil, buttons, nil)

	newDisplayWindow.SetContent(container.NewVBox(
		i,
		fortune,
		rightBorder,
	))

	newDisplayWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case "Escape":
			closeFunc()
		case "Return":
			nextFunc()
		}
	})

	// Hack to make sure we don't accidently quit.
	newDisplayWindow.SetCloseIntercept(func() {
		newDisplayWindow.Hide()
	})

	newDisplayWindow.CenterOnScreen()
	newDisplayWindow.SetFixedSize(true)
	newDisplayWindow.Show()

	// Destroy old window if it exists.
	if displayWindow != nil && !displayWindow.Content().Visible() {
		displayWindow.Close()
	}

	displayWindow = newDisplayWindow
}
