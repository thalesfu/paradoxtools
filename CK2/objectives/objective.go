package objectives

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type Objective struct {
	Code                  string           `paradox_type:"map_key" json:"code,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Description           string           `json:"description,omitempty"`
	CanCancel             pserialize.PBool `paradox_field:"can_cancel" json:"can_cancel,omitempty"`
	ExpectationOfLiege    pserialize.PBool `paradox_field:"expectation_of_liege" json:"expectation_of_liege,omitempty"`
	AiCapitalKingdomFocus pserialize.PBool `paradox_field:"ai_capital_kingdom_focus" json:"ai_capital_kingdom_focus,omitempty"`
	Exclusive             pserialize.PBool `paradox_field:"exclusive" json:"exclusive,omitempty"`
	RelHeadLoyalist       pserialize.PBool `paradox_field:"rel_head_loyalist" json:"rel_head_loyalist,omitempty"`
	CancelOnLeaderDeath   pserialize.PBool `paradox_field:"cancel_on_leader_death" json:"cancel_on_leader_death,omitempty"`
	MilitaryPlot          pserialize.PBool `paradox_field:"military_plot" json:"military_plot,omitempty"`
	VassalRankPlot        pserialize.PBool `paradox_field:"vassal_rank_plot" json:"vassal_rank_plot,omitempty"`
	IntriguePlot          pserialize.PBool `paradox_field:"intrigue_plot" json:"intrigue_plot,omitempty"`
	MurderPlot            pserialize.PBool `paradox_field:"murder_plot" json:"murder_plot,omitempty"`
	VassalIntriguePlot    pserialize.PBool `paradox_field:"vassal_intrigue_plot" json:"vassal_intrigue_plot,omitempty"`
	Type                  string           `paradox_field:"type" json:"type,omitempty"`
	Text                  string           `paradox_field:"text" json:"text,omitempty"`
	Fertility             float32          `paradox_field:"fertility" json:"fertility,omitempty"`
	Aggression            float32          `paradox_field:"aggression" json:"aggression,omitempty"`
	GlobalRevoltRisk      float32          `paradox_field:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	WarningLevel          float32          `paradox_field:"warning_level" json:"warning_level,omitempty"`
	PlotPowerModifier     float32          `paradox_field:"plot_power_modifier" json:"plot_power_modifier,omitempty"`
	Health                float32          `paradox_field:"health" json:"health,omitempty"`
	Stewardship           int              `paradox_field:"stewardship" json:"stewardship,omitempty"`
	TownOpinion           int              `paradox_field:"town_opinion" json:"town_opinion,omitempty"`
	Intrigue              int              `paradox_field:"intrigue" json:"intrigue,omitempty"`
	SexAppealOpinion      int              `paradox_field:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	Martial               int              `paradox_field:"martial" json:"martial,omitempty"`
	CombatRating          int              `paradox_field:"combat_rating" json:"combat_rating,omitempty"`
	Diplomacy             int              `paradox_field:"diplomacy" json:"diplomacy,omitempty"`
	Learning              int              `paradox_field:"learning" json:"learning,omitempty"`
	ChurchOpinion         int              `paradox_field:"church_opinion" json:"church_opinion,omitempty"`
}

func LoadAllObjective(path string) map[string]*Objective {

	translations := localisation.LoadAllTranslations(path)
	dir := filepath.Join(path, "common", "objectives")
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*Objective)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(dir, filename)

			content, ok := golangutils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*Objective](content)

				if o {
					for k, v := range *ts {
						result[k] = v
					}
				}
			}
		}
	}

	for _, m := range result {
		m.Name = translations[m.Code]

		if m.Name == "" {
			m.Name = translations[m.Code+"_title"]
		}

		m.Description = translations[m.Code+"_desc"]
	}

	return result
}
