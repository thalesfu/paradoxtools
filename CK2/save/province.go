package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Province struct {
	ID                     int                        `paradox_type:"map_key" json:"id,omitempty"`
	Name                   string                     `paradox_field:"name" json:"name,omitempty"`
	MaxSettlements         int                        `paradox_field:"max_settlements" json:"max_settlements,omitempty"`
	Winter                 int                        `paradox_field:"winter" json:"winter,omitempty"`
	Religion               *Religion                  `paradox_field:"religion" paradox_type:"map_value" paradox_map_name:"religion" json:"religion,omitempty"`
	Modifier               []*Modifier                `paradox_field:"modifier" paradox_type:"list" json:"modifier,omitempty"`
	Technology             *Technology                `paradox_field:"technology" paradox_type:"map_value" paradox_map_name:"technology" json:"technology,omitempty"`
	Loot                   []float32                  `paradox_field:"loot" paradox_type:"field_list" json:"loot,omitempty"`
	PrimarySettlement      string                     `paradox_field:"primary_settlement" json:"primary_settlement,omitempty"`
	Barons                 map[string]*Baron          `paradox_type:"map" paradox_map_key_pattern:"^b_" json:"barons,omitempty"`
	Disease                *Disease                   `paradox_field:"disease" json:"disease,omitempty"`
	Flags                  map[string]pserialize.Year `paradox_field:"flags" paradox_type:"map_key"  json:"flags,omitempty"`
	Vars                   map[string]float32         `paradox_field:"vars" paradox_type:"map_key"  json:"vars,omitempty"`
	Fort                   *Fort                      `paradox_field:"fort" json:"fort,omitempty"`
	TradePost              *TradePost                 `paradox_field:"trade_post" json:"trade_post,omitempty"`
	SettlementConstruction *SettlementConstruction    `paradox_field:"settlement_construction" json:"settlement_construction,omitempty"`
	DelayedEvents          []*Event                   `paradox_field:"delayed_event" paradox_type:"list" json:"delayed_event,omitempty"`
}

type Technology struct {
	TechLevels []float32 `paradox_field:"tech_levels" paradox_type:"field_list" json:"tech_levels,omitempty"`
}
