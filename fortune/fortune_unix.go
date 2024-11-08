//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd

package fortune

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"

	"github.com/zquestz/fortunate/config"
)

var (
	listMatcher = regexp.MustCompile(`\s+(\d+\.\d+\%)\s([\w|-]+)`)
)

// Run runs fortune.
func Run() (string, error) {
	args := []string{}

	if config.AppConfig.ShortFortunes {
		args = append(args, "-s")
	}

	if config.AppConfig.LongFortunes {
		args = append(args, "-l")
	}

	if len(config.AppConfig.FortuneLists) > 0 {
		args = append(args, config.AppConfig.FortuneLists...)
	}

	cmd := exec.Command("fortune", args...)
	output, err := cmd.Output()

	return strings.TrimSpace(string(output)), err
}

// Lists returns a list of installed fortune lists.
func Lists() ([]string, error) {
	cmd := exec.Command("fortune", "-f")
	lists, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	list := []string{}

	scanner := bufio.NewScanner(strings.NewReader(string(lists)))
	for scanner.Scan() {
		row := scanner.Text()
		matches := listMatcher.FindStringSubmatch(row)
		if len(matches) < 3 {
			continue
		}

		list = append(list, matches[2])
	}

	return list, nil
}
