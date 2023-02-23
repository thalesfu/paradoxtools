package feud

type County interface {
	Feud

	// GetID returns the id of the possession.
	GetID() string

	// GetParent returns the parent of the possession.
	GetParent() Duke

	// GetChildren returns the children of the possession.
	GetChildren() map[string]Barony

	// SetParent sets the parent of the possession.
	SetParent(d Duke)

	// AddChild adds a child to the possession.
	AddChild(b Barony)
}

type BaseCounty struct {
	ID        string
	Title     string
	TitleName string
	TitleCode string
	Parent    Duke
	Baronies  map[string]Barony
}

func (c *BaseCounty) GetID() string {
	return c.ID
}

func (c *BaseCounty) GetTitle() string {
	return c.Title
}

func (c *BaseCounty) GetTitleName() string {
	return c.TitleName
}

func (c *BaseCounty) GetTitleCode() string {
	return c.TitleCode
}

func (c *BaseCounty) GetParent() Duke {
	return c.Parent
}

func (c *BaseCounty) GetChildren() map[string]Barony {
	return c.Baronies
}

func (c *BaseCounty) SetParent(d Duke) {
	c.Parent = d
}

func (c *BaseCounty) AddChild(b Barony) {
	c.Baronies[b.GetTitleCode()] = b
	b.SetParent(c)
}
