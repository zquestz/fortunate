//go:build windows

package fortune

import (
	"errors"
)

var (
	CookieSupported = false
)

// Run runs fortune.
func Run() (string, string, error) {
	return "", "", errors.New("not supported")
}

// Lists returns a list of installed fortune lists.
func Lists() ([]string, error) {
	return []string{}, errors.New("not supported")
}

// CheckCookieSupported checks if the locally installed fortune supports cookie display.
func CheckCookieSupported() error {
	return errors.New("not supported")
}
