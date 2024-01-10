package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Character struct {
	ID                  int                  `paradox_type:"map_key" json:"id,omitempty"`
	Name                string               `paradox_field:"bn" paradox_text:"escaped" json:"bn,omitempty"`
	Hist                pserialize.PBool     `paradox_field:"hist" json:"hist,omitempty"`
	Pi                  pserialize.PBool     `paradox_field:"pi" json:"pi,omitempty"`
	Birthday            pserialize.Year      `paradox_field:"b_d" json:"b_d,omitempty"`
	DeathDay            pserialize.Year      `paradox_field:"d_d" json:"d_d,omitempty"`
	Father              int                  `paradox_field:"fat" json:"fat,omitempty"`
	Spouse              int                  `paradox_field:"spouse" json:"spouse,omitempty"`
	Attributes          []int                `paradox_field:"att" paradox_type:"field_list" json:"att,omitempty"`
	Traits              []int                `paradox_field:"tr" paradox_type:"field_list" json:"tr,omitempty"`
	Dynasty             int                  `paradox_field:"dnt" json:"dnt,omitempty"`
	DNA                 string               `paradox_field:"dna" json:"dna,omitempty"`
	Government          string               `paradox_field:"gov" json:"gov,omitempty"`
	Fertility           float32              `paradox_field:"fer" json:"fer,omitempty"`
	Health              float32              `paradox_field:"health" json:"health,omitempty"`
	Consort             []int                `paradox_field:"consort" paradox_type:"list" json:"consort,omitempty"`
	Prs                 float32              `paradox_field:"prs" json:"prs,omitempty"`
	Piety               float32              `paradox_field:"piety" json:"piety,omitempty"`
	Wealth              float32              `paradox_field:"wealth" json:"wealth,omitempty"`
	Emp                 int                  `paradox_field:"emp" json:"emp,omitempty"`
	Host                int                  `paradox_field:"host" json:"host,omitempty"`
	Curinc              float32              `paradox_field:"curinc" json:"curinc,omitempty"`
	Emi                 float32              `paradox_field:"emi" json:"emi,omitempty"`
	Eme                 float32              `paradox_field:"eme" json:"eme,omitempty"`
	Eyi                 float32              `paradox_field:"eyi" json:"eyi,omitempty"`
	Eypi                float32              `paradox_field:"eypi" json:"eypi,omitempty"`
	Action              string               `paradox_field:"action" json:"action,omitempty"`
	ACDate              pserialize.Year      `paradox_field:"acdate" json:"acdate,omitempty"`
	ACLoc               int                  `paradox_field:"acloc" json:"acloc,omitempty"`
	AmbitionDate        pserialize.Year      `paradox_field:"ambition_date" json:"ambition_date,omitempty"`
	FocusDate           pserialize.Year      `paradox_field:"focus_date" json:"focus_date,omitempty"`
	CPos                string               `paradox_field:"cpos" json:"cpos,omitempty"`
	CPosStart           int                  `paradox_field:"cpos_start" json:"cpos_start,omitempty"`
	Vars                *Vars                `paradox_field:"vars" json:"vars,omitempty"`
	Flags               *Flag                `paradox_field:"flags" json:"flags,omitempty"`
	Modifier            *Modifier            `paradox_field:"md" json:"md,omitempty"`
	Claim               *Claim               `paradox_field:"claim" json:"claim,omitempty"`
	Lgr                 *Lgr                 `paradox_field:"lgr" json:"lgr,omitempty"`
	SavedEventTarget    *SavedEventTarget    `paradox_field:"saved_event_target" json:"saved_event_target,omitempty"`
	OffMapCurrencies    map[int]float32      `paradox_field:"offmap_currencies" json:"offmap_currencies,omitempty"`
	DMN                 *DMN                 `paradox_field:"dmn" json:"dmn,omitempty"`
	CharacterPlayerData *CharacterPlayerData `paradox_field:"character_player_data" json:"character_player_data,omitempty"`
	Job                 string               `paradox_field:"job" json:"job,omitempty"`
}
