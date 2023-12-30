package save

type Army struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}
