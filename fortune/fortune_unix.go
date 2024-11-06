//go:build linux || darwin
// +build linux darwin

package fortune

import (
	"os/exec"
	"strings"
)

// Run runs fortune.
func Run() (string, error) {
	cmd := exec.Command("fortune")
	output, err := cmd.Output()

	return strings.TrimSpace(string(output)), err
}
