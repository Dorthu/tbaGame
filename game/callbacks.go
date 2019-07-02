package game

import (
	"github.com/jroimartin/gocui"
)

func updateScreen(g *gocui.Gui) {
	DrawRoom(g)
	SetText("")
	DrawDialog(g)
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	player.move(-1, 0, FacingUp)
	updateScreen(g)
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	player.move(0, 1, FacingRight)
	updateScreen(g)
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	player.move(1, 0, FacingDown)
	updateScreen(g)
	return nil
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	player.move(0, -1, FacingLeft)
	updateScreen(g)
	return nil
}

func Interact(g *gocui.Gui, v *gocui.View) error {
	player.interact()
	DrawDialog(g)
	return nil
}

func SetRoom(filename string, g *gocui.Gui) {
	room := loadRoom(filename)

	currentRoom = *room
	updateScreen(g)
}
