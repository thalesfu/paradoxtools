package save

type SupplyInSave struct {
	Supplies float64 `paradox_field:"supplies" json:"supplies,omitempty"`
	Fuel     float64 `paradox_field:"fuel" json:"fuel,omitempty"`
}
