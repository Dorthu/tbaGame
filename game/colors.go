package game

import (
	"errors"
	"fmt"
)

type colorType string

const (
	ColorRed    = "\u001b[31m"
	ColorBlue   = "\u001b[34m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorWhite  = "\u001b[37m"
	ColorReset  = "\u001b[0m"
)

type backgroundColorType string

const (
	BackgroundRed    = "\u001b[41m"
	BackgroundGreen  = "\u001b[42m"
	BackgroundYellow = "\u001b[43m"
	BackgroundBlue   = "\u001b[44m"
	BackgroundReset  = "\u001b[0m"
)

func (c *colorType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string
	if err := unmarshal(&raw); err != nil {
		return err
	}
	color := colorType(raw)

	switch color {
	case "white":
		*c = ColorWhite
	case "red":
		*c = ColorRed
	case "green":
		*c = ColorGreen
	case "blue":
		*c = ColorBlue
	case "yellow":
		*c = ColorYellow
	default:
		return errors.New(fmt.Sprintf("Invalid color: %s", raw))
	}

	return nil
}

func (c *backgroundColorType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string
	if err := unmarshal(&raw); err != nil {
		return err
	}
	color := backgroundColorType(raw)

	switch color {
	case "red":
		*c = BackgroundRed
	case "green":
		*c = BackgroundGreen
	case "yellow":
		*c = BackgroundYellow
	case "blue":
		*c = BackgroundBlue
	default:
		return errors.New(fmt.Sprintf("Invalid background color: %s", raw))
	}

	return nil
}
