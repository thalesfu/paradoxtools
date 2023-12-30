package save

type TradePost struct {
	Type  string      `paradox_field:"type" json:"type,omitempty"`
	Owner int         `paradox_field:"owner" json:"owner,omitempty"`
	Levy  *ArmyDetail `paradox_field:"levy" json:"levy,omitempty"`
}
