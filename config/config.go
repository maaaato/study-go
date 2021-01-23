package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

const ConfigPath = "monitor.toml"

// ConfigToml Struct
type ConfToml struct {
	TailFile      string        `toml:"TailFile"`
	PositionFile  string        `toml:"PositionFile"`
	SearchStart   string        `toml:"SearchStart"`
	SearchEnd     string        `toml:"SearchEnd"`
	TagName       string        `toml:"TagName"`
	TagStartValue string        `toml:"TagStartValue"`
	TagEndValue   string        `toml:"TagEndValue"`
	Delay         time.Duration `toml:"Delay"`
}

// New Read config.
func New() (*ConfToml, error) {
	var c ConfToml
	_, err := toml.DecodeFile(ConfigPath, &c)
	if err != nil {
		fmt.Println(err)
	}

	// check error.
	if c.TailFile == "" {
		return &c, errors.New("please set tailfile.")
	}

	if c.PositionFile == "" {
		return &c, errors.New("please set postion file.")
	}

	if c.SearchStart == "" {
		return &c, errors.New("please set search start.")
	}

	if c.SearchEnd == "" {
		return &c, errors.New("please set search end.")
	}

	return &c, nil
}
