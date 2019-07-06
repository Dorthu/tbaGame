package game

type Space struct {
	Char        string
	Solid       bool
	Description string
	Color       colorType
	Background  backgroundColorType
	CurItem     *Item
	Action      *Action
	Trigger     *Trigger
}

type Room struct {
	Grid [5][9]Space `yaml:grid`
}

var currentRoom Room

func GetRoom() *Room {
	return &currentRoom
}

/// methods of Room
func (r *Room) spaceFree(loc Point) bool {
	if loc.x < 0 || loc.x > 4 || loc.y < 0 || loc.y > 8 {
		return false
	}

	if r.Grid[loc.x][loc.y].Solid {
		return false
	}

	return true
}

func (r *Room) spaceAt(loc Point) *Space {
	if loc.x < 0 || loc.x > 4 || loc.y < 0 || loc.y > 8 {
		return nil
	}

	return &r.Grid[loc.x][loc.y]
}
