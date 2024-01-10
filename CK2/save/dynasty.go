package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Dynasty struct {
	ID            int              `paradox_type:"map_key" json:"id,omitempty"`
	Name          string           `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	Culture       string           `paradox_field:"culture" json:"culture,omitempty"`
	Religion      string           `paradox_field:"religion" json:"religion,omitempty"`
	CoatOfArms    *CoatOfArms      `paradox_field:"coat_of_arms" json:"coat_of_arms,omitempty"`
	SetCoatOfArms pserialize.PBool `paradox_field:"set_coat_of_arms" json:"set_coat_of_arms,omitempty"`
}
