package game

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	default:
		return errors.New(fmt.Sprintf("Invalid color: %s", raw))
	}

	return nil
}

type Space struct {
	Char  string `yaml:char`
	Solid bool   `yaml:solid`
	// description string
	Color colorType `yaml:color`
}

type Room struct {
	Grid [5][9]Space `yaml:grid`
}

/// var testRoom = Room{grid: [5][9]Space{
/// 	{{'+', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'+', true}},
/// 	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
/// 	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
/// 	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
/// 	{{'+', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'+', true}},
/// }}

var testRoom Room

func doSetup() error {
	testRoom.load()
	return nil
}

var _ = doSetup()

func (r *Room) load() *Room {
	raw_yaml, err := ioutil.ReadFile("testRoom.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw_yaml, r)

	if err != nil {
		panic(err)
	}

	return r
}

// var currentRoom Room = testRoom

func GetRoom() *Room {
	return &testRoom
}

/// methods of Room
func (r *Room) spaceFree(loc Point) bool {
	if loc.x < 0 || loc.x > 5 || loc.y < 0 || loc.y > 9 {
		return false
	}

	if r.Grid[loc.x][loc.y].Solid {
		return false
	}

	return true
}
