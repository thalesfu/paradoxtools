package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Religion struct {
	ID                          string           `paradox_type:"map_key" json:"id,omitempty"`
	Parent                      string           `paradox_field:"parent" json:"parent,omitempty"`
	OriginalParent              string           `paradox_field:"original_parent" json:"original_parent,omitempty"`
	WasHeresy                   pserialize.PBool `paradox_field:"was_heresy" json:"was_heresy,omitempty"`
	AllowVikingInvasionOverride pserialize.PBool `paradox_field:"allow_viking_invasion_override" json:"allow_viking_invasion_override,omitempty"`
	Reformed                    string           `paradox_field:"reformed" json:"reformed,omitempty"`
	OriginalReformed            string           `paradox_field:"original_reformed" json:"original_reformed,omitempty"`
	HolySites                   []string         `paradox_field:"holy_sites" paradox_type:"field_list" json:"holy_sites,omitempty"`
	Authorities                 []*Authority     `paradox_field:"authority" paradox_type:"list" json:"authority,omitempty"`
}

type Authority struct {
	Date     pserialize.Year `paradox_field:"d" json:"date,omitempty"`
	Modifier string          `paradox_field:"modifier" json:"modifier,omitempty"`
}
