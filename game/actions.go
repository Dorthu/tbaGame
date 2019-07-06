package game

type Action struct {
	Needs       *Item
	NewChar     string
	ChangeSolid bool
	Message     string
	executed    bool
}

/// actually does the thing the action does
func (a *Action) execute(s *Space, p *Player) bool {
	if a.executed {
		return false
	}

	hasItem := false

	for _, i := range p.inventory.Items {
		if i.Name == a.Needs.Name { /// TODO - this should work by comparison
			hasItem = true
		}
	}

	if !hasItem {
		return false
	}

	a.executed = true

	if a.NewChar != "" {
		s.Char = a.NewChar
	}

	if a.ChangeSolid {
		s.Solid = !s.Solid
	}

	if a.Message != "" {
		SetText(a.Message)
	}

	return true
}
