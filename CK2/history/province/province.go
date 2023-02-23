package province

type Province interface {
	// GetID returns the ID of the province.
	GetID() string

	// GetTitle returns the title of the province.
	GetTitleName() string

	// GetTitleCode returns the title code of the province.
	GetTitleCode() string
}

type BaseProvince struct {
	ID        string
	Title     string
	TitleCode string
}

func (p *BaseProvince) GetID() string {
	return p.ID
}

func (p *BaseProvince) GetTitleName() string {
	return p.Title
}

func (p *BaseProvince) GetTitleCode() string {
	return p.TitleCode
}
