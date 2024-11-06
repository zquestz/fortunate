package notify

import (
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
)

var fortuneTicker *time.Ticker

func NotifyFortune() {
	output, _ := fortune.Run()

	Notify(config.GUIAppName, config.GUIAppName, output, "")
}

func FortuneTicker() {
	// Never start panic'd timers.
	if config.AppConfig.FortuneTimer <= 0 {
		return
	}

	fortuneTicker = time.NewTicker(time.Hour * time.Duration(config.AppConfig.FortuneTimer))
	defer fortuneTicker.Stop()

	for {
		select {
		case <-fortuneTicker.C:
			NotifyFortune()
		}
	}
}

func FortuneTickerReset() {
	// Nil means we never started the go routine, start it up.
	if fortuneTicker == nil {
		go FortuneTicker()
		return
	}

	if config.AppConfig.FortuneTimer <= 0 {
		fortuneTicker.Stop()
		return
	}

	fortuneTicker.Reset(time.Hour * time.Duration(config.AppConfig.FortuneTimer))
}
