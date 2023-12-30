package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type BuildingConstruction struct {
	StartDate pserialize.Year `paradox_field:"start_date" json:"start_date,omitempty"`
	Date      pserialize.Year `paradox_field:"date" json:"date,omitempty"`
	Progress  float32         `paradox_field:"progress" json:"progress,omitempty"`
	Days      int             `paradox_field:"days" json:"days,omitempty"`
	Location  int             `paradox_field:"location" json:"location,omitempty"`
	Building  int             `paradox_field:"building" json:"building,omitempty"`
}
