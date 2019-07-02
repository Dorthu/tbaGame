package game

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var currentText string

func SetText(s string) {
	currentText = s
}

/// updating display funcs
func DrawDialog(g *gocui.Gui) error {
	v, err := g.View("dialog")

	if err != nil {
		return err
	}

	v.Clear()
	fmt.Fprint(v, currentText)

	return nil
}
