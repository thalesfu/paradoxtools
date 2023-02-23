package feud

type Duke interface {
	Feud

	// GetParent returns the parent of the possession.
	GetParent() Kingdom

	// GetChildren returns the children of the possession.
	GetChildren() map[string]County

	// SetParent sets the parent of the possession.
	SetParent(k Kingdom)

	// AddChild adds a child to the possession.
	AddChild(c County)
}

type BaseDuke struct {
	Title     string
	TitleName string
	TitleCode string
	Parent    Kingdom
	Counties  map[string]County
}

func (d *BaseDuke) GetTitle() string {
	return d.Title
}

func (d *BaseDuke) GetTitleName() string {
	return d.TitleName
}

func (d *BaseDuke) GetTitleCode() string {
	return d.TitleCode
}

func (d *BaseDuke) GetParent() Kingdom {
	return d.Parent
}

func (d *BaseDuke) GetChildren() map[string]County {
	return d.Counties
}

func (d *BaseDuke) SetParent(k Kingdom) {
	d.Parent = k
}

func (d *BaseDuke) AddChild(c County) {
	d.Counties[c.GetTitleCode()] = c
	c.SetParent(d)
}
