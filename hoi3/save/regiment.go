package save

type Regiment struct {
	ID                        *UnitID   `paradox_field:"id" json:"id,omitempty"`
	Name                      string    `paradox_field:"name" json:"name,omitempty"`
	Home                      int       `paradox_field:"home" json:"home,omitempty"`
	Type                      string    `paradox_field:"type" json:"type,omitempty"`
	Location                  string    `paradox_field:"location" json:"location,omitempty"`
	Organisation              float64   `paradox_field:"organisation" json:"organisation,omitempty"`
	Strength                  float64   `paradox_field:"strength" json:"strength,omitempty"`
	HighestStrength           float64   `paradox_field:"highest" json:"highest_xp,omitempty"`
	Experience                float64   `paradox_field:"experience" json:"experience,omitempty"`
	MobileWarfare             []float64 `paradox_field:"mobile_warfare" paradox_type:"field_list" json:"mobile_warfare,omitempty"`
	ElasticDefence            []float64 `paradox_field:"elastic_defence" paradox_type:"field_list" json:"elastic_defence,omitempty"`
	DelayDoctrine             []float64 `paradox_field:"delay_doctrine" paradox_type:"field_list" json:"delay_doctrine,omitempty"`
	IntegratedSupportDoctrine []float64 `paradox_field:"integrated_support_doctrine" paradox_type:"field_list" json:"integrated_support_doctrine,omitempty"`
	CentralPlanning           []float64 `paradox_field:"central_planning" paradox_type:"field_list" json:"central_planning,omitempty"`
	MassAssault               []float64 `paradox_field:"mass_assault" paradox_type:"field_list" json:"mass_assault,omitempty"`
	LargeFront                []float64 `paradox_field:"large_front" paradox_type:"field_list" json:"large_front,omitempty"`
	GuerillaWarfare           []float64 `paradox_field:"guerilla_warfare" paradox_type:"field_list" json:"guerilla_warfare,omitempty"`
	InfantryWarfare           []float64 `paradox_field:"infantry_warfare" paradox_type:"field_list" json:"infantry_warfare,omitempty"`
	MedicalEvacuation         []float64 `paradox_field:"medical_evacuation" paradox_type:"field_list" json:"medical_evacuation,omitempty"`
	SmallArmsTechnology       []float64 `paradox_field:"smallarms_technology" paradox_type:"field_list" json:"smallarms_technology,omitempty"`
	InfantrySupport           []float64 `paradox_field:"infantry_support" paradox_type:"field_list" json:"infantry_support,omitempty"`
	InfantryGuns              []float64 `paradox_field:"infantry_guns" paradox_type:"field_list" json:"infantry_guns,omitempty"`
	InfantryAt                []float64 `paradox_field:"infantry_at" paradox_type:"field_list" json:"infantry_at,omitempty"`
	NightGoggles              []float64 `paradox_field:"night_goggles" paradox_type:"field_list" json:"night_goggles,omitempty"`
	DesertWarfareEquipment    []float64 `paradox_field:"desert_warfare_equipment" paradox_type:"field_list" json:"desert_warfare_equipment,omitempty"`
	JungleWarfareEquipment    []float64 `paradox_field:"jungle_warfare_equipment" paradox_type:"field_list" json:"jungle_warfare_equipment,omitempty"`
	MountainWarfareEquipment  []float64 `paradox_field:"mountain_warfare_equipment" paradox_type:"field_list" json:"mountain_warfare_equipment,omitempty"`
	ArticWarfareEquipment     []float64 `paradox_field:"artic_warfare_equipment" paradox_type:"field_list" json:"artic_warfare_equipment,omitempty"`
}

func (r *Regiment) GetStrength() *UnitStrength {
	return &UnitStrength{
		Organisation:    r.Organisation,
		Strength:        r.Strength,
		HighestStrength: r.HighestStrength,
	}
}
