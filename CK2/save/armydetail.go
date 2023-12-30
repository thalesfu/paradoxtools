package save

type ArmyDetail struct {
	Li []float32 `paradox_field:"li" paradox_type:"field_list" json:"li,omitempty"`
	Hi []float32 `paradox_field:"hi" paradox_type:"field_list" json:"hi,omitempty"`
	Pi []float32 `paradox_field:"pi" paradox_type:"field_list" json:"pi,omitempty"`
	Lc []float32 `paradox_field:"lc" paradox_type:"field_list" json:"lc,omitempty"`
	Ar []float32 `paradox_field:"ar" paradox_type:"field_list" json:"ar,omitempty"`
}
