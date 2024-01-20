package modifier

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type Modifier struct {
	Code                                  string           `paradox_type:"map_key" json:"code,omitempty"`
	Name                                  string           `json:"name,omitempty"`
	Description                           string           `json:"description,omitempty"`
	Major                                 pserialize.PBool `paradox_field:"major" json:"major,omitempty"`
	LocalRevoltRisk                       float32          `paradox_field:"local_revolt_risk" json:"local_revolt_risk,omitempty"`
	LandMorale                            float32          `paradox_field:"land_morale" json:"land_morale,omitempty"`
	MonthlyCharacterPrestige              float32          `paradox_field:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	MonthlyCharacterPiety                 float32          `paradox_field:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	MonthlyCharacterWealth                float32          `paradox_field:"monthly_character_wealth" json:"monthly_character_wealth,omitempty"`
	Fertility                             float32          `paradox_field:"fertility" json:"fertility,omitempty"`
	BuildCostModifier                     float32          `paradox_field:"build_cost_modifier" json:"build_cost_modifier,omitempty"`
	BuildTimeModifier                     float32          `paradox_field:"build_time_modifier" json:"build_time_modifier,omitempty"`
	CastleLevySize                        float32          `paradox_field:"castle_levy_size" json:"castle_levy_size,omitempty"`
	GlobalTaxModifier                     float32          `paradox_field:"global_tax_modifier" json:"global_tax_modifier,omitempty"`
	LocalTaxModifier                      float32          `paradox_field:"local_tax_modifier" json:"local_tax_modifier,omitempty"`
	LocalBuildTimeModifier                float32          `paradox_field:"local_build_time_modifier" json:"local_build_time_modifier,omitempty"`
	TechGrowthModifierCulture             float32          `paradox_field:"tech_growth_modifier_culture" json:"tech_growth_modifier_culture,omitempty"`
	TechGrowthModifierMilitary            float32          `paradox_field:"tech_growth_modifier_military" json:"tech_growth_modifier_military,omitempty"`
	CastleTaxModifier                     float32          `paradox_field:"castle_tax_modifier" json:"castle_tax_modifier,omitempty"`
	LevyReinforceRate                     float32          `paradox_field:"levy_reinforce_rate" json:"levy_reinforce_rate,omitempty"`
	CityTaxModifier                       float32          `paradox_field:"city_tax_modifier" json:"city_tax_modifier,omitempty"`
	TempleTaxModifier                     float32          `paradox_field:"temple_tax_modifier" json:"temple_tax_modifier,omitempty"`
	TribalTaxModifier                     float32          `paradox_field:"tribal_tax_modifier" json:"tribal_tax_modifier,omitempty"`
	NomadTaxModifier                      float32          `paradox_field:"nomad_tax_modifier" json:"nomad_tax_modifier,omitempty"`
	DiseaseDefence                        float32          `paradox_field:"disease_defence" json:"disease_defence,omitempty"`
	LocalBuildCostModifier                float32          `paradox_field:"local_build_cost_modifier" json:"local_build_cost_modifier,omitempty"`
	LevySize                              float32          `paradox_field:"levy_size" json:"levy_size,omitempty"`
	TribalLevySize                        float32          `paradox_field:"tribal_levy_size" json:"tribal_levy_size,omitempty"`
	ManpowerGrowth                        float32          `paradox_field:"manpower_growth" json:"manpower_growth,omitempty"`
	DefensivePlotPowerModifier            float32          `paradox_field:"defensive_plot_power_modifier" json:"defensive_plot_power_modifier,omitempty"`
	GlobalRevoltRisk                      float32          `paradox_field:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	PopulationGrowth                      float32          `paradox_field:"population_growth" json:"population_growth,omitempty"`
	TaxIncome                             float32          `paradox_field:"tax_income" json:"tax_income,omitempty"`
	ArrestChanceModifier                  float32          `paradox_field:"arrest_chance_modifier" json:"arrest_chance_modifier,omitempty"`
	PlotPowerModifier                     float32          `paradox_field:"plot_power_modifier" json:"plot_power_modifier,omitempty"`
	PlotDiscoveryChance                   float32          `paradox_field:"plot_discovery_chance" json:"plot_discovery_chance,omitempty"`
	Health                                float32          `paradox_field:"health" json:"health,omitempty"`
	MaxManpowerMult                       float32          `paradox_field:"max_manpower_mult" json:"max_manpower_mult,omitempty"`
	MaxPopulationMult                     float32          `paradox_field:"max_population_mult" json:"max_population_mult,omitempty"`
	GarrisonSize                          float32          `paradox_field:"garrison_size" json:"garrison_size,omitempty"`
	HealthPenalty                         float32          `paradox_field:"health_penalty" json:"health_penalty,omitempty"`
	GalleysPerc                           float32          `paradox_field:"galleys_perc" json:"galleys_perc,omitempty"`
	MilitaryTechpoints                    float32          `paradox_field:"military_techpoints" json:"military_techpoints,omitempty"`
	CultureTechpoints                     float32          `paradox_field:"culture_techpoints" json:"culture_techpoints,omitempty"`
	EconomyTechpoints                     float32          `paradox_field:"economy_techpoints" json:"economy_techpoints,omitempty"`
	TechGrowthModifier                    float32          `paradox_field:"tech_growth_modifier" json:"tech_growth_modifier,omitempty"`
	AddPietyModifier                      float32          `paradox_field:"add_piety_modifier" json:"add_piety_modifier,omitempty"`
	AddPrestigeModifier                   float32          `paradox_field:"add_prestige_modifier" json:"add_prestige_modifier,omitempty"`
	GlobalLevySize                        float32          `paradox_field:"global_levy_size" json:"global_levy_size,omitempty"`
	CityVassalTaxModifier                 float32          `paradox_field:"city_vassal_tax_modifier" json:"city_vassal_tax_modifier,omitempty"`
	SiegeDefence                          float32          `paradox_field:"siege_defence" json:"siege_defence,omitempty"`
	MultiplicativeTradePostIncomeModifier float32          `paradox_field:"multiplicative_trade_post_income_modifier" json:"multiplicative_trade_post_income_modifier,omitempty"`
	LocalBuildTimeTempleModifier          float32          `paradox_field:"local_build_time_temple_modifier" json:"local_build_time_temple_modifier,omitempty"`
	GlobalMovementSpeed                   float32          `paradox_field:"global_movement_speed" json:"global_movement_speed,omitempty"`
	SiegeSpeed                            float32          `paradox_field:"siege_speed" json:"siege_speed,omitempty"`
	SupplyLimit                           float32          `paradox_field:"supply_limit" json:"supply_limit,omitempty"`
	Attrition                             float32          `paradox_field:"attrition" json:"attrition,omitempty"`
	ReligionFlex                          float32          `paradox_field:"religion_flex" json:"religion_flex,omitempty"`
	CultureFlex                           float32          `paradox_field:"culture_flex" json:"culture_flex,omitempty"`
	LandOrganisation                      float32          `paradox_field:"land_organisation" json:"land_organisation,omitempty"`
	MurderPlotPowerModifier               float32          `paradox_field:"murder_plot_power_modifier" json:"murder_plot_power_modifier,omitempty"`
	Icon                                  int              `paradox_field:"icon" json:"icon,omitempty"`
	Intrigue                              int              `paradox_field:"intrigue" json:"intrigue,omitempty"`
	Martial                               int              `paradox_field:"martial" json:"martial,omitempty"`
	ChurchOpinion                         int              `paradox_field:"church_opinion" json:"church_opinion,omitempty"`
	Stewardship                           int              `paradox_field:"stewardship" json:"stewardship,omitempty"`
	Learning                              int              `paradox_field:"learning" json:"learning,omitempty"`
	GeneralOpinion                        int              `paradox_field:"general_opinion" json:"general_opinion,omitempty"`
	CombatRating                          int              `paradox_field:"combat_rating" json:"combat_rating,omitempty"`
	Tradevalue                            int              `paradox_field:"tradevalue" json:"tradevalue,omitempty"`
	TradeRouteWealth                      int              `paradox_field:"trade_route_wealth" json:"trade_route_wealth,omitempty"`
	TradeRouteValue                       int              `paradox_field:"trade_route_value" json:"trade_route_value,omitempty"`
	TownOpinion                           int              `paradox_field:"town_opinion" json:"town_opinion,omitempty"`
	SexAppealOpinion                      int              `paradox_field:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	CastleOpinion                         int              `paradox_field:"castle_opinion" json:"castle_opinion,omitempty"`
	TechGrowthModifierEconomy             int              `paradox_field:"tech_growth_modifier_economy" json:"tech_growth_modifier_economy,omitempty"`
	SameReligionOpinion                   int              `paradox_field:"same_religion_opinion" json:"same_religion_opinion,omitempty"`
	VassalOpinion                         int              `paradox_field:"vassal_opinion" json:"vassal_opinion,omitempty"`
	TribalOpinion                         int              `paradox_field:"tribal_opinion" json:"tribal_opinion,omitempty"`
	ClanSentiment                         int              `paradox_field:"clan_sentiment" json:"clan_sentiment,omitempty"`
	DiplomacyPenalty                      int              `paradox_field:"diplomacy_penalty" json:"diplomacy_penalty,omitempty"`
	MartialPenalty                        int              `paradox_field:"martial_penalty" json:"martial_penalty,omitempty"`
	StewardshipPenalty                    int              `paradox_field:"stewardship_penalty" json:"stewardship_penalty,omitempty"`
	IntriguePenalty                       int              `paradox_field:"intrigue_penalty" json:"intrigue_penalty,omitempty"`
	LearningPenalty                       int              `paradox_field:"learning_penalty" json:"learning_penalty,omitempty"`
	LiegeOpinion                          int              `paradox_field:"liege_opinion" json:"liege_opinion,omitempty"`
	DynastyOpinion                        int              `paradox_field:"dynasty_opinion" json:"dynasty_opinion,omitempty"`
	SocietyInfluence                      int              `paradox_field:"society_influence" json:"society_influence,omitempty"`
	MonthlyGrace                          int              `paradox_field:"monthly_grace" json:"monthly_grace,omitempty"`
	DuelistOpinion                        int              `paradox_field:"duelist_opinion" json:"duelist_opinion,omitempty"`
	CruelOpinion                          int              `paradox_field:"cruel_opinion" json:"cruel_opinion,omitempty"`
	StrategistOpinion                     int              `paradox_field:"strategist_opinion" json:"strategist_opinion,omitempty"`
	KindOpinion                           int              `paradox_field:"kind_opinion" json:"kind_opinion,omitempty"`
	Diplomacy                             int              `paradox_field:"Diplomacy" json:"Diplomacy,omitempty"`
	WrothOpinion                          int              `paradox_field:"wroth_opinion" json:"wroth_opinion,omitempty"`
	ZealousOpinion                        int              `paradox_field:"zealous_opinion" json:"zealous_opinion,omitempty"`
	DeceitfulOpinion                      int              `paradox_field:"deceitful_opinion" json:"deceitful_opinion,omitempty"`
	AmbitiousOpinion                      int              `paradox_field:"ambitious_opinion" json:"ambitious_opinion,omitempty"`
	DaysOfSupply                          int              `paradox_field:"days_of_supply" json:"days_of_supply,omitempty"`
	BraveOpinion                          int              `paradox_field:"brave_opinion" json:"brave_opinion,omitempty"`
	PatientOpinion                        int              `paradox_field:"patient_opinion" json:"patient_opinion,omitempty"`
	HumbleOpinion                         int              `paradox_field:"humble_opinion" json:"humble_opinion,omitempty"`
	ProudOpinion                          int              `paradox_field:"proud_opinion" json:"proud_opinion,omitempty"`
	CharitableOpinion                     int              `paradox_field:"charitable_opinion" json:"charitable_opinion,omitempty"`
}

func LoadAllModifier(path string) map[string]*Modifier {

	translations := localisation.LoadAllTranslations(path)
	religionPath := filepath.Join(path, "common", "event_modifiers")
	files, err := os.ReadDir(religionPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*Modifier)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(religionPath, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*Modifier](content)

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
		m.Description = translations[m.Code+"_desc"]
	}

	return result
}
