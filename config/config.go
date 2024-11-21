package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/zquestz/go-ucl"
)

const (
	AppName    = "fortunate"
	GUIAppName = "Fortunate"
	Version    = "1.1.0"
)

// Config type for application configuration.
type Config struct {
	DisplayVersion bool     `json:"-"`
	IconTheme      string   `json:"iconTheme"`
	FortuneTimer   int      `json:"fortuneTimer,string"`
	ShortFortunes  bool     `json:"shortFortunes,string"`
	LongFortunes   bool     `json:"longFortunes,string"`
	ShowCookie     bool     `json:"showCookie,string"`
	FortuneLists   []string `json:"fortuneLists"`
}

// AppConfig stores the app configuration.
var AppConfig Config

// Load reads the configuration from ~/.config/fortunate/config
// and loads it into the Config struct.
// The config is in UCL format.
func (c *Config) Load() error {
	conf, err := c.loadConfig()
	if err != nil {
		return err
	}

	// There are cases when we don't have a configuration.
	if conf != nil {
		err = c.applyConf(conf)
		if err != nil {
			return err
		}
	}

	return nil
}

// Save saves the configuration to ~/.config/fortunate/config.
func (c *Config) Save() error {
	configurationPath, err := c.configPath()
	if err != nil {
		return err
	}

	os.MkdirAll(configurationPath, 0755)

	var buf bytes.Buffer
	err = ucl.Encode(&buf, AppConfig, "  ", "json", "")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(configurationPath, "config"), buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) loadConfig() ([]byte, error) {
	configurationPath, err := c.configPath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath.Join(configurationPath, "config"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer f.Close()

	ucl.Ucldebug = false
	data, err := ucl.NewParser(f).Ucl()
	if err != nil {
		return nil, err
	}

	conf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *Config) configPath() (string, error) {
	h, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(h, ".config", AppName), nil
}

func (c *Config) applyConf(conf []byte) error {
	err := json.Unmarshal(conf, c)
	if err != nil {
		return err
	}

	return nil
}
