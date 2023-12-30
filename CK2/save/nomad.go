package save

type Nomad struct {
	Name       string      ` paradox_type:"map_key" json:"name,omitempty"`
	Population int         `paradox_field:"population" json:"population,omitempty"`
	Manpower   int         `paradox_field:"manpower" json:"manpower,omitempty"`
	Capital    string      `paradox_field:"capital" json:"capital,omitempty"`
	Technology *Technology `paradox_field:"technology" json:"technology,omitempty"`
}
