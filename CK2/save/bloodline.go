package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Bloodline struct {
	ID            int                        `paradox_type:"map_key" json:"id,omitempty"`
	Type          string                     `paradox_field:"type" json:"type,omitempty"`
	Flags         map[string]pserialize.Year `paradox_field:"flags" paradox_type:"map_key"  json:"flags,omitempty"`
	Created       pserialize.Year            `paradox_field:"created" json:"created,omitempty"`
	Member        []int                      `paradox_field:"member" paradox_type:"list" json:"member,omitempty"`
	Inheritance   string                     `paradox_field:"inheritance" json:"inheritance,omitempty"`
	AllowBastards pserialize.PBool           `paradox_field:"allow_bastards" json:"allow_bastards,omitempty"`
}
