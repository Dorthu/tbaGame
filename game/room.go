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

type Space struct {
	Char        string
	Solid       bool
	Description string
	Color       colorType
	CurItem     *Item
}

type Room struct {
	Grid [5][9]Space `yaml:grid`
}

var currentRoom Room

func GetRoom() *Room {
	return &currentRoom
}

/// methods of Room
func (r *Room) spaceFree(loc Point) bool {
	if loc.x < 0 || loc.x > 4 || loc.y < 0 || loc.y > 8 {
		return false
	}

	if r.Grid[loc.x][loc.y].Solid {
		return false
	}

	return true
}

func (r *Room) spaceAt(loc Point) *Space {
	if loc.x < 0 || loc.x > 4 || loc.y < 0 || loc.y > 8 {
		return nil
	}

	return &r.Grid[loc.x][loc.y]
}
