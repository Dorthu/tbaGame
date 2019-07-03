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
			space := room.Grid[i][j]
			/// print the background color code
			fmt.Fprint(buf, string(space.Background))

			/// print whatever's in the space
			if player.loc.x == i && player.loc.y == j {
				fmt.Fprint(buf, string(ColorRed)+string(player.facing.char)+string(ColorReset))
			} else {
				if space.CurItem != nil {
					item := space.CurItem
					fmt.Fprint(buf, string(item.Color)+string(item.Overworld)+string(ColorReset))
				} else {
					fmt.Fprint(buf, string(space.Color)+string(space.Char))
				}
			}
			/// reset the color
			fmt.Fprint(buf, string(ColorReset))
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}
