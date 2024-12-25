package save

type ProvinceInSave struct {
	ID               string           `paradox_type:"map_key" json:"id,omitempty"`
	Owner            string           `paradox_field:"owner" json:"owner,omitempty"`
	Controller       string           `paradox_field:"controller" json:"controller,omitempty"`
	Core             []string         `paradox_field:"core" paradox_type:"list" json:"core,omitempty"`
	Point            int              `paradox_field:"points" json:"point,omitempty"`
	Manpower         float64          `paradox_field:"manpower" json:"manpower,omitempty"`
	Leadership       float64          `paradox_field:"leadership" json:"leadership,omitempty"`
	CurrentProducing *ProducingInSave `paradox_field:"current_producing" json:"current_producing,omitempty"`
	MaxProducing     *ProducingInSave `paradox_field:"max_producing" json:"max_producing,omitempty"`
	AirBase          []float64        `paradox_field:"air_base" paradox_type:"field_list" json:"air_base,omitempty"`
	NavalBase        []float64        `paradox_field:"naval_base" paradox_type:"field_list" json:"naval_base,omitempty"`
	LandFort         []float64        `paradox_field:"land_fort" paradox_type:"field_list" json:"land_fort,omitempty"`
	CoastalFort      []float64        `paradox_field:"coastal_fort" paradox_type:"field_list" json:"coastal_fort,omitempty"`
	AntiAir          []float64        `paradox_field:"anti_air" paradox_type:"field_list" json:"anti_air,omitempty"`
	RadarStation     []float64        `paradox_field:"radar_station" paradox_type:"field_list" json:"radar_station,omitempty"`
	NuclearReactor   []float64        `paradox_field:"nuclear_reactor" paradox_type:"field_list" json:"nuclear_reactor,omitempty"`
	RocketTest       []float64        `paradox_field:"rocket_test" paradox_type:"field_list" json:"rocket_test,omitempty"`
	Industry         []float64        `paradox_field:"industry" paradox_type:"field_list" json:"industry,omitempty"`
	Infra            []float64        `paradox_field:"infra" paradox_type:"field_list" json:"infra,omitempty"`
	Pool             *SupplyInSave    `paradox_field:"pool" json:"pool,omitempty"`
	Drawn            *SupplyInSave    `paradox_field:"drawn" json:"drawn,omitempty"`
	LastDrawn        *SupplyInSave    `paradox_field:"last_drawn" json:"last_drawn,omitempty"`
	Throughput       *SupplyInSave    `paradox_field:"throughput" json:"throughput,omitempty"`
	LastThroughput   *SupplyInSave    `paradox_field:"last_throughput" json:"last_throughput,omitempty"`
}
