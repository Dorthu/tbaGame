package game

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Point struct {
	x, y int
}

var playerLoc Point = Point{x: 2, y: 4}

/// move funcs
func move(g *gocui.Gui, deltaX int, deltaY int) {
	newSpace := Point{x: playerLoc.x + deltaX, y: playerLoc.y + deltaY}

	if spaceFree(newSpace) {
		playerLoc.x += deltaX
		playerLoc.y += deltaY
	}

	DrawRoom(g)
}

func spaceFree(loc Point) bool {
	if loc.x < 0 || loc.x > 5 || loc.y < 0 || loc.y > 9 {
		return false
	}

	room := GetRoom()

	if room.grid[loc.x][loc.y].solid {
		return false
	}

	return true
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	move(g, -1, 0)
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	move(g, 0, 1)
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	move(g, 1, 0)
	return nil
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	move(g, 0, -1)
	return nil
}

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
			if playerLoc.x == i && playerLoc.y == j {
				fmt.Fprint(buf, "x")
			} else {
				fmt.Fprint(buf, string(room.grid[i][j].char))
			}
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}
