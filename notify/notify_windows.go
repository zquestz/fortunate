//go:build windows

package notify

import "errors"

// Notify displays a desktop notification.
func Notify(appName string, title string, text string) error {
	return errors.New("not supported")
}
