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
				fmt.Fprint(buf, "\u001b[31m"+string(player.facing.char)+"\u001b[0m")
			} else {
				/// TODO: use the Space's color
				fmt.Fprint(buf, "\u001b[32m"+string(room.Grid[i][j].Char)+"\u001b[0m")
			}
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}
