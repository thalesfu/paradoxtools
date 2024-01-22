package building

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type BuildingGroup struct {
	Code        string               `paradox_type:"map_key" json:"code,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Buildings   map[string]*Building `paradox_type:"map" paradox_map_key_pattern:".*"  json:"buildings,omitempty"`
}

type Building struct {
	Code                     string           `paradox_type:"map_key" json:"code,omitempty"`
	Name                     string           `json:"name,omitempty"`
	Description              string           `json:"description,omitempty"`
	Group                    string           `json:"group,omitempty"`
	GroupName                string           `json:"group_name,omitempty"`
	GroupDescription         string           `json:"group_description,omitempty"`
	Port                     pserialize.PBool `paradox_field:"port" json:"port,omitempty"`
	AddNumberToName          pserialize.PBool `paradox_field:"add_number_to_name" json:"add_number_to_name,omitempty"`
	Scouting                 pserialize.PBool `paradox_field:"scouting" json:"scouting,omitempty"`
	Desc                     string           `paradox_field:"desc" json:"desc,omitempty"`
	UpgradesFrom             string           `paradox_field:"upgrades_from" json:"upgrades_from,omitempty"`
	Replaces                 string           `paradox_field:"replaces" json:"replaces,omitempty"`
	ConvertToTribal          string           `paradox_field:"convert_to_tribal" json:"convert_to_tribal,omitempty"`
	ConvertToCity            string           `paradox_field:"convert_to_city" json:"convert_to_city,omitempty"`
	ConvertToCastle          string           `paradox_field:"convert_to_castle" json:"convert_to_castle,omitempty"`
	FortLevel                float32          `paradox_field:"fort_level" json:"fort_level,omitempty"`
	ExtraTechBuildingStart   float32          `paradox_field:"extra_tech_building_start" json:"extra_tech_building_start,omitempty"`
	LevySize                 float32          `paradox_field:"levy_size" json:"levy_size,omitempty"`
	TaxIncome                float32          `paradox_field:"tax_income" json:"tax_income,omitempty"`
	GarrisonSize             float32          `paradox_field:"garrison_size" json:"garrison_size,omitempty"`
	LandMorale               float32          `paradox_field:"land_morale" json:"land_morale,omitempty"`
	LevyReinforceRate        float32          `paradox_field:"levy_reinforce_rate" json:"levy_reinforce_rate,omitempty"`
	LocalRevoltRisk          float32          `paradox_field:"local_revolt_risk" json:"local_revolt_risk,omitempty"`
	LiegePrestige            float32          `paradox_field:"liege_prestige" json:"liege_prestige,omitempty"`
	ArchersOffensive         float32          `paradox_field:"archers_offensive" json:"archers_offensive,omitempty"`
	PikemenDefensive         float32          `paradox_field:"pikemen_defensive" json:"pikemen_defensive,omitempty"`
	HeavyInfantryDefensive   float32          `paradox_field:"heavy_infantry_defensive" json:"heavy_infantry_defensive,omitempty"`
	HorseArchersOffensive    float32          `paradox_field:"horse_archers_offensive" json:"horse_archers_offensive,omitempty"`
	HorseArchersMorale       float32          `paradox_field:"horse_archers_morale" json:"horse_archers_morale,omitempty"`
	KnightsOffensive         float32          `paradox_field:"knights_offensive" json:"knights_offensive,omitempty"`
	LightCavalryOffensive    float32          `paradox_field:"light_cavalry_offensive" json:"light_cavalry_offensive,omitempty"`
	PikemenMorale            float32          `paradox_field:"pikemen_morale" json:"pikemen_morale,omitempty"`
	HeavyInfantryOffensive   float32          `paradox_field:"heavy_infantry_offensive" json:"heavy_infantry_offensive,omitempty"`
	LightCavalryDefensive    float32          `paradox_field:"light_cavalry_defensive" json:"light_cavalry_defensive,omitempty"`
	CamelCavalryDefensive    float32          `paradox_field:"camel_cavalry_defensive" json:"camel_cavalry_defensive,omitempty"`
	LightInfantryOffensive   float32          `paradox_field:"light_infantry_offensive" json:"light_infantry_offensive,omitempty"`
	LightInfantryDefensive   float32          `paradox_field:"light_infantry_defensive" json:"light_infantry_defensive,omitempty"`
	KnightsDefensive         float32          `paradox_field:"knights_defensive" json:"knights_defensive,omitempty"`
	PikemenOffensive         float32          `paradox_field:"pikemen_offensive" json:"pikemen_offensive,omitempty"`
	HeavyInfantryMorale      float32          `paradox_field:"heavy_infantry_morale" json:"heavy_infantry_morale,omitempty"`
	WarElephantsOffensive    float32          `paradox_field:"war_elephants_offensive" json:"war_elephants_offensive,omitempty"`
	WarElephantsDefensive    float32          `paradox_field:"war_elephants_defensive" json:"war_elephants_defensive,omitempty"`
	ArchersDefensive         float32          `paradox_field:"archers_defensive" json:"archers_defensive,omitempty"`
	LightInfantryMorale      float32          `paradox_field:"light_infantry_morale" json:"light_infantry_morale,omitempty"`
	LightCavalryMorale       float32          `paradox_field:"light_cavalry_morale" json:"light_cavalry_morale,omitempty"`
	TechGrowthModifier       float32          `paradox_field:"tech_growth_modifier" json:"tech_growth_modifier,omitempty"`
	CultureTechpoints        float32          `paradox_field:"culture_techpoints" json:"culture_techpoints,omitempty"`
	LiegePiety               float32          `paradox_field:"liege_piety" json:"liege_piety,omitempty"`
	EconomyTechpoints        float32          `paradox_field:"economy_techpoints" json:"economy_techpoints,omitempty"`
	DiseaseDefence           float32          `paradox_field:"disease_defence" json:"disease_defence,omitempty"`
	LocalBuildTimeModifier   float32          `paradox_field:"local_build_time_modifier" json:"local_build_time_modifier,omitempty"`
	LocalBuildCostModifier   float32          `paradox_field:"local_build_cost_modifier" json:"local_build_cost_modifier,omitempty"`
	MilitaryTechpoints       float32          `paradox_field:"military_techpoints" json:"military_techpoints,omitempty"`
	MonthlyCharacterPrestige float32          `paradox_field:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	Fertility                float32          `paradox_field:"fertility" json:"fertility,omitempty"`
	MonthlyCharacterPiety    float32          `paradox_field:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	HorseArchersDefensive    float32          `paradox_field:"horse_archers_defensive" json:"horse_archers_defensive,omitempty"`
	GlobalMovementSpeed      float32          `paradox_field:"global_movement_speed" json:"global_movement_speed,omitempty"`
	PopulationGrowth         float32          `paradox_field:"population_growth" json:"population_growth,omitempty"`
	NomadTaxModifier         float32          `paradox_field:"nomad_tax_modifier" json:"nomad_tax_modifier,omitempty"`
	GlobalSupplyLimit        float32          `paradox_field:"global_supply_limit" json:"global_supply_limit,omitempty"`
	MovedCapitalMonthsMult   float32          `paradox_field:"moved_capital_months_mult" json:"moved_capital_months_mult,omitempty"`
	HordeMaintenenceCost     float32          `paradox_field:"horde_maintenence_cost" json:"horde_maintenence_cost,omitempty"`
	GlobalWinterSupply       float32          `paradox_field:"global_winter_supply" json:"global_winter_supply,omitempty"`
	MaxPopulationMult        float32          `paradox_field:"max_population_mult" json:"max_population_mult,omitempty"`
	GlobalTradeRouteValue    float32          `paradox_field:"global_trade_route_value" json:"global_trade_route_value,omitempty"`
	CamelCavalryOffensive    float32          `paradox_field:"camel_cavalry_offensive" json:"camel_cavalry_offensive,omitempty"`
	TradeRouteWealth         float32          `paradox_field:"trade_route_wealth" json:"trade_route_wealth,omitempty"`
	GoldCost                 int              `paradox_field:"gold_cost" json:"gold_cost,omitempty"`
	BuildTime                int              `paradox_field:"build_time" json:"build_time,omitempty"`
	AiCreationFactor         int              `paradox_field:"ai_creation_factor" json:"ai_creation_factor,omitempty"`
	LightInfantry            int              `paradox_field:"light_infantry" json:"light_infantry,omitempty"`
	Archers                  int              `paradox_field:"archers" json:"archers,omitempty"`
	Retinuesize              int              `paradox_field:"retinuesize" json:"retinuesize,omitempty"`
	HeavyInfantry            int              `paradox_field:"heavy_infantry" json:"heavy_infantry,omitempty"`
	Pikemen                  int              `paradox_field:"pikemen" json:"pikemen,omitempty"`
	LightCavalry             int              `paradox_field:"light_cavalry" json:"light_cavalry,omitempty"`
	Knights                  int              `paradox_field:"knights" json:"knights,omitempty"`
	CourtSizeModifier        int              `paradox_field:"court_size_modifier" json:"court_size_modifier,omitempty"`
	Galleys                  int              `paradox_field:"galleys" json:"galleys,omitempty"`
	HorseArchers             int              `paradox_field:"horse_archers" json:"horse_archers,omitempty"`
	CamelCavalry             int              `paradox_field:"camel_cavalry" json:"camel_cavalry,omitempty"`
	WarElephants             int              `paradox_field:"war_elephants" json:"war_elephants,omitempty"`
	Tradevalue               int              `paradox_field:"tradevalue" json:"tradevalue,omitempty"`
	MaxTradeposts            int              `paradox_field:"max_tradeposts" json:"max_tradeposts,omitempty"`
	Stewardship              int              `paradox_field:"stewardship" json:"stewardship,omitempty"`
	Diplomacy                int              `paradox_field:"diplomacy" json:"diplomacy,omitempty"`
	Martial                  int              `paradox_field:"martial" json:"martial,omitempty"`
	ChurchOpinion            int              `paradox_field:"church_opinion" json:"church_opinion,omitempty"`
	Learning                 int              `paradox_field:"learning" json:"learning,omitempty"`
	Intrigue                 int              `paradox_field:"intrigue" json:"intrigue,omitempty"`
	HospitalLevel            int              `paradox_field:"hospital_level" json:"hospital_level,omitempty"`
	MaxPopulation            int              `paradox_field:"max_population" json:"max_population,omitempty"`
	CommanderLimit           int              `paradox_field:"commander_limit" json:"commander_limit,omitempty"`
	ClanSentiment            int              `paradox_field:"clan_sentiment" json:"clan_sentiment,omitempty"`
	GlobalTradeRouteWealth   int              `paradox_field:"global_trade_route_wealth" json:"global_trade_route_wealth,omitempty"`
	GlobalTradevalue         int              `paradox_field:"global_tradevalue" json:"global_tradevalue,omitempty"`
	AiFeudalModifier         int              `paradox_field:"ai_feudal_modifier" json:"ai_feudal_modifier,omitempty"`
	AiRepublicModifier       int              `paradox_field:"ai_republic_modifier" json:"ai_republic_modifier,omitempty"`
	PietyCost                int              `paradox_field:"piety_cost" json:"piety_cost,omitempty"`
	PrestigeCost             int              `paradox_field:"prestige_cost" json:"prestige_cost,omitempty"`
}

func LoadAllBuildings(path string) map[string]*BuildingGroup {

	translations := localisation.LoadAllTranslations(path)
	buildingPath := filepath.Join(path, "common", "buildings")
	files, err := os.ReadDir(buildingPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*BuildingGroup)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(buildingPath, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*BuildingGroup](content)

				if o {
					for k, v := range *ts {
						if g, ok := result[k]; ok {
							for kk, vv := range g.Buildings {
								g.Buildings[kk] = vv
							}
						} else {
							result[k] = v

						}

					}
				}
			}
		}
	}

	for _, buildingGroup := range result {
		buildingGroup.Name = translations[buildingGroup.Code]
		buildingGroup.Description = translations[buildingGroup.Code+"_desc"]

		for _, building := range buildingGroup.Buildings {
			building.Name = translations[building.Code]
			building.Description = translations[building.Code+"_desc"]
			building.Group = buildingGroup.Code
			building.GroupName = buildingGroup.Name
			building.GroupDescription = buildingGroup.Description
		}
	}

	return result
}
