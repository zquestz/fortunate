package fyneapp

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
	"github.com/zquestz/fortunate/notify"
	"github.com/zquestz/fortunate/theme"
)

var settingsWindow fyne.Window

func Settings() {
	// Window was already hidden, just show it.
	if settingsWindow != nil {
		settingsWindow.Show()
		return
	}

	settingsWindow = appGUI.NewWindow("Settings")

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

	shortFortunes := widget.NewCheck("Short", nil)
	shortFortunes.SetChecked(config.AppConfig.ShortFortunes)

	longFortunes := widget.NewCheck("Long", nil)
	longFortunes.SetChecked(config.AppConfig.LongFortunes)

	fortuneLength := container.NewHBox(shortFortunes, longFortunes)

	persistNotifications := widget.NewCheck("", nil)
	persistNotifications.SetChecked(config.AppConfig.PersistNotifications)

	lists, err := fortune.Lists()
	if err != nil {
		notify.NotifyError(err)
	}

	fortuneLists := widget.NewCheckGroup(lists, nil)
	fortuneLists.SetSelected(config.AppConfig.FortuneLists)

	scrollableLists := container.NewScroll(fortuneLists)
	scrollableLists.SetMinSize(fyne.Size{
		Height: 250,
		Width:  300,
	})

	cancelFunc := func() {
		settingsWindow.Hide()
	}

	submitFunc := func() {
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
		config.AppConfig.ShortFortunes = shortFortunes.Checked
		config.AppConfig.LongFortunes = longFortunes.Checked
		config.AppConfig.PersistNotifications = persistNotifications.Checked
		config.AppConfig.FortuneLists = fortuneLists.Selected
		err := config.AppConfig.Save()
		if err != nil {
			notify.NotifyError(err)
		}
		theme.SetIconTheme()
		notify.FortuneTickerReset()
		settingsWindow.Hide()
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Icon Theme", Widget: iconTheme},
			{Text: "Fortune Timer", Widget: notifyFortuneTimes},
			{Text: "Fortune Length", Widget: fortuneLength},
			{Text: "Persist Notifications", Widget: persistNotifications},
			{Text: "Fortune Lists", Widget: scrollableLists},
		},
		OnCancel: cancelFunc,
		OnSubmit: submitFunc,
	}

	settingsWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case "Escape":
			cancelFunc()
		case "Return":
			submitFunc()
		}
	})

	// Hack to make sure we don't accidently quit.
	settingsWindow.SetCloseIntercept(func() {
		settingsWindow.Hide()
	})

	settingsWindow.SetContent(form)
	settingsWindow.CenterOnScreen()
	settingsWindow.SetFixedSize(true)
	settingsWindow.Show()
}
