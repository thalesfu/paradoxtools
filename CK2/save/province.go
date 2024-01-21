package save

import (
	"github.com/thalesfu/paradoxtools/CK2/history/province"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"time"
)

type Province struct {
	ID                     int                        `paradox_type:"map_key" json:"id,omitempty"`
	Code                   string                     `json:"code,omitempty"`
	Name                   string                     `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	MaxSettlements         int                        `paradox_field:"max_settlements" json:"max_settlements,omitempty"`
	Winter                 int                        `paradox_field:"winter" json:"winter,omitempty"`
	Culture                string                     `paradox_field:"culture" json:"culture,omitempty"`
	Religion               string                     `paradox_field:"religion" json:"religion,omitempty"`
	Modifiers              []*Modifier                `paradox_field:"modifier" paradox_type:"list" json:"modifier,omitempty"`
	Technology             *Technology                `paradox_field:"technology" paradox_type:"map_value" paradox_map_name:"technology" json:"technology,omitempty"`
	Loot                   []float32                  `paradox_field:"loot" paradox_type:"field_list" json:"loot,omitempty"`
	PrimarySettlement      string                     `paradox_field:"primary_settlement" json:"primary_settlement,omitempty"`
	Barons                 map[string]*Baron          `paradox_type:"map" paradox_map_key_pattern:"^b_" json:"barons,omitempty"`
	Disease                *Disease                   `paradox_field:"disease" json:"disease,omitempty"`
	Flags                  map[string]pserialize.Year `paradox_field:"flags" json:"flags,omitempty"`
	Vars                   map[string]float32         `paradox_field:"vars" json:"vars,omitempty"`
	Fort                   *Fort                      `paradox_field:"fort" json:"fort,omitempty"`
	TradePost              *TradePost                 `paradox_field:"trade_post" json:"trade_post,omitempty"`
	SettlementConstruction *SettlementConstruction    `paradox_field:"settlement_construction" json:"settlement_construction,omitempty"`
	DelayedEvents          []*Event                   `paradox_field:"delayed_event" paradox_type:"list" json:"delayed_event,omitempty"`
	PlayID                 int                        `description:"game play id" json:"play_id,omitempty"`
	PlayDate               time.Time                  `description:"game play date" json:"play_date,omitempty"`
}

type Technology struct {
	TechLevels []float32 `paradox_field:"tech_levels" paradox_type:"field_list" json:"tech_levels,omitempty"`
}

func processProvinces(saveFile *SaveFile, translations map[string]string) {
	for _, p := range saveFile.Provinces {
		p.PlayID = saveFile.PlayThroughID
		p.PlayDate = time.Time(saveFile.Date)

		if province, ok := province.ProvinceIdList[p.ID]; ok {
			p.Code = province.GetTitleCode()

			if p.Name == "" {
				p.Name = province.GetTitleName()
			}
		}

		if p.Name == "" {
			p.Name = translations[p.Code]
		}

		if p.Name == "" {
			p.Name = translations[p.Code+"_adj"]
		}
	}
}
