package fyneapp

import (
	"strings"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var settingsWindow fyne.Window

// settings creates/shows the settings window.
func settings() {
	if settingsWindow != nil && settingsWindow.Content().Visible() {
		settingsWindow.Show()
		return
	}

	newSettingsWindow := appGUI.NewWindow("Settings")

	ctrlW := &desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault}
	newSettingsWindow.Canvas().AddShortcut(ctrlW, func(shortcut fyne.Shortcut) {
		newSettingsWindow.Hide()
	})

	iconTheme := buildIconTheme()
	notifyFortuneTimes := buildNotifyFortuneTimes()

	shortFortunes := widget.NewCheck("Short", nil)
	shortFortunes.SetChecked(config.AppConfig.ShortFortunes)

	longFortunes := widget.NewCheck("Long", nil)
	longFortunes.SetChecked(config.AppConfig.LongFortunes)

	fortuneLength := container.NewHBox(shortFortunes, longFortunes)

	showCookie := widget.NewCheck("", nil)

	if !fortune.CookieSupported {
		config.AppConfig.ShowCookie = false
	}

	showCookie.SetChecked(config.AppConfig.ShowCookie)

	lists, err := fortune.Lists()
	if err != nil {
		notifyError(err)
	}

	fortuneLists := widget.NewCheckGroup(lists, nil)
	fortuneLists.SetSelected(config.AppConfig.FortuneLists)

	scrollableLists := container.NewScroll(fortuneLists)
	scrollableLists.SetMinSize(fyne.Size{
		Height: 250,
		Width:  300,
	})

	cancelFunc := func() {
		newSettingsWindow.Hide()
	}

	submitFunc := func() {
		config.AppConfig.IconTheme = strings.ToLower(iconTheme.Selected)
		setNotifyFortuneTimes(notifyFortuneTimes)
		config.AppConfig.ShortFortunes = shortFortunes.Checked
		config.AppConfig.LongFortunes = longFortunes.Checked
		config.AppConfig.ShowCookie = showCookie.Checked
		config.AppConfig.FortuneLists = fortuneLists.Selected
		err := config.AppConfig.Save()
		if err != nil {
			notifyError(err)
		}
		setSystrayIcon()
		fortuneTickerReset()
		newSettingsWindow.Hide()
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Icon Theme", Widget: iconTheme},
			{Text: "Fortune Timer", Widget: notifyFortuneTimes},
			{Text: "Fortune Length", Widget: fortuneLength},
		},
		OnCancel: cancelFunc,
		OnSubmit: submitFunc,
	}

	if fortune.CookieSupported {
		form.Items = append(form.Items, &widget.FormItem{Text: "Show Cookie", Widget: showCookie})
	}

	form.Items = append(form.Items, &widget.FormItem{Text: "Fortune Lists", Widget: scrollableLists})

	newSettingsWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case "Escape":
			cancelFunc()
		case "Return":
			submitFunc()
		}
	})

	newSettingsWindow.SetCloseIntercept(func() {
		newSettingsWindow.Hide()
	})

	newSettingsWindow.SetContent(form)
	newSettingsWindow.SetFixedSize(true)
	newSettingsWindow.CenterOnScreen()
	newSettingsWindow.Show()

	// Destroy old window if it exists.
	if settingsWindow != nil && !settingsWindow.Content().Visible() {
		settingsWindow.Close()
	}

	settingsWindow = newSettingsWindow
}

func buildIconTheme() *widget.Select {
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

	return iconTheme
}

func buildNotifyFortuneTimes() *widget.Select {
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

	return notifyFortuneTimes
}

func setNotifyFortuneTimes(notifyFortuneTimes *widget.Select) {
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
}
