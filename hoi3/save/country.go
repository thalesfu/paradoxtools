package save

type CountryInSave struct {
	ID         string           `paradox_type:"map_key" json:"id,omitempty"`
	Theatres   []*UnitContainer `paradox_field:"theatre" paradox_type:"list"  json:"theatres,omitempty"`
	ArmyGroups []*UnitContainer `paradox_field:"armygroup" paradox_type:"list" json:"armygroup,omitempty"`
	Armies     []*UnitContainer `paradox_field:"army" paradox_type:"list" json:"army,omitempty"`
	Corps      []*UnitContainer `paradox_field:"corps" paradox_type:"list" json:"corps,omitempty"`
	Divisions  []*UnitContainer `paradox_field:"division" paradox_type:"list" json:"division,omitempty"`
	Regiments  []*Regiment      `paradox_field:"regiment" paradox_type:"list" json:"regiments,omitempty"`
}

func (c *CountryInSave) GetProvincesStrength() map[string]*UnitStrength {
	provincesStrength := make(map[string]*UnitStrength)

	if len(c.Theatres) > 0 {
		for _, t := range c.Theatres {
			tStrength := t.GetProvincesStrength()
			for province, strength := range tStrength {
				if _, ok := provincesStrength[province]; !ok {
					provincesStrength[province] = &UnitStrength{}
				}
				provincesStrength[province].Strength += strength.Strength
				provincesStrength[province].HighestStrength += strength.HighestStrength
				provincesStrength[province].Organisation += strength.Organisation
			}
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
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

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
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

	if len(c.Corps) > 0 {
		for _, c := range c.Corps {
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

	if len(c.Divisions) > 0 {
		for _, d := range c.Divisions {
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

	if len(c.Regiments) > 0 {
		for _, r := range c.Regiments {
			location := r.Location

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

func (c *CountryInSave) GetStrength() *UnitStrength {
	strength := &UnitStrength{}

	if len(c.Theatres) > 0 {
		for _, t := range c.Theatres {
			tStrength := t.GetStrength()
			strength.Strength += tStrength.Strength
			strength.HighestStrength += tStrength.HighestStrength
			strength.Organisation += tStrength.Organisation
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
			agStrength := ag.GetStrength()
			strength.Strength += agStrength.Strength
			strength.HighestStrength += agStrength.HighestStrength
			strength.Organisation += agStrength.Organisation
		}
	}

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
			aStrength := a.GetStrength()
			strength.Strength += aStrength.Strength
			strength.HighestStrength += aStrength.HighestStrength
			strength.Organisation += aStrength.Organisation
		}
	}

	if len(c.Corps) > 0 {
		for _, c := range c.Corps {
			cStrength := c.GetStrength()
			strength.Strength += cStrength.Strength
			strength.HighestStrength += cStrength.HighestStrength
			strength.Organisation += cStrength.Organisation
		}
	}

	if len(c.Divisions) > 0 {
		for _, d := range c.Divisions {
			dStrength := d.GetStrength()
			strength.Strength += dStrength.Strength
			strength.HighestStrength += dStrength.HighestStrength
			strength.Organisation += dStrength.Organisation
		}
	}

	if len(c.Regiments) > 0 {
		for _, r := range c.Regiments {
			rStrength := r.GetStrength()
			strength.Strength += rStrength.Strength
			strength.HighestStrength += rStrength.HighestStrength
			strength.Organisation += rStrength.Organisation
		}
	}

	return strength
}
