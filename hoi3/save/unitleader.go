package save

type UnitLeader struct {
	ID   string `paradox_field:"id" json:"id,omitempty"`
	Type string `paradox_field:"type" json:"type,omitempty"`
}
