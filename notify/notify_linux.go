//go:build linux
// +build linux

package notify

import (
	"os/exec"

	"github.com/zquestz/fortunate/config"
)

// Notify displays a desktop notification.
func Notify(appName string, title string, text string, iconPath string) {
	cmd := exec.Command("notify-send", "-e", "-a", config.GUIAppName, "-i", config.AppName, title, text)
	cmd.Run()
}
