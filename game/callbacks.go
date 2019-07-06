package game

import (
	"github.com/jroimartin/gocui"
)

func updateScreen(g *gocui.Gui) {
	DrawRoom(g)
	DrawDialog(g)
}

func MoveUp(g *gocui.Gui, v *gocui.View) error {
	player.move(-1, 0, FacingUp)
	SetText("")
	updateScreen(g)
	return nil
}

func MoveRight(g *gocui.Gui, v *gocui.View) error {
	player.move(0, 1, FacingRight)
	SetText("")
	updateScreen(g)
	return nil
}

func MoveDown(g *gocui.Gui, v *gocui.View) error {
	player.move(1, 0, FacingDown)
	SetText("")
	updateScreen(g)
	return nil
}

func MoveLeft(g *gocui.Gui, v *gocui.View) error {
	player.move(0, -1, FacingLeft)
	SetText("")
	updateScreen(g)
	return nil
}

func Interact(g *gocui.Gui, v *gocui.View) error {
	player.interact()
	updateScreen(g)
	return nil
}

func SetRoom(filename string, g *gocui.Gui) {
	changeRoom(filename)
	updateScreen(g)
}

func changeRoom(filename string) {
	room := loadRoom("rooms/" + filename)

	currentRoom = *room
}
