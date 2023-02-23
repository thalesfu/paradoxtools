package feud

type Feud interface {
	// GetTitle returns the title of the possession.
	GetTitle() string

	// GetTitleName returns the title name of the possession.
	GetTitleName() string

	// GetTitleCode returns the title code of the possession.
	GetTitleCode() string
}
