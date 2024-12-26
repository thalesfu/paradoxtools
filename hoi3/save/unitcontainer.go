package save

import (
	"github.com/thalesfu/paradoxtools/utils/pserialize"
)

type UnitContainer struct {
	ID            *UnitID          `paradox_field:"id" json:"id,omitempty"`
	Name          string           `paradox_field:"name" json:"name,omitempty"`
	Base          string           `paradox_field:"base" json:"base,omitempty"`
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
	Navies        []*UnitContainer `paradox_field:"navy" paradox_type:"list" json:"navy,omitempty"`
	Ships         []*Ship          `paradox_field:"ship" paradox_type:"list" json:"ships,omitempty"`
	Airs          []*UnitContainer `paradox_field:"air" paradox_type:"list" json:"airs,omitempty"`
	Wings         []*Wing          `paradox_field:"wing" paradox_type:"list" json:"wings,omitempty"`
}

func (u *UnitContainer) GetProvincesAirCount() map[string]map[string]*UnitCount {
	provincesCount := make(map[string]map[string]*UnitCount)

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
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

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
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

	if len(u.Corps) > 0 {
		for _, corp := range u.Corps {
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

	if len(u.Ships) > 0 {
		for _, ship := range u.Ships {
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

	if len(u.Airs) > 0 {
		for _, air := range u.Airs {
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

	if len(u.Wings) > 0 {
		for _, wing := range u.Wings {
			if _, ok := provincesCount[u.Location]; !ok {
				provincesCount[u.Location] = make(map[string]*UnitCount)
			}

			if _, ok := provincesCount[u.Location][wing.Type]; !ok {
				provincesCount[u.Location][wing.Type] = &UnitCount{
					Type: wing.Type,
				}
			}
			provincesCount[u.Location][wing.Type].Count += 1
			provincesCount[u.Location][wing.Type].Strength += wing.Strength
			provincesCount[u.Location][wing.Type].HighestStrength += wing.HighestStrength
			provincesCount[u.Location][wing.Type].Organisation += wing.Organisation
		}
	}

	return provincesCount
}

func (u *UnitContainer) GetAirCount() map[string]*UnitCount {
	unitCounts := make(map[string]*UnitCount)

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
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

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
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

	if len(u.Corps) > 0 {
		for _, c := range u.Corps {
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

	if len(u.Ships) > 0 {
		for _, s := range u.Ships {
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

	if len(u.Airs) > 0 {
		for _, a := range u.Airs {
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

	if len(u.Wings) > 0 {
		for _, wing := range u.Wings {
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

func (u *UnitContainer) GetProvincesNavyCount() map[string]map[string]*UnitCount {
	provincesCount := make(map[string]map[string]*UnitCount)

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
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

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
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

	if len(u.Corps) > 0 {
		for _, corp := range u.Corps {
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

	if len(u.Navies) > 0 {
		for _, navy := range u.Navies {
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

	if len(u.Ships) > 0 {
		for _, ship := range u.Ships {
			if _, ok := provincesCount[u.Location]; !ok {
				provincesCount[u.Location] = make(map[string]*UnitCount)
			}

			if _, ok := provincesCount[u.Location][ship.Type]; !ok {
				provincesCount[u.Location][ship.Type] = &UnitCount{
					Type: ship.Type,
				}
			}
			provincesCount[u.Location][ship.Type].Count += 1
			provincesCount[u.Location][ship.Type].Strength += ship.Strength
			provincesCount[u.Location][ship.Type].HighestStrength += ship.HighestStrength
			provincesCount[u.Location][ship.Type].Organisation += ship.Organisation
		}
	}

	return provincesCount
}

func (u *UnitContainer) GetNavyCount() map[string]*UnitCount {
	unitCounts := make(map[string]*UnitCount)

	if len(u.ArmyGroups) > 0 {
		for _, ag := range u.ArmyGroups {
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

	if len(u.Armies) > 0 {
		for _, a := range u.Armies {
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

	if len(u.Corps) > 0 {
		for _, c := range u.Corps {
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

	if len(u.Navies) > 0 {
		for _, n := range u.Navies {
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

	if len(u.Ships) > 0 {
		for _, s := range u.Ships {
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
