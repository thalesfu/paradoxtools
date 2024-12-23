package save

type ProvinceInSave struct {
	ID    string    `paradox_type:"map_key" json:"id,omitempty"`
	Owner string    `json:"owner,omitempty"`
	Infra []float32 `paradox_field:"infra" paradox_type:"field_list" json:"infra,omitempty"`
}
