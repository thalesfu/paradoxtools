package feud

type Kingdom interface {
	Feud

	// GetParent returns the parent of the possession.
	GetParent() Empire

	// GetChildren returns the children of the possession.
	GetChildren() map[string]Duke

	// SetParent sets the parent of the possession.
	SetParent(e Empire)

	// AddChild adds a child to the possession.
	AddChild(d Duke)
}

type BaseKingdom struct {
	Title     string
	TitleName string
	TitleCode string
	Parent    Empire
	Dukes     map[string]Duke
}

func (k *BaseKingdom) GetTitle() string {
	return k.Title
}

func (k *BaseKingdom) GetTitleName() string {
	return k.TitleName
}

func (k *BaseKingdom) GetTitleCode() string {
	return k.TitleCode
}

func (k *BaseKingdom) GetParent() Empire {
	return k.Parent
}

func (k *BaseKingdom) GetChildren() map[string]Duke {
	return k.Dukes
}

func (k *BaseKingdom) SetParent(e Empire) {
	k.Parent = e
}

func (k *BaseKingdom) AddChild(d Duke) {
	k.Dukes[d.GetTitleCode()] = d
	d.SetParent(k)
}
