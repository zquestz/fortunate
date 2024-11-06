package theme

import (
	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/icon"

	"fyne.io/systray"
)

func SetIconTheme() {
	switch iconTheme := config.AppConfig.IconTheme; iconTheme {
	case "dark":
		systray.SetTemplateIcon(icon.DataDark, icon.DataDark)
	case "light":
		systray.SetTemplateIcon(icon.DataLight, icon.DataLight)
	default:
		systray.SetTemplateIcon(icon.Data, icon.Data)
	}
}
