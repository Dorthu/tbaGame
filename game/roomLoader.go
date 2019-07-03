package game

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/// These types are the specs for the actual types that are used in rooms.
type roomSpec struct {
	Version int            `yaml:version`
	Grid    [5][9]gridSpec `yaml:grid`
	Spaces  []spaceSpec    `yaml:spaces`
	Items   []Item         `yaml:items`
}

type gridSpec struct {
	SpaceType string `yaml:"type"`
	Has       string `yaml:has`
}

type spaceSpec struct {
	SpaceType   string              `yaml:"type"`
	Char        string              `yaml:"char"`
	Solid       bool                `yaml:"solid"`
	Color       colorType           `yaml:"color"`
	Description string              `yaml:"description"`
	Background  backgroundColorType `yaml:background`
}

func (s *spaceSpec) makeSpace() Space {
	return Space{
		Char:        s.Char,
		Solid:       s.Solid,
		Color:       s.Color,
		Description: s.Description,
		Background:  s.Background,
	}
}

/// given a filename, reads the file and loads the contents as a roomSpec, then
/// converts that roomSpec into a Room and returns a reference to it
func loadRoom(filename string) *Room {
	var spec roomSpec
	raw_yaml, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw_yaml, &spec)

	if err != nil {
		panic(err)
	}

	r := spec.toRoom()

	return r
}

func (spec *roomSpec) getItem(name string) *Item {
	for _, item := range spec.Items {
		if item.Name == name {
			return &item
		}
	}

	panic(fmt.Sprintf("No item named %s", name))
}

func (spec *roomSpec) toRoom() *Room {
	var r Room

	spaceMap := make(map[string]spaceSpec)

	for _, space := range spec.Spaces {
		spaceMap[space.SpaceType] = space
	}

	for i, col := range spec.Grid {
		for j, cell := range col {
			cellSpec := spaceMap[cell.SpaceType]

			if cellSpec.SpaceType == "" {
				panic(errors.New(fmt.Sprintf("Invalid space type: %s", cell.SpaceType)))
			}

			r.Grid[i][j] = cellSpec.makeSpace()

			if cell.Has != "" {
				heldItem := spec.getItem(cell.Has)

				r.Grid[i][j].CurItem = heldItem
			}
		}
	}

	return &r
}
