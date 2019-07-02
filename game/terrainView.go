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
				fmt.Fprint(buf, string(ColorRed)+string(player.facing.char)+string(ColorReset))
			} else {
				space := room.Grid[i][j]
				if space.CurItem != nil {
					item := space.CurItem
					fmt.Fprint(buf, string(item.Color)+string(item.Overworld)+string(ColorReset))
				} else {
					fmt.Fprint(buf, string(space.Color)+string(space.Char)+string(ColorReset))
				}
			}
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}
