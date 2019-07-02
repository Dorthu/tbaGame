package game

type Item struct {
	Name        string `yaml:name`
	Char        string `yaml:char`
	Description string `yaml:description`
}

type Inventory struct {
	Items []Item
}
