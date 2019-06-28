package game

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

/// draws the room
func DrawRoom(g *gocui.Gui) error {
	buf, err := g.View("terrain")

	if err != nil {
		return err
	}

	buf.Clear()

	room := GetRoom()

	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if player.loc.x == i && player.loc.y == j {
				fmt.Fprint(buf, string(player.facing.char))
			} else {
				fmt.Fprint(buf, string(room.grid[i][j].char))
			}
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}
