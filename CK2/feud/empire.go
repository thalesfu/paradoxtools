package feud

type Empire interface {
	Feud

	// GetChildren returns the children of the possession.
	GetChildren() map[string]Kingdom

	// AddChild adds a child to the possession.
	AddChild(k Kingdom)
}

type BaseEmpire struct {
	Title     string
	TitleName string
	TitleCode string
	Kingdoms  map[string]Kingdom
}

func (e *BaseEmpire) GetTitle() string {
	return e.Title
}

func (e *BaseEmpire) GetTitleName() string {
	return e.TitleName
}

func (e *BaseEmpire) GetTitleCode() string {
	return e.TitleCode
}

func (e *BaseEmpire) GetChildren() map[string]Kingdom {
	return e.Kingdoms
}

func (e *BaseEmpire) AddChild(k Kingdom) {
	e.Kingdoms[k.GetTitleCode()] = k
	k.SetParent(e)
}
