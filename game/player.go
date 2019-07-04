package game

import "fmt"

/// type definitions
type Point struct {
	x, y int
}

type playerFacing struct {
	x, y int
	char rune
}

func (p *Point) inDirection(f *playerFacing) Point {
	return Point{x: p.x + f.x, y: p.y + f.y}
}

type Player struct {
	loc       Point
	facing    playerFacing
	inventory Inventory
}

/// constants
var FacingUp playerFacing = playerFacing{-1, 0, '^'}
var FacingRight playerFacing = playerFacing{0, 1, '>'}
var FacingDown playerFacing = playerFacing{1, 0, 'v'}
var FacingLeft playerFacing = playerFacing{0, -1, '<'}

/// the player
var player Player = Player{loc: Point{x: 2, y: 4}, facing: FacingUp}

/// methods of Player struct
func (p *Player) move(deltaX int, deltaY int, newFacing playerFacing) {
	newSpace := Point{x: p.loc.x + deltaX, y: p.loc.y + deltaY}

	if GetRoom().spaceFree(newSpace) {
		p.loc.x += deltaX
		p.loc.y += deltaY
	}

	p.facing = newFacing
}

func (p *Player) interact() {
	room := GetRoom()

	space := room.spaceAt(p.loc.inDirection(&p.facing))

	if space == nil {
		return
	}

	item := space.CurItem
	action := space.Action
	if item != nil {
		p.inventory.Items = append(p.inventory.Items, item)
		space.CurItem = nil
		SetText(fmt.Sprintf("Got %s!", item.Name))
	} else if action != nil && action.execute(space, p) {
		/// pass
	} else if space.Description != "" {
		SetText(space.Description)
	}

}
