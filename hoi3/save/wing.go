package save

type Wing struct {
	ID                            *UnitID   `paradox_field:"id" json:"id,omitempty"`
	Name                          string    `paradox_field:"name" json:"name,omitempty"`
	Home                          int       `paradox_field:"home" json:"home,omitempty"`
	Type                          string    `paradox_field:"type" json:"type,omitempty"`
	Location                      string    `paradox_field:"location" json:"location,omitempty"`
	Organisation                  float64   `paradox_field:"organisation" json:"organisation,omitempty"`
	Strength                      float64   `paradox_field:"strength" json:"strength,omitempty"`
	HighestStrength               float64   `paradox_field:"highest" json:"highest_xp,omitempty"`
	Experience                    float64   `paradox_field:"experience" json:"experience,omitempty"`
	Builder                       string    `paradox_field:"builder" json:"builder,omitempty"`
	AaMissile                     []float64 `paradox_field:"aamissile" paradox_type:"field_list" json:"aamissile,omitempty"`
	AeroEngine                    []float64 `paradox_field:"aeroengine" paradox_type:"field_list" json:"aeroengine,omitempty"`
	AirLaunchedTorpedo            []float64 `paradox_field:"air_launched_torpedo" paradox_type:"field_list" json:"air_launched_torpedo,omitempty"`
	BasicAeroEngine               []float64 `paradox_field:"basic_aeroengine" paradox_type:"field_list" json:"basic_aeroengine,omitempty"`
	BasicAircraftMachineGun       []float64 `paradox_field:"basic_aircraft_machinegun" paradox_type:"field_list" json:"basic_aircraft_machinegun,omitempty"`
	BasicBomb                     []float64 `paradox_field:"basic_bomb" paradox_type:"field_list" json:"basic_bomb,omitempty"`
	BasicMediumFuelTank           []float64 `paradox_field:"basic_medium_fueltank" paradox_type:"field_list" json:"basic_medium_fueltank,omitempty"`
	BasicSingleEngineAirframe     []float64 `paradox_field:"basic_single_engine_airframe" paradox_type:"field_list" json:"basic_single_engine_airframe,omitempty"`
	BasicSmallFuelTank            []float64 `paradox_field:"basic_small_fueltank" paradox_type:"field_list" json:"basic_small_fueltank,omitempty"`
	BasicTwinEngineAirframe       []float64 `paradox_field:"basic_twin_engine_airframe" paradox_type:"field_list" json:"basic_twin_engine_airframe,omitempty"`
	CasGroundCrewTraining         []float64 `paradox_field:"cas_groundcrew_training" paradox_type:"field_list" json:"cas_groundcrew_training,omitempty"`
	CasPilotTraining              []float64 `paradox_field:"cas_pilot_training" paradox_type:"field_list" json:"cas_pilot_training,omitempty"`
	DropTanks                     []float64 `paradox_field:"drop_tanks" paradox_type:"field_list" json:"drop_tanks,omitempty"`
	FighterGroundCrewTraining     []float64 `paradox_field:"fighter_groundcrew_training" paradox_type:"field_list" json:"fighter_groundcrew_training,omitempty"`
	FighterPilotTraining          []float64 `paradox_field:"fighter_pilot_training" paradox_type:"field_list" json:"fighter_pilot_training,omitempty"`
	FleetAuxiliaryCarrierDoctrine []float64 `paradox_field:"fleet_auxiliary_carrier_doctrine" paradox_type:"field_list" json:"fleet_auxiliary_carrier_doctrine,omitempty"`
	JetEngine                     []float64 `paradox_field:"jet_engine" paradox_type:"field_list" json:"jet_engine,omitempty"`
	MediumAirSearchRadar          []float64 `paradox_field:"medium_airsearch_radar" paradox_type:"field_list" json:"medium_airsearch_radar,omitempty"`
	MediumBomb                    []float64 `paradox_field:"medium_bomb" paradox_type:"field_list" json:"medium_bomb,omitempty"`
	MediumFuelTank                []float64 `paradox_field:"medium_fueltank" paradox_type:"field_list" json:"medium_fueltank,omitempty"`
	MediumNavigationRadar         []float64 `paradox_field:"medium_navagation_radar" paradox_type:"field_list" json:"medium_navagation_radar,omitempty"`
	NavGroundCrewTraining         []float64 `paradox_field:"nav_groundcrew_training" paradox_type:"field_list" json:"nav_groundcrew_training,omitempty"`
	NavPilotTraining              []float64 `paradox_field:"nav_pilot_training" paradox_type:"field_list" json:"nav_pilot_training,omitempty"`
	PilotRescue                   []float64 `paradox_field:"pilot_rescue" paradox_type:"field_list" json:"pilot_rescue,omitempty"`
	RadarGuidedBomb               []float64 `paradox_field:"radar_guided_bomb" paradox_type:"field_list" json:"radar_guided_bomb,omitempty"`
	RadarGuidedMissile            []float64 `paradox_field:"radar_guided_missile" paradox_type:"field_list" json:"radar_guided_missile,omitempty"`
	SingleEngineAircraftArmament  []float64 `paradox_field:"single_engine_aircraft_armament" paradox_type:"field_list" json:"single_engine_aircraft_armament,omitempty"`
	SingleEngineAirframe          []float64 `paradox_field:"single_engine_airframe" paradox_type:"field_list" json:"single_engine_airframe,omitempty"`
	SmallAirSearchRadar           []float64 `paradox_field:"small_airsearch_radar" paradox_type:"field_list" json:"small_airsearch_radar,omitempty"`
	SmallNavigationRadar          []float64 `paradox_field:"small_navagation_radar" paradox_type:"field_list" json:"small_navagation_radar,omitempty"`
	SmallBomb                     []float64 `paradox_field:"small_bomb" paradox_type:"field_list" json:"small_bomb,omitempty"`
	SmallFuelTank                 []float64 `paradox_field:"small_fueltank" paradox_type:"field_list" json:"small_fueltank,omitempty"`
	TacGroundCrewTraining         []float64 `paradox_field:"tac_groundcrew_training" paradox_type:"field_list" json:"tac_groundcrew_training,omitempty"`
	TacPilotTraining              []float64 `paradox_field:"tac_pilot_training" paradox_type:"field_list" json:"tac_pilot_training,omitempty"`
	TwinEngineAircraftArmament    []float64 `paradox_field:"twin_engine_aircraft_armament" paradox_type:"field_list" json:"twin_engine_aircraft_armament,omitempty"`
	TwinEngineAirframe            []float64 `paradox_field:"twin_engine_airframe" paradox_type:"field_list" json:"twin_engine_airframe,omitempty"`
}
