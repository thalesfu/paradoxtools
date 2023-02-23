package feud

type Barony interface {
	Feud

	// GetParent returns the parent of the possession.
	GetParent() County

	// SetParent sets the parent of the possession.
	SetParent(c County)
}

type BaseBarony struct {
	Title     string
	TitleName string
	TitleCode string
	Parent    County
}

func (b *BaseBarony) GetTitle() string {
	return b.Title
}

func (b *BaseBarony) GetTitleName() string {
	return b.TitleName
}

func (b *BaseBarony) GetTitleCode() string {
	return b.TitleCode
}

func (b *BaseBarony) GetParent() County {
	return b.Parent
}

func (b *BaseBarony) SetParent(c County) {
	b.Parent = c
}
