package save

type DMN struct {
	Armies                  []*Army     `paradox_field:"army" paradox_type:"list"  json:"army,omitempty"`
	Capital                 string      `paradox_field:"capital" json:"capital,omitempty"`
	Primary                 *Title      `paradox_field:"primary" paradox_type:"map_value" paradox_map_name:"title" json:"primary,omitempty"`
	LiegeTroops             *ArmyDetail `paradox_field:"liege_troops" json:"liege_troops,omitempty"`
	RaisedLiegeTroops       []int       `paradox_field:"raised_liege_troops" paradox_type:"field_list" json:"raised_liege_troops,omitempty"`
	MyLiegelevyContribution int         `paradox_field:"my_liegelevy_contribution" json:"my_liegelevy_contribution,omitempty"`
	PeaceMonths             int         `paradox_field:"peace_months" json:"peace_months,omitempty"`
}
