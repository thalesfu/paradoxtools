package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Disease struct {
	First         pserialize.PBool `paradox_field:"first" json:"first,omitempty"`
	Months        int              `paradox_field:"months" json:"months,omitempty"`
	TimePeriod    int              `paradox_field:"timeperiod" json:"timeperiod,omitempty"`
	InfectionDate pserialize.Year  `paradox_field:"infection_date" json:"infection_date,omitempty"`
}
