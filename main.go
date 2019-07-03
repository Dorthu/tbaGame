package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"

	"github.com/dorthu/tbaGame/game"
)

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

	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, game.Interact); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, game.MoveUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, game.MoveRight); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, game.MoveDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, game.MoveLeft); err != nil {
		log.Panicln(err)
	}

	/// let's get it started
	game.SetRoom("rooms/outside.yaml", g)

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

	if _, err := g.SetView("terrain", cornerX, 0, cornerX+10, 6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		game.DrawRoom(g)
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
		fmt.Fprintln(v, "Arrow keys: Move.  Space: interact.")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
