//go:build linux || dragonfly || freebsd || netbsd || openbsd

package notify

import (
	"os/exec"

	"github.com/zquestz/fortunate/config"
)

// Notify displays a desktop notification.
func Notify(appName string, title string, text string) error {
	cmdArgs := []string{}
	if !config.AppConfig.PersistNotifications {
		cmdArgs = append(cmdArgs, "-e")
	}
	cmdArgs = append(cmdArgs, "-a", config.GUIAppName, "-i", config.AppName, title, text)
	cmd := exec.Command("notify-send", cmdArgs...)
	return cmd.Run()
}
