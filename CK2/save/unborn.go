package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Unborn struct {
	Mother      *SimplePeople    `paradox_field:"mother" json:"mother,omitempty"`
	Father      *SimplePeople    `paradox_field:"father" json:"father,omitempty"`
	Date        pserialize.Year  `paradox_field:"date" json:"date,omitempty"`
	IsBastard   pserialize.PBool `paradox_field:"is_bastard" json:"is_bastard,omitempty"`
	KnownFather pserialize.PBool `paradox_field:"known_father" json:"known_father,omitempty"`
	Message     pserialize.PBool `paradox_field:"message" json:"message,omitempty"`
}
