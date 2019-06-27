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

func GetRoom() *Room {
	return &testRoom
}
