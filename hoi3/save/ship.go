package save

type Ship struct {
	ID                               *UnitID          `paradox_field:"id" json:"id,omitempty"`
	Name                             string           `paradox_field:"name" json:"name,omitempty"`
	Type                             string           `paradox_field:"type" json:"type,omitempty"`
	Organisation                     float64          `paradox_field:"organisation" json:"organisation,omitempty"`
	Strength                         float64          `paradox_field:"strength" json:"strength,omitempty"`
	HighestStrength                  float64          `paradox_field:"highest" json:"highest_xp,omitempty"`
	Experience                       float64          `paradox_field:"experience" json:"experience,omitempty"`
	Builder                          string           `paradox_field:"builder" json:"builder,omitempty"`
	Airs                             []*UnitContainer `paradox_field:"air" paradox_type:"list" json:"airs,omitempty"`
	Wings                            []*Wing          `paradox_field:"wing" paradox_type:"list" json:"wings,omitempty"`
	AmphibiousAssaultUnits           []float64        `paradox_field:"amphibious_assault_units" paradox_type:"field_list" json:"amphibious_assault_units,omitempty"`
	AmphibiousInvasionTechnology     []float64        `paradox_field:"amphibious_invasion_technology" paradox_type:"field_list" json:"amphibious_invasion_technology,omitempty"`
	BattleFleetConcentrationDoctrine []float64        `paradox_field:"battlefleet_concentration_doctrine" paradox_type:"field_list" json:"battlefleet_concentration_doctrine,omitempty"`
	BattleShipAntiaircraft           []float64        `paradox_field:"battleship_antiaircraft" paradox_type:"field_list" json:"battleship_antiaircraft,omitempty"`
	BattleShipArmour                 []float64        `paradox_field:"battleship_armour" paradox_type:"field_list" json:"battleship_armour,omitempty"`
	BattleShipCrewTraining           []float64        `paradox_field:"battleship_crew_training" paradox_type:"field_list" json:"battleship_crew_training,omitempty"`
	BattleShipEngine                 []float64        `paradox_field:"battleship_engine" paradox_type:"field_list" json:"battleship_engine,omitempty"`
	CapitalShipArmament              []float64        `paradox_field:"capitalship_armament" paradox_type:"field_list" json:"capitalship_armament,omitempty"`
	CarrierAntiaircraft              []float64        `paradox_field:"carrier_antiaircraft" paradox_type:"field_list" json:"carrier_antiaircraft,omitempty"`
	CarrierArmour                    []float64        `paradox_field:"carrier_armour" paradox_type:"field_list" json:"carrier_armour,omitempty"`
	CarrierCrewTraining              []float64        `paradox_field:"carrier_crew_training" paradox_type:"field_list" json:"carrier_crew_training,omitempty"`
	CarrierEngine                    []float64        `paradox_field:"carrier_engine" paradox_type:"field_list" json:"carrier_engine,omitempty"`
	CarrierGroupDoctrine             []float64        `paradox_field:"carrier_group_doctrine" paradox_type:"field_list" json:"carrier_group_doctrine,omitempty"`
	CarrierHanger                    []float64        `paradox_field:"carrier_hanger" paradox_type:"field_list" json:"carrier_hanger,omitempty"`
	CruiserCrewTraining              []float64        `paradox_field:"cruiser_crew_training" paradox_type:"field_list" json:"cruiser_crew_training,omitempty"`
	CruiserWarfare                   []float64        `paradox_field:"cruiser_warfare" paradox_type:"field_list" json:"cruiser_warfare,omitempty"`
	DestroyerAntiaircraft            []float64        `paradox_field:"destroyer_antiaircraft" paradox_type:"field_list" json:"destroyer_antiaircraft,omitempty"`
	DestroyerArmament                []float64        `paradox_field:"destroyer_armament" paradox_type:"field_list" json:"destroyer_armament,omitempty"`
	DestroyerArmour                  []float64        `paradox_field:"destroyer_armour" paradox_type:"field_list" json:"destroyer_armour,omitempty"`
	DestroyerCrewTraining            []float64        `paradox_field:"destroyer_crew_training" paradox_type:"field_list" json:"destroyer_crew_training,omitempty"`
	DestroyerEngine                  []float64        `paradox_field:"destroyer_engine" paradox_type:"field_list" json:"destroyer_engine,omitempty"`
	DestroyerEscortRole              []float64        `paradox_field:"destroyer_escort_role" paradox_type:"field_list" json:"destroyer_escort_role"`
	ElectricPoweredTorpedo           []float64        `paradox_field:"electric_powered_torpedo" paradox_type:"field_list" json:"electric_powered_torpedo,omitempty"`
	FleetAuxiliarySubmarineDoctrine  []float64        `paradox_field:"fleet_auxiliary_submarine_doctrine" paradox_type:"field_list" json:"fleet_auxiliary_submarine_doctrine,omitempty"`
	HeavyCruiserAntiaircraft         []float64        `paradox_field:"heavycruiser_antiaircraft" paradox_type:"field_list" json:"heavycruiser_antiaircraft,omitempty"`
	HeavyCruiserArmament             []float64        `paradox_field:"heavycruiser_armament" paradox_type:"field_list" json:"heavycruiser_armament,omitempty"`
	HeavyCruiserArmour               []float64        `paradox_field:"heavycruiser_armour" paradox_type:"field_list" json:"heavycruiser_armour,omitempty"`
	HeavyCruiserEngine               []float64        `paradox_field:"heavycruiser_engine" paradox_type:"field_list" json:"heavycruiser_engine,omitempty"`
	LargeWarshipRadar                []float64        `paradox_field:"largewarship_radar" paradox_type:"field_list" json:"largewarship_radar,omitempty"`
	LightCruiserAntiaircraft         []float64        `paradox_field:"lightcruiser_antiaircraft" paradox_type:"field_list" json:"lightcruiser_antiaircraft,omitempty"`
	LightCruiserArmament             []float64        `paradox_field:"lightcruiser_armament" paradox_type:"field_list" json:"lightcruiser_armament,omitempty"`
	LightCruiserArmour               []float64        `paradox_field:"lightcruiser_armour" paradox_type:"field_list" json:"lightcruiser_armour,omitempty"`
	LightCruiserCrewTraining         []float64        `paradox_field:"light_cruiser_crew_training" paradox_type:"field_list" json:"light_cruiser_crew_training,omitempty"`
	LightCruiserEngine               []float64        `paradox_field:"lightcruiser_engine" paradox_type:"field_list" json:"lightcruiser_engine,omitempty"`
	LightCruiserEscortRole           []float64        `paradox_field:"light_cruiser_escort_role" paradox_type:"field_list" json:"light_cruiser_escort_role,omitempty"`
	RadarTraining                    []float64        `paradox_field:"radar_training" paradox_type:"field_list" json:"radar_training,omitempty"`
	SmallWarshipAsw                  []float64        `paradox_field:"smallwarship_asw" paradox_type:"field_list" json:"smallwarship_asw,omitempty"`
	SmallWarshipRadar                []float64        `paradox_field:"smallwarship_radar" paradox_type:"field_list" json:"smallwarship_radar,omitempty"`
	Spotting                         []float64        `paradox_field:"spotting" paradox_type:"field_list" json:"spotting,omitempty"`
	SubmarineAirWarningEquipment     []float64        `paradox_field:"submarine_airwarningequipment" paradox_type:"field_list" json:"submarine_airwarningequipment,omitempty"`
	SubmarineAntiaircraft            []float64        `paradox_field:"submarine_antiaircraft" paradox_type:"field_list" json:"submarine_antiaircraft,omitempty"`
	SubmarineCrewTraining            []float64        `paradox_field:"submarine_crew_training" paradox_type:"field_list" json:"submarine_crew_training,omitempty"`
	SubmarineEngine                  []float64        `paradox_field:"submarine_engine" paradox_type:"field_list" json:"submarine_engine,omitempty"`
	SubmarineHull                    []float64        `paradox_field:"submarine_hull" paradox_type:"field_list" json:"submarine_hull,omitempty"`
	SubmarineSonar                   []float64        `paradox_field:"submarine_sonar" paradox_type:"field_list" json:"submarine_sonar,omitempty"`
	SubmarineTorpedoes               []float64        `paradox_field:"submarine_torpedoes" paradox_type:"field_list" json:"submarine_torpedoes,omitempty"`
}

func (s *Ship) GetProvincesAirCount() map[string]map[string]*UnitCount {
	provincesCount := make(map[string]map[string]*UnitCount)

	if len(s.Airs) > 0 {
		for _, air := range s.Airs {
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

func (s *Ship) GetAirCount() map[string]*UnitCount {
	unitCounts := make(map[string]*UnitCount)

	if len(s.Airs) > 0 {
		for _, a := range s.Airs {
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

	if len(s.Wings) > 0 {
		for _, wing := range s.Wings {
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
