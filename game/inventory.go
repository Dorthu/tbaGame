package game

type Item struct {
	Name        string    `yaml:name`
	Image       string    `yaml:char`
	Description string    `yaml:description`
	Overworld   string    `yaml:overworld`
	Color       colorType `yaml:color`
}

type Inventory struct {
	Items []*Item
}
