package game

import (
	"fmt"
	"strconv"
	"strings"
)

type effectType string

const (
	effectLoadRoom = "loadRoom"
)

type Trigger struct {
	Effect effectType
	Data   string
}

func (t *Trigger) fire(p *Player) {
	switch t.Effect {
	case effectLoadRoom:
		doLoadRoom(t, p)
	}
}

func doLoadRoom(t *Trigger, p *Player) {
	data := strings.Split(t.Data, " ")

	if len(data) != 3 {
		panic("Invalid length for loadRoom data!")
	}

	roomName := data[0]
	playerX, err := strconv.Atoi(data[1])

	if err != nil {
		panic(fmt.Sprintf("Invalid value for player X in loadRoom trigger: %s", data[1]))
	}

	playerY, err := strconv.Atoi(data[2])

	if err != nil {
		panic(fmt.Sprintf("Invalid value for player Y in loadRoom trigger: %s", data[2]))
	}

	p.loc.x = playerX
	p.loc.y = playerY
	changeRoom(roomName)
}
