//go:build windows

package fortune

import "errors"

// Run runs fortune.
func Run() (string, string, error) {
	return "", "", errors.New("not supported")
}

// Lists returns a list of installed fortune lists.
func Lists() ([]string, error) {
	return []string{}, errors.New("not supported")
}
