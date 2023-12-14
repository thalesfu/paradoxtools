package save

type SaveFile struct {
	Version      string          `paradox_field:"version" paradox_type:"field" json:"version"`
	DynastyTitle []*DynastyTitle `paradox_field:"dyn_title" paradox_type:"list" json:"dyn_title"`
}

type DynastyTitle struct {
	Title string `paradox_field:"title" paradox_type:"field" json:"title"`
}
