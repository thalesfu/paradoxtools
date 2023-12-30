package save

type MercenaryType struct {
	Name          string  `paradox_field:"name" json:"name,omitempty"`
	Father        int     `paradox_field:"father" json:"father,omitempty"`
	LevySize      float32 `paradox_field:"levy_size" json:"levy_size,omitempty"`
	LightInfantry float32 `paradox_field:"light_infantry" json:"light_infantry,omitempty"`
	HeavyInfantry float32 `paradox_field:"heavy_infantry" json:"heavy_infantry,omitempty"`
	Pikemen       float32 `paradox_field:"pikemen" json:"pikemen,omitempty"`
	LightCavalry  float32 `paradox_field:"light_cavalry" json:"light_cavalry,omitempty"`
	Knights       float32 `paradox_field:"knights" json:"knights,omitempty"`
	Archers       float32 `paradox_field:"archers" json:"archers,omitempty"`
	Galleys       float32 `paradox_field:"galleys" json:"galleys,omitempty"`
	NumSubUnits   float32 `paradox_field:"num_subunits" json:"num_subunits,omitempty"`
	HorseArchers  float32 `paradox_field:"horse_archers" json:"horse_archers,omitempty"`
}
