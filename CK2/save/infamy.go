package save

type Infamy struct {
	TotalInfamy float32          `paradox_field:"total_infamy" json:"total_infamy,omitempty"`
	Histories   []*InfamyHistory `paradox_field:"history" paradox_type:"list" json:"history,omitempty"`
}

type InfamyHistory struct {
	Name  string  `paradox_field:"name" json:"name,omitempty"`
	Value float32 `paradox_field:"value" json:"value,omitempty"`
}
