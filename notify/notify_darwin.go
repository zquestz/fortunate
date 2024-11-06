//go:build darwin
// +build darwin

package notify

// Notify displays a desktop notification.
func Notify(appName string, title string, text string) error {
	return nil
}
