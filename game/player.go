package game

/// type definitions
type Point struct {
	x, y int
}

type playerFacing struct {
	x, y int
	char rune
}

type Player struct {
	loc    Point
	facing playerFacing
}

/// constants
var FacingUp playerFacing = playerFacing{-1, 0, '^'}
var FacingRight playerFacing = playerFacing{0, -1, '>'}
var FacingDown playerFacing = playerFacing{1, 0, 'v'}
var FacingLeft playerFacing = playerFacing{0, 1, '<'}

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
