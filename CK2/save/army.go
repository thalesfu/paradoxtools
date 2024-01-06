package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type ArmyID struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}

type Army struct {
	ArmyID      *ArmyID    `paradox_field:"id" json:"id,omitempty"`
	SubUnits    []*SubUnit `paradox_field:"sub_unit" paradox_type:"list" json:"sub_unit,omitempty"`
	Name        string     `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	Previous    int        `paradox_field:"previous" json:"previous,omitempty"`
	Location    int        `paradox_field:"location" json:"location,omitempty"`
	Flank       []int      `paradox_field:"flank" paradox_type:"field_list" json:"flank,omitempty"`
	FlankLeft   int        `paradox_field:"flank_left" json:"flank_left,omitempty"`
	FlankCenter int        `paradox_field:"flank_center" json:"flank_center,omitempty"`
	FlankRight  int        `paradox_field:"flank_right" json:"flank_right,omitempty"`
	Attachments []*ArmyID  `paradox_field:"attachments" paradox_type:"field_list" json:"attachments,omitempty"`
	Base        int        `paradox_field:"base" json:"base,omitempty"`
}

type SubUnit struct {
	ArmyID *ArmyID         `paradox_field:"id" json:"id,omitempty"`
	Title  *Title          `paradox_field:"title" paradox_type:"entity" paradox_default_field:"id" json:"title,omitempty"`
	Home   int             `paradox_field:"home" json:"home,omitempty"`
	Type   int             `paradox_field:"type" json:"type,omitempty"`
	Leader int             `paradox_field:"leader" json:"leader,omitempty"`
	Owner  int             `paradox_field:"owner" json:"owner,omitempty"`
	Date   pserialize.Year `paradox_field:"date" json:"date,omitempty"`
	Troops *ArmyDetail     `paradox_field:"troops" json:"troops,omitempty"`
}
