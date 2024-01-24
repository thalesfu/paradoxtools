package save

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"time"
)

type Character struct {
	ID   int    `paradox_type:"map_key" json:"id,omitempty"`
	Name string `paradox_field:"bn" paradox_text:"escaped" json:"name,omitempty"`

	Flags               map[string]pserialize.Year `paradox_field:"flags" json:"flags,omitempty"`
	ACDate              pserialize.Year            `paradox_field:"acdate" json:"acdate,omitempty"`
	AmbitionDate        pserialize.Year            `paradox_field:"ambition_date" json:"ambition_date,omitempty"`
	FocusDate           pserialize.Year            `paradox_field:"focus_date" json:"focus_date,omitempty"`
	Birthday            pserialize.Year            `paradox_field:"b_d" json:"birthday,omitempty"`
	DeathDay            pserialize.Year            `paradox_field:"d_d" json:"deathday,omitempty"`
	IsHistory           pserialize.PBool           `paradox_field:"hist" json:"ishistory,omitempty"`
	IsFemale            pserialize.PBool           `paradox_field:"fem" json:"isfemale,omitempty"`
	Pi                  pserialize.PBool           `paradox_field:"pi" json:"pi,omitempty"`
	MovedCapital        pserialize.PBool           `paradox_field:"moved_capital" json:"moved_capital,omitempty"`
	IsBastard           pserialize.PBool           `paradox_field:"bstd" json:"bstd,omitempty"`
	ForcedPortrait      pserialize.PBool           `paradox_field:"forced_portrait" json:"forced_portrait,omitempty"`
	ForceHost           pserialize.PBool           `paradox_field:"force_host" json:"force_host,omitempty"`
	InHiding            pserialize.PBool           `paradox_field:"in_hiding" json:"in_hiding,omitempty"`
	OffmapRuler         pserialize.PBool           `paradox_field:"offmap_ruler" json:"offmap_ruler,omitempty"`
	Ee                  pserialize.PBool           `paradox_field:"ee" json:"ee,omitempty"`
	IsPlayer            pserialize.PBool           `paradox_field:"player" json:"isplayer,omitempty"`
	Occluded            pserialize.PBool           `paradox_field:"occluded" json:"occluded,omitempty"`
	Dynnam              pserialize.PBool           `paradox_field:"dynnam" json:"dynnam,omitempty"`
	Ruler               pserialize.PBool           `paradox_field:"ruler" json:"ruler,omitempty"`
	Action              string                     `paradox_field:"action" json:"action,omitempty"`
	Nick                *Nick                      `paradox_field:"nick" paradox_type:"entity" paradox_default_field:"nickname" json:"nick,omitempty"`
	CD                  string                     `paradox_field:"c_d" json:"c_d,omitempty"`
	DNA                 string                     `paradox_field:"dna" json:"dna,omitempty"`
	Government          string                     `paradox_field:"gov" json:"gov,omitempty"`
	CPos                string                     `paradox_field:"cpos" json:"cpos,omitempty"`
	Job                 string                     `paradox_field:"job" json:"job,omitempty"`
	Culture             string                     `paradox_field:"cul" json:"culture,omitempty"`
	GFXCulture          string                     `paradox_field:"g_cul" json:"gfx_culture,omitempty"`
	Religion            string                     `paradox_field:"rel" json:"religion,omitempty"`
	SecretReligion      string                     `paradox_field:"secret_religion" json:"secret_religion,omitempty"`
	Titles              []string                   `paradox_field:"title" paradox_type:"list" json:"titles,omitempty"`
	Vars                map[string]float32         `paradox_field:"vars" json:"vars,omitempty"`
	Fertility           float32                    `paradox_field:"fer" json:"fer,omitempty"`
	Health              float32                    `paradox_field:"health" json:"health,omitempty"`
	Prestige            float32                    `paradox_field:"prs" json:"prestige,omitempty"`
	Piety               float32                    `paradox_field:"piety" json:"piety,omitempty"`
	Wealth              float32                    `paradox_field:"wealth" json:"wealth,omitempty"`
	Emi                 float32                    `paradox_field:"emi" json:"emi,omitempty"`
	Eyi                 float32                    `paradox_field:"eyi" json:"eyi,omitempty"`
	Eypi                float32                    `paradox_field:"eypi" json:"eypi,omitempty"`
	Score               float32                    `paradox_field:"score" json:"score,omitempty"`
	Curinc              float32                    `paradox_field:"curinc" json:"curinc,omitempty"`
	Eme                 float32                    `paradox_field:"eme" json:"eme,omitempty"`
	Curexp              float32                    `paradox_field:"curexp" json:"curexp,omitempty"`
	Father              int                        `paradox_field:"fat" json:"fat,omitempty"`
	RealFather          int                        `paradox_field:"rfat" json:"rfat,omitempty"`
	Mother              int                        `paradox_field:"mot" json:"mot,omitempty"`
	Dynasty             int                        `paradox_field:"dnt" json:"dnt,omitempty"`
	Emp                 int                        `paradox_field:"emp" json:"emp,omitempty"`
	Host                int                        `paradox_field:"host" json:"host,omitempty"`
	Retrat              int                        `paradox_field:"retrat" json:"retrat,omitempty"`
	Guardian            int                        `paradox_field:"guardian" json:"guardian,omitempty"`
	WarTarget           int                        `paradox_field:"war_target" json:"war_target,omitempty"`
	CPosStart           int                        `paradox_field:"cpos_start" json:"cpos_start,omitempty"`
	MovedCapitalMonths  int                        `paradox_field:"moved_capital_months" json:"moved_capital_months,omitempty"`
	Society             int                        `paradox_field:"society" json:"society,omitempty"`
	ACLoc               int                        `paradox_field:"acloc" json:"acloc,omitempty"`
	Rfat                int                        `paradox_field:"rfat" json:"rfat,omitempty"`
	LastObjective       int                        `paradox_field:"last_objective" json:"last_objective,omitempty"`
	Btrh                int                        `paradox_field:"btrh" json:"btrh,omitempty"`
	Regent              int                        `paradox_field:"regent" json:"regent,omitempty"`
	ConsortOf           int                        `paradox_field:"consort_of" json:"consort_of,omitempty"`
	Killer              int                        `paradox_field:"killer" json:"killer,omitempty"`
	Lge                 int                        `paradox_field:"lge" json:"lge,omitempty"`
	Consort             []int                      `paradox_field:"consort" paradox_type:"list" json:"consort,omitempty"`
	Spouse              []int                      `paradox_field:"spouse" paradox_type:"list" json:"spouse,omitempty"`
	Lover               []int                      `paradox_field:"lover" paradox_type:"list" json:"lover,omitempty"`
	Attributes          []int                      `paradox_field:"att" paradox_type:"field_list" json:"att,omitempty"`
	Traits              []int                      `paradox_field:"tr" paradox_type:"field_list" json:"tr,omitempty"`
	OffMapCurrencies    map[int]float32            `paradox_field:"offmap_currencies" json:"offmap_currencies,omitempty"`
	Modifier            []*Modifier                `paradox_field:"md" paradox_type:"list" json:"md,omitempty"`
	Claim               []*Claim                   `paradox_field:"claim" paradox_type:"list" json:"claim,omitempty"`
	Lgr                 *Lgr                       `paradox_field:"lgr" json:"lgr,omitempty"`
	SavedEventTarget    *SavedEventTarget          `paradox_field:"saved_event_target" json:"saved_event_target,omitempty"`
	DMN                 *DMN                       `paradox_field:"dmn" json:"dmn,omitempty"`
	CharacterPlayerData *CharacterPlayerData       `paradox_field:"character_player_data" json:"character_player_data,omitempty"`
	PlayID              int                        `description:"game play id" json:"play_id,omitempty"`
	PlayDate            time.Time                  `description:"game play date" json:"play_date,omitempty"`
}

func processCharacters(saveFile *SaveFile, translations map[string]string) {
	for _, c := range saveFile.Characters {
		c.PlayID = saveFile.PlayThroughID
		c.PlayDate = time.Time(saveFile.Date)

		if c.Name == "" {
			k := fmt.Sprintf("%d#|name", c.ID)
			c.Name = translations[k]
		}

		if c.Nick != nil {
			if c.Nick.Name == "" && c.Nick.NickName != "" {
				c.Nick.Name = translations[c.Nick.NickName]
			}
		}
	}
}
