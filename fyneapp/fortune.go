package fyneapp

import (
	"fmt"
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
)

var fortuneTicker *time.Ticker

// notifyFortune sends a notification containing a fortune.
func notifyFortune() {
	cookie, content, err := fortune.Run()
	if err != nil {
		notifyError(err)
	}

	var title string

	if config.AppConfig.ShowCookie {
		title = fmt.Sprintf("%s %s", config.GUIAppName, cookie)
	} else {
		title = config.GUIAppName
	}

	notify(title, content)
}

// startFortuneTicker starts the fortune notification ticker.
func startFortuneTicker() {
	// Never start panic'd timers.
	if config.AppConfig.FortuneTimer <= 0 {
		return
	}

	fortuneTicker = time.NewTicker(time.Hour * time.Duration(config.AppConfig.FortuneTimer))
	defer fortuneTicker.Stop()

	for {
		select {
		case <-fortuneTicker.C:
			notifyFortune()
		}
	}
}

// fortuneTickerReset changes the timing of the fortune ticker.
func fortuneTickerReset() {
	// Nil means we never started the go routine, start it up.
	if fortuneTicker == nil {
		go startFortuneTicker()
		return
	}

	if config.AppConfig.FortuneTimer <= 0 {
		fortuneTicker.Stop()
		return
	}

	fortuneTicker.Reset(time.Hour * time.Duration(config.AppConfig.FortuneTimer))
}
