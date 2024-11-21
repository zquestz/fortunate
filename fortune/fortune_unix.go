//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd

package fortune

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/zquestz/fortunate/config"

	"github.com/Masterminds/semver/v3"
)

var (
	listMatcher        = regexp.MustCompile(`\s+(\d+\.\d+\%)\s([\w|-]+)`)
	cookieMatcher      = regexp.MustCompile(`\((.*)\)`)
	CookieSupported    = true
	fortuneBinary      = "fortune"
	localFortuneBinary = "/usr/local/bin/fortune"
)

// Run runs fortune.
func Run() (string, string, error) {
	args := []string{}

	if cookieEnabled() {
		args = append(args, "-c")
	}

	if config.AppConfig.ShortFortunes {
		args = append(args, "-s")
	}

	if config.AppConfig.LongFortunes {
		args = append(args, "-l")
	}

	if len(config.AppConfig.FortuneLists) > 0 {
		args = append(args, config.AppConfig.FortuneLists...)
	}

	cmd := exec.Command(fortuneBinary, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	if cookieEnabled() {
		splitOutput := strings.Split(string(output), "\n")
		if len(splitOutput) < 3 {
			return "", "", errors.New("failed to parse fortune output")
		}

		cookie := shortenCookie(splitOutput[0])
		content := strings.Join(splitOutput[2:], "\n")

		return cookie, content, nil
	}

	return "", string(output), nil
}

// Lists returns a list of installed fortune lists.
func Lists() ([]string, error) {
	cmd := exec.Command(fortuneBinary, "-f")
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

	sort.Strings(list)

	return list, nil
}

// CheckCookieSupported checks if the locally installed fortune supports cookie display.
func CheckCookieSupported() error {
	err := checkFortuneExists()
	if err != nil {
		CookieSupported = false
		return err
	}

	cmd := exec.Command(fortuneBinary, "-v")
	output, err := cmd.Output()
	if err != nil {
		CookieSupported = false
		return err
	}

	splitOutput := strings.Split(string(output), " ")
	if len(splitOutput) < 3 {
		CookieSupported = false
		return nil
	}

	version := strings.TrimSpace(splitOutput[2])

	versionConstraint, err := semver.NewConstraint(">= 2.22.0, < 1000")
	if err != nil {
		CookieSupported = false
		return err
	}

	semVer, err := semver.NewVersion(version)
	if err != nil {
		CookieSupported = false
		return err
	}

	CookieSupported = versionConstraint.Check(semVer)

	return nil
}

// checkFortuneExists checks for the existence of fortune.
func checkFortuneExists() error {
	cmd := exec.Command(fortuneBinary, "-v")
	_, err := cmd.Output()
	if err != nil {
		_, statErr := os.Stat(localFortuneBinary)
		if statErr == nil {
			fortuneBinary = localFortuneBinary
			return nil
		}
	}

	return err
}

// cookieEnabled returns if cookie support should be enabled.
func cookieEnabled() bool {
	if config.AppConfig.ShowCookie && CookieSupported {
		return true
	}

	return false
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
