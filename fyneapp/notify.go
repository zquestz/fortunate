package fyneapp

import (
	"fmt"

	"github.com/zquestz/fortunate/config"

	"fyne.io/fyne/v2"
)

// notify displays a desktop notification.
func notify(title string, content string) {
	notification := fyne.NewNotification(title, content)

	appGUI.SendNotification(notification)
}

// notifyError sends an error notification.
func notifyError(err error) {
	notify(fmt.Sprintf("%s Error", config.GUIAppName), err.Error())
}
