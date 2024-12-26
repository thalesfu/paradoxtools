package save

type CountryInSave struct {
	ID         string           `paradox_type:"map_key" json:"id,omitempty"`
	Theatres   []*UnitContainer `paradox_field:"theatre" paradox_type:"list"  json:"theatres,omitempty"`
	ArmyGroups []*UnitContainer `paradox_field:"armygroup" paradox_type:"list" json:"armygroup,omitempty"`
	Armies     []*UnitContainer `paradox_field:"army" paradox_type:"list" json:"army,omitempty"`
	Corps      []*UnitContainer `paradox_field:"corps" paradox_type:"list" json:"corps,omitempty"`
	Divisions  []*UnitContainer `paradox_field:"division" paradox_type:"list" json:"division,omitempty"`
	Regiments  []*Regiment      `paradox_field:"regiment" paradox_type:"list" json:"regiments,omitempty"`
	Navies     []*UnitContainer `paradox_field:"navy" paradox_type:"list" json:"navy,omitempty"`
	Ships      []*Ship          `paradox_field:"ship" paradox_type:"list" json:"ships,omitempty"`
	Airs       []*UnitContainer `paradox_field:"air" paradox_type:"list" json:"airs,omitempty"`
	Wings      []*Wing          `paradox_field:"wing" paradox_type:"list" json:"wings,omitempty"`
}

func (c *CountryInSave) GetProvincesAirCount() map[string]map[string]*UnitCount {
	provincesCount := make(map[string]map[string]*UnitCount)

	if len(c.Theatres) > 0 {
		for _, theatre := range c.Theatres {
			tuc := theatre.GetProvincesAirCount()
			for province, uc := range tuc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
			aguc := ag.GetProvincesAirCount()
			for province, uc := range aguc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
			auc := a.GetProvincesAirCount()
			for province, uc := range auc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Corps) > 0 {
		for _, corp := range c.Corps {
			cuc := corp.GetProvincesAirCount()
			for province, uc := range cuc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Ships) > 0 {
		for _, ship := range c.Ships {
			suc := ship.GetProvincesAirCount()
			for province, uc := range suc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Airs) > 0 {
		for _, air := range c.Airs {
			auc := air.GetProvincesAirCount()
			for province, uc := range auc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	return provincesCount
}

func (c *CountryInSave) GetAirCount() map[string]*UnitCount {
	unitCounts := make(map[string]*UnitCount)

	if len(c.Theatres) > 0 {
		for _, theatre := range c.Theatres {
			tUnitCounts := theatre.GetAirCount()
			for t, c := range tUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
			agUnitCounts := ag.GetAirCount()
			for t, c := range agUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
			aUnitCounts := a.GetAirCount()
			for t, c := range aUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Corps) > 0 {
		for _, c := range c.Corps {
			cUnitCounts := c.GetAirCount()
			for t, c := range cUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Ships) > 0 {
		for _, s := range c.Ships {
			sUnitCounts := s.GetAirCount()
			for t, c := range sUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Airs) > 0 {
		for _, a := range c.Airs {
			airUnitCounts := a.GetAirCount()
			for t, c := range airUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Wings) > 0 {
		for _, wing := range c.Wings {
			t := wing.Type
			if _, ok := unitCounts[t]; !ok {
				unitCounts[t] = &UnitCount{
					Type: t,
				}
			}
			unitCounts[t].Count += 1
			unitCounts[t].Strength += wing.Strength
			unitCounts[t].HighestStrength += wing.HighestStrength
			unitCounts[t].Organisation += wing.Organisation
		}
	}

	return unitCounts
}

func (c *CountryInSave) GetProvincesNavyCount() map[string]map[string]*UnitCount {
	provincesCount := make(map[string]map[string]*UnitCount)

	if len(c.Theatres) > 0 {
		for _, theatre := range c.Theatres {
			tuc := theatre.GetProvincesNavyCount()
			for province, uc := range tuc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
			aguc := ag.GetProvincesNavyCount()
			for province, uc := range aguc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
			auc := a.GetProvincesNavyCount()
			for province, uc := range auc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Corps) > 0 {
		for _, corp := range c.Corps {
			cuc := corp.GetProvincesNavyCount()
			for province, uc := range cuc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	if len(c.Navies) > 0 {
		for _, navy := range c.Navies {
			nuc := navy.GetProvincesNavyCount()
			for province, uc := range nuc {
				if _, ok := provincesCount[province]; !ok {
					provincesCount[province] = make(map[string]*UnitCount)
				}

				for t, c := range uc {
					if _, ok := provincesCount[province][t]; !ok {
						provincesCount[province][t] = &UnitCount{
							Type: t,
						}
					}
					provincesCount[province][t].Count += c.Count
					provincesCount[province][t].Strength += c.Strength
					provincesCount[province][t].HighestStrength += c.HighestStrength
					provincesCount[province][t].Organisation += c.Organisation
				}
			}
		}
	}

	return provincesCount
}

func (c *CountryInSave) GetNavyCount() map[string]*UnitCount {
	unitCounts := make(map[string]*UnitCount)

	if len(c.Theatres) > 0 {
		for _, t := range c.Theatres {
			tUnitCounts := t.GetNavyCount()
			for t, c := range tUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.ArmyGroups) > 0 {
		for _, ag := range c.ArmyGroups {
			agUnitCounts := ag.GetNavyCount()
			for t, c := range agUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Armies) > 0 {
		for _, a := range c.Armies {
			aUnitCounts := a.GetNavyCount()
			for t, c := range aUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Corps) > 0 {
		for _, c := range c.Corps {
			cUnitCounts := c.GetNavyCount()
			for t, c := range cUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Navies) > 0 {
		for _, n := range c.Navies {
			nUnitCounts := n.GetNavyCount()
			for t, c := range nUnitCounts {
				if _, ok := unitCounts[t]; !ok {
					unitCounts[t] = &UnitCount{
						Type: t,
					}
				}
				unitCounts[t].Count += c.Count
				unitCounts[t].Strength += c.Strength
				unitCounts[t].HighestStrength += c.HighestStrength
				unitCounts[t].Organisation += c.Organisation
			}
		}
	}

	if len(c.Ships) > 0 {
		for _, s := range c.Ships {
			t := s.Type
			if _, ok := unitCounts[t]; !ok {
				unitCounts[t] = &UnitCount{
					Type: t,
				}
			}
			unitCounts[t].Count += 1
			unitCounts[t].Strength += s.Strength
			unitCounts[t].HighestStrength += s.HighestStrength
			unitCounts[t].Organisation += s.Organisation
		}
	}

	return unitCounts
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
