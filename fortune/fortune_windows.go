//go:build windows
// +build windows

package fortune

import "errors"

// Run runs fortune.
func Run() (string, error) {
	return "", errors.New("not supported")
}
