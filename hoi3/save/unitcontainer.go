package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type UnitContainer struct {
	ID            *UnitID          `paradox_field:"id" json:"id,omitempty"`
	Name          string           `paradox_field:"name" json:"name,omitempty"`
	Leader        *UnitLeader      `paradox_field:"leader" json:"leader,omitempty"`
	Location      string           `paradox_field:"location" json:"location,omitempty"`
	IsPrioritized pserialize.PBool `paradox_field:"is_prioritized" json:"is_priority,omitempty"`
	CanReinforce  pserialize.PBool `paradox_field:"can_reinforce" json:"is_strategic,omitempty"`
	CanUpgrade    pserialize.PBool `paradox_field:"can_upgrade" json:"is_defensive,omitempty"`
	Fuel          float64          `paradox_field:"fuel" json:"fuel,omitempty"`
	Supplies      float64          `paradox_field:"supplies" json:"supply,omitempty"`
	ArmyGroups    []*UnitContainer `paradox_field:"armygroup" paradox_type:"list" json:"armygroup,omitempty"`
	Armies        []*UnitContainer `paradox_field:"army" paradox_type:"list" json:"army,omitempty"`
	Corps         []*UnitContainer `paradox_field:"corps" paradox_type:"list" json:"corps,omitempty"`
	Divisions     []*UnitContainer `paradox_field:"division" paradox_type:"list" json:"division,omitempty"`
	Regiments     []*Regiment      `paradox_field:"regiment" paradox_type:"list" json:"regiments,omitempty"`
}

func (u *UnitContainer) GetProvincesStrength() map[string]*UnitStrength {
	provincesStrength := make(map[string]*UnitStrength)

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
			agStrength := ag.GetProvincesStrength()
			for province, strength := range agStrength {
				if _, ok := provincesStrength[province]; !ok {
					provincesStrength[province] = &UnitStrength{}
				}
				provincesStrength[province].Strength += strength.Strength
				provincesStrength[province].HighestStrength += strength.HighestStrength
				provincesStrength[province].Organisation += strength.Organisation
			}
		}
	}

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
			aStrength := a.GetProvincesStrength()
			for province, strength := range aStrength {
				if _, ok := provincesStrength[province]; !ok {
					provincesStrength[province] = &UnitStrength{}
				}
				provincesStrength[province].Strength += strength.Strength
				provincesStrength[province].HighestStrength += strength.HighestStrength
				provincesStrength[province].Organisation += strength.Organisation
			}
		}
	}

	if len(u.Corps) > 0 {
		for _, c := range u.Corps {
			cStrength := c.GetProvincesStrength()
			for province, strength := range cStrength {
				if _, ok := provincesStrength[province]; !ok {
					provincesStrength[province] = &UnitStrength{}
				}
				provincesStrength[province].Strength += strength.Strength
				provincesStrength[province].HighestStrength += strength.HighestStrength
				provincesStrength[province].Organisation += strength.Organisation
			}
		}
	}

	if len(u.Divisions) > 0 {
		for _, d := range u.Divisions {
			dStrength := d.GetProvincesStrength()
			for province, strength := range dStrength {
				if _, ok := provincesStrength[province]; !ok {
					provincesStrength[province] = &UnitStrength{}
				}
				provincesStrength[province].Strength += strength.Strength
				provincesStrength[province].HighestStrength += strength.HighestStrength
				provincesStrength[province].Organisation += strength.Organisation
			}
		}
	}

	if len(u.Regiments) > 0 {
		for _, r := range u.Regiments {
			location := r.Location
			if location == "" {
				location = u.Location
			}

			if _, ok := provincesStrength[location]; !ok {
				provincesStrength[location] = &UnitStrength{}
			}
			rStrength := r.GetStrength()
			provincesStrength[location].Strength += rStrength.Strength
			provincesStrength[location].HighestStrength += rStrength.HighestStrength
			provincesStrength[location].Organisation += rStrength.Organisation
		}
	}

	return provincesStrength
}

func (u *UnitContainer) GetStrength() *UnitStrength {
	strength := &UnitStrength{}

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
			agStrength := ag.GetStrength()
			strength.Strength += agStrength.Strength
			strength.HighestStrength += agStrength.HighestStrength
			strength.Organisation += agStrength.Organisation
		}
	}

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
			aStrength := a.GetStrength()
			strength.Strength += aStrength.Strength
			strength.HighestStrength += aStrength.HighestStrength
			strength.Organisation += aStrength.Organisation
		}
	}

	if len(u.Corps) > 0 {
		for _, c := range u.Corps {
			cStrength := c.GetStrength()
			strength.Strength += cStrength.Strength
			strength.HighestStrength += cStrength.HighestStrength
			strength.Organisation += cStrength.Organisation
		}
	}

	if len(u.Divisions) > 0 {
		for _, d := range u.Divisions {
			dStrength := d.GetStrength()
			strength.Strength += dStrength.Strength
			strength.HighestStrength += dStrength.HighestStrength
			strength.Organisation += dStrength.Organisation
		}
	}

	if len(u.Regiments) > 0 {
		for _, r := range u.Regiments {
			rStrength := r.GetStrength()
			strength.Strength += rStrength.Strength
			strength.HighestStrength += rStrength.HighestStrength
			strength.Organisation += rStrength.Organisation
		}
	}

	return strength
}
