package save

type CoatOfArms struct {
	Data     []int  `paradox_field:"data" paradox_type:"field_list" json:"data,omitempty"`
	Religion string `paradox_field:"religion" json:"religion,omitempty"`
}
