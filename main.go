package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

type Point struct {
	x, y int
}

var playerLoc Point = Point{x: 2, y: 4}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, showText); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	terrainView(g)
	dialogView(g)
	return nil
}

/// Creates the graphical panel where the player can view the scene
func terrainView(g *gocui.Gui) error {
	screenX, _ := g.Size()
	middleX := screenX / 2
	cornerX := middleX - 5
	/// width = 6, height = 10

	if v, err := g.SetView("terrain", cornerX, 0, cornerX+10, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		drawRoom(g, v)
	}

	return nil
}

/// Creates the view where the text appears
func dialogView(g *gocui.Gui) error {
	offsetY := 7
	screenX, _ := g.Size()
	middleX := screenX / 2
	cornerX := middleX - 20
	/// width = 40, height = 4
	if v, err := g.SetView("dialog", cornerX, offsetY, cornerX+40, offsetY+4); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, "now it does things")
	}

	return nil
}

func showText(g *gocui.Gui, v *gocui.View) error {
	v, err := g.View("dialog")

	if err != nil {
		return err
	}

	v.Clear()
	fmt.Fprint(v, "Changed it")

	return nil
}

/// draws the room
func drawRoom(g *gocui.Gui, v *gocui.View) error {
	buf, err := g.View("terrain")

	if err != nil {
		return err
	}

	buf.Clear()

	roomLines := [5][9]rune{
		{'+', '-', '-', '-', '-', '-', '-', '-', '+'},
		{'|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '|'},
		{'|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '|'},
		{'|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '|'},
		{'+', '-', '-', '-', '-', '-', '-', '-', '+'},
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 9; j++ {
			if playerLoc.x == i && playerLoc.y == j {
				fmt.Fprint(buf, "x")
			} else {
				fmt.Fprint(buf, string(roomLines[i][j]))
			}
		}
		fmt.Fprintln(buf, "")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
