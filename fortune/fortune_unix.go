//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd

package fortune

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/zquestz/fortunate/config"
)

var (
	listMatcher   = regexp.MustCompile(`\s+(\d+\.\d+\%)\s([\w|-]+)`)
	cookieMatcher = regexp.MustCompile(`\((.*)\)`)
)

// Run runs fortune.
func Run() (string, string, error) {
	args := []string{"-c"}

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
	if err != nil {
		return "", "", err
	}

	splitOutput := strings.Split(string(output), "\n")
	if len(splitOutput) < 3 {
		return "", "", errors.New("failed to parse fortune output")
	}

	cookie := shortenCookie(splitOutput[0])
	content := strings.Join(splitOutput[2:], "\n")

	return cookie, content, nil
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

// shortenCookie removes the path from the cookie.
func shortenCookie(cookie string) string {
	matches := cookieMatcher.FindStringSubmatch(cookie)
	if len(matches) < 2 {
		return cookie
	}

	cookieSlice := strings.Split(matches[1], string(filepath.Separator))
	shortenedCookie := cookieSlice[len(cookieSlice)-1]

	return fmt.Sprintf("(%s)", shortenedCookie)
}
