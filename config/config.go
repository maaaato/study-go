package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"time"
)

const ConfigPath = "/usr/local/src/monitor.toml"

type ConfToml struct {
	TailFile      string        `toml:"TailFile"`
	PostionFile   string        `toml:"PostionFile"`
	SearchStart   string        `toml:"SearchStart"`
	SearchEnd     string        `toml:"SearchEnd"`
	TagName       string        `toml:"TagName"`
	TagStartValue string        `toml:"TagStartValue"`
	TagEndValue   string        `toml:"TagEndValue"`
	Delay         time.Duration `toml:"Delay"`
}

// Read config.
func New() (*ConfToml, error) {
	c := &ConfToml{}
	_, err := toml.DecodeFile(ConfigPath, c)
	if err != nil {
		panic(err)
	}

	// check error.
	if c.TailFile == "" {
		return c, errors.New("please set tailfile.")
	}

	if c.PostionFile == "" {
		return c, errors.New("please set postion file.")
	}

	if c.SearchStart == "" {
		return c, errors.New("please set search start.")
	}

	if c.SearchEnd == "" {
		return c, errors.New("please set search end.")
	}

	return c, nil
}
