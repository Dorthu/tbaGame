package game

import (
	"github.com/jroimartin/gocui"
)

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	player.move(-1, 0, FacingUp)
	DrawRoom(g)
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	player.move(0, 1, FacingRight)
	DrawRoom(g)
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	player.move(1, 0, FacingDown)
	DrawRoom(g)
	return nil
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	player.move(0, -1, FacingLeft)
	DrawRoom(g)
	return nil
}
