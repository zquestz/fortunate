package notify

import (
	"fmt"
	"time"

	"github.com/zquestz/fortunate/config"
	"github.com/zquestz/fortunate/fortune"
)

var fortuneTicker *time.Ticker

func NotifyFortune() error {
	output, err := fortune.Run()
	if err != nil {
		return err
	}

	return Notify(config.GUIAppName, config.GUIAppName, output)
}

func NotifyError(err error) error {
	return Notify(config.GUIAppName, fmt.Sprintf("%s Error", config.GUIAppName), err.Error())
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
			err := NotifyFortune()
			if err != nil {
				NotifyError(err)
			}
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
