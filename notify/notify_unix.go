//go:build linux || dragonfly || freebsd || netbsd || openbsd
// +build linux dragonfly freebsd netbsd openbsd

package notify

import (
	"os/exec"

	"github.com/zquestz/fortunate/config"
)

// Notify displays a desktop notification.
func Notify(appName string, title string, text string) error {
	cmd := exec.Command("notify-send", "-e", "-a", config.GUIAppName, "-i", config.AppName, title, text)
	return cmd.Run()
}
