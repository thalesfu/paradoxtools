package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type LawVote struct {
	Voter int    `paradox_field:"voter" json:"voter,omitempty"`
	Law   string `paradox_field:"law" json:"law,omitempty"`
}

type RejectedLawChanges struct {
	Actor int             `paradox_field:"actor" json:"actor,omitempty"`
	Law   string          `paradox_field:"law" json:"law,omitempty"`
	Date  pserialize.Year `paradox_field:"date" json:"date,omitempty"`
}
