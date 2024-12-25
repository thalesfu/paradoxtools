package save

type ProducingInSave struct {
	Energy        float64 `paradox_field:"energy" json:"energy,omitempty"`
	Metal         float64 `paradox_field:"metal" json:"metal,omitempty"`
	RareMaterials float64 `paradox_field:"rare_materials" json:"rare_materials,omitempty"`
	CrudeOil      float64 `paradox_field:"crude_oil" json:"crude_oil,omitempty"`
}
