package save

type Nomination struct {
	Voter   *NominationPerson `paradox_field:"voter" json:"voter,omitempty"`
	Nominee *NominationPerson `paradox_field:"nominee" json:"nominee,omitempty"`
}

type NominationPerson struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}
