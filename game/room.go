package game

type Space struct {
	char  rune
	solid bool
}

type Room struct {
	grid [5][9]Space
}

var testRoom = Room{grid: [5][9]Space{
	{{'+', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'+', true}},
	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
	{{'|', true}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {' ', false}, {'|', true}},
	{{'+', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'-', true}, {'+', true}},
}}

var currentRoom Room = testRoom

func GetRoom() *Room {
	return &currentRoom
}

/// methods of Room
func (r *Room) spaceFree(loc Point) bool {
	if loc.x < 0 || loc.x > 5 || loc.y < 0 || loc.y > 9 {
		return false
	}

	if r.grid[loc.x][loc.y].solid {
		return false
	}

	return true
}
