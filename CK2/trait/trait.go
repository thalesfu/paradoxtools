package trait

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type CommandModifier struct {
	Terrain        string  `paradox_field:"terrain" json:"terrain,omitempty"`
	MoraleOffence  float32 `paradox_field:"morale_offence" json:"morale_offence,omitempty"`
	Pursue         float32 `paradox_field:"pursue" json:"pursue,omitempty"`
	Defence        float32 `paradox_field:"defence" json:"defence,omitempty"`
	MoraleDefence  float32 `paradox_field:"morale_defence" json:"morale_defence,omitempty"`
	LightInfantry  float32 `paradox_field:"light_infantry" json:"light_infantry,omitempty"`
	HeavyInfantry  float32 `paradox_field:"heavy_infantry" json:"heavy_infantry,omitempty"`
	Cavalry        float32 `paradox_field:"cavalry" json:"cavalry,omitempty"`
	Random         float32 `paradox_field:"random" json:"random,omitempty"`
	Speed          float32 `paradox_field:"speed" json:"speed,omitempty"`
	Retreat        float32 `paradox_field:"retreat" json:"retreat,omitempty"`
	Damage         float32 `paradox_field:"damage" json:"damage,omitempty"`
	Center         float32 `paradox_field:"center" json:"center,omitempty"`
	Flank          float32 `paradox_field:"flank" json:"flank,omitempty"`
	Siege          float32 `paradox_field:"siege" json:"siege,omitempty"`
	ReligiousEnemy float32 `paradox_field:"religious_enemy" json:"religious_enemy,omitempty"`
	NarrowFlank    float32 `paradox_field:"narrow_flank" json:"narrow_flank,omitempty"`
	WarElephants   float32 `paradox_field:"war_elephants" json:"war_elephants,omitempty"`
	WinterCombat   float32 `paradox_field:"winter_combat" json:"winter_combat,omitempty"`
	LightCavalry   float32 `paradox_field:"light_cavalry" json:"light_cavalry,omitempty"`
	Knights        float32 `paradox_field:"knights" json:"knights,omitempty"`
	WinterSupply   int     `paradox_field:"winter_supply" json:"winter_supply,omitempty"`
}

type Trait struct {
	Code                            string           `paradox_type:"map_key" json:"code,omitempty"`
	ID                              int              `paradox_type:"map_index" json:"id,omitempty"`
	Name                            string           `json:"name,omitempty"`
	Description                     string           `json:"description,omitempty"`
	Education                       pserialize.PBool `paradox_field:"education" json:"education,omitempty"`
	IsHealth                        pserialize.PBool `paradox_field:"is_health" json:"is_health,omitempty"`
	IsIllness                       pserialize.PBool `paradox_field:"is_illness" json:"is_illness,omitempty"`
	Customizer                      pserialize.PBool `paradox_field:"customizer" json:"customizer,omitempty"`
	SuccessionGfx                   pserialize.PBool `paradox_field:"succession_gfx" json:"succession_gfx,omitempty"`
	Incapacitating                  pserialize.PBool `paradox_field:"incapacitating" json:"incapacitating,omitempty"`
	IsEpidemic                      pserialize.PBool `paradox_field:"is_epidemic" json:"is_epidemic,omitempty"`
	Religious                       pserialize.PBool `paradox_field:"religious" json:"religious,omitempty"`
	Random                          pserialize.PBool `paradox_field:"random" json:"random,omitempty"`
	Inbred                          pserialize.PBool `paradox_field:"inbred" json:"inbred,omitempty"`
	Lifestyle                       pserialize.PBool `paradox_field:"lifestyle" json:"lifestyle,omitempty"`
	Personality                     pserialize.PBool `paradox_field:"personality" json:"personality,omitempty"`
	Vice                            pserialize.PBool `paradox_field:"vice" json:"vice,omitempty"`
	Virtue                          pserialize.PBool `paradox_field:"virtue" json:"virtue,omitempty"`
	Leader                          pserialize.PBool `paradox_field:"leader" json:"leader,omitempty"`
	Cached                          pserialize.PBool `paradox_field:"cached" json:"cached,omitempty"`
	Pilgrimage                      pserialize.PBool `paradox_field:"pilgrimage" json:"pilgrimage,omitempty"`
	Agnatic                         pserialize.PBool `paradox_field:"agnatic" json:"agnatic,omitempty"`
	CannotMarry                     pserialize.PBool `paradox_field:"cannot_marry" json:"cannot_marry,omitempty"`
	CannotInherit                   pserialize.PBool `paradox_field:"cannot_inherit" json:"cannot_inherit,omitempty"`
	Blinding                        pserialize.PBool `paradox_field:"blinding" json:"blinding,omitempty"`
	RebelInherited                  pserialize.PBool `paradox_field:"rebel_inherited" json:"rebel_inherited,omitempty"`
	ToleratesChristian              pserialize.PBool `paradox_field:"tolerates_christian" json:"tolerates_christian,omitempty"`
	ToleratesMuslim                 pserialize.PBool `paradox_field:"tolerates_muslim" json:"tolerates_muslim,omitempty"`
	ToleratesPaganGroup             pserialize.PBool `paradox_field:"tolerates_pagan_group" json:"tolerates_pagan_group,omitempty"`
	ToleratesZoroastrianGroup       pserialize.PBool `paradox_field:"tolerates_zoroastrian_group" json:"tolerates_zoroastrian_group,omitempty"`
	ToleratesJewishGroup            pserialize.PBool `paradox_field:"tolerates_jewish_group" json:"tolerates_jewish_group,omitempty"`
	ToleratesIndianGroup            pserialize.PBool `paradox_field:"tolerates_indian_group" json:"tolerates_indian_group,omitempty"`
	InHiding                        pserialize.PBool `paradox_field:"in_hiding" json:"in_hiding,omitempty"`
	Childhood                       pserialize.PBool `paradox_field:"childhood" json:"childhood,omitempty"`
	CanHoldTitles                   pserialize.PBool `paradox_field:"can_hold_titles" json:"can_hold_titles,omitempty"`
	IsSymptom                       pserialize.PBool `paradox_field:"is_symptom" json:"is_symptom,omitempty"`
	Immortal                        pserialize.PBool `paradox_field:"immortal" json:"immortal,omitempty"`
	Hidden                          pserialize.PBool `paradox_field:"hidden" json:"hidden,omitempty"`
	HiddenFromOthers                pserialize.PBool `paradox_field:"hidden_from_others" json:"hidden_from_others,omitempty"`
	Attribute                       string           `paradox_field:"attribute" json:"attribute,omitempty"`
	MaleInsultAdj                   string           `paradox_field:"male_insult_adj" json:"male_insult_adj,omitempty"`
	FemaleInsultAdj                 string           `paradox_field:"female_insult_adj" json:"female_insult_adj,omitempty"`
	ChildInsultAdj                  string           `paradox_field:"child_insult_adj" json:"child_insult_adj,omitempty"`
	MaleComplimentAdj               string           `paradox_field:"male_compliment_adj" json:"male_compliment_adj,omitempty"`
	FemaleComplimentAdj             string           `paradox_field:"female_compliment_adj" json:"female_compliment_adj,omitempty"`
	ChildComplimentAdj              string           `paradox_field:"child_compliment_adj" json:"child_compliment_adj,omitempty"`
	MaleInsult                      string           `paradox_field:"male_insult" json:"male_insult,omitempty"`
	FemaleInsult                    string           `paradox_field:"female_insult" json:"female_insult,omitempty"`
	ChildInsult                     string           `paradox_field:"child_insult" json:"child_insult,omitempty"`
	MaleCompliment                  string           `paradox_field:"male_compliment" json:"male_compliment,omitempty"`
	FemaleCompliment                string           `paradox_field:"female_compliment" json:"female_compliment,omitempty"`
	IsTribal                        string           `paradox_field:"is_tribal" json:"is_tribal,omitempty"`
	ChildCompliment                 string           `paradox_field:"child_compliment" json:"child_compliment,omitempty"`
	ReligionGroup                   string           `paradox_field:"religion_group" json:"religion_group,omitempty"`
	Terrain                         string           `paradox_field:"terrain" json:"terrain,omitempty"`
	Religion                        string           `paradox_field:"religion" json:"religion,omitempty"`
	HasBloodlineFlag                string           `paradox_field:"has_bloodline_flag" json:"has_bloodline_flag,omitempty"`
	IsRuler                         string           `paradox_field:"is_ruler" json:"is_ruler,omitempty"`
	IsFemale                        string           `paradox_field:"is_female" json:"is_female,omitempty"`
	IsTheocracy                     string           `paradox_field:"is_theocracy" json:"is_theocracy,omitempty"`
	ControlsReligion                string           `paradox_field:"controls_religion" json:"controls_religion,omitempty"`
	ReligiousBranch                 string           `paradox_field:"religious_branch" json:"religious_branch,omitempty"`
	Prisoner                        string           `paradox_field:"prisoner" json:"prisoner,omitempty"`
	Race                            string           `paradox_field:"race" json:"race,omitempty"`
	HasReligionFeature              string           `paradox_field:"has_religion_feature" json:"has_religion_feature,omitempty"`
	Character                       string           `paradox_field:"character" json:"character,omitempty"`
	SocietyMemberOf                 string           `paradox_field:"society_member_of" json:"society_member_of,omitempty"`
	IsCloseRelative                 string           `paradox_field:"is_close_relative" json:"is_close_relative,omitempty"`
	Trait                           string           `paradox_field:"trait" json:"trait,omitempty"`
	Ai                              string           `paradox_field:"ai" json:"ai,omitempty"`
	IsNomadic                       string           `paradox_field:"is_nomadic" json:"is_nomadic,omitempty"`
	HasDharmicReligionTrigger       string           `paradox_field:"has_dharmic_religion_trigger" json:"has_dharmic_religion_trigger,omitempty"`
	Culture                         string           `paradox_field:"culture" json:"culture,omitempty"`
	HasCharacterFlag                string           `paradox_field:"has_character_flag" json:"has_character_flag,omitempty"`
	GraphicalCulture                string           `paradox_field:"graphical_culture" json:"graphical_culture,omitempty"`
	Fertility                       float32          `paradox_field:"fertility" json:"fertility,omitempty"`
	Health                          float32          `paradox_field:"health" json:"health,omitempty"`
	FertilityPenalty                float32          `paradox_field:"fertility_penalty" json:"fertility_penalty,omitempty"`
	HealthPenalty                   float32          `paradox_field:"health_penalty" json:"health_penalty,omitempty"`
	MonthlyCharacterPiety           float32          `paradox_field:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	GlobalTaxModifier               float32          `paradox_field:"global_tax_modifier" json:"global_tax_modifier,omitempty"`
	MonthlyCharacterPrestige        float32          `paradox_field:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	MonthlyCharacterWealth          float32          `paradox_field:"monthly_character_wealth" json:"monthly_character_wealth,omitempty"`
	WonderBuildCostModifier         float32          `paradox_field:"wonder_build_cost_modifier" json:"wonder_build_cost_modifier,omitempty"`
	WonderBuildTimeModifier         float32          `paradox_field:"wonder_build_time_modifier" json:"wonder_build_time_modifier,omitempty"`
	GlobalLevySize                  float32          `paradox_field:"global_levy_size" json:"global_levy_size,omitempty"`
	MaxManpowerMult                 float32          `paradox_field:"max_manpower_mult" json:"max_manpower_mult,omitempty"`
	GlobalRevoltRisk                float32          `paradox_field:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	Attrition                       float32          `paradox_field:"attrition" json:"attrition,omitempty"`
	Intrigue                        int              `paradox_field:"intrigue" json:"intrigue,omitempty"`
	Stewardship                     int              `paradox_field:"stewardship" json:"stewardship,omitempty"`
	CombatRating                    int              `paradox_field:"combat_rating" json:"combat_rating,omitempty"`
	Martial                         int              `paradox_field:"martial" json:"martial,omitempty"`
	Diplomacy                       int              `paradox_field:"diplomacy" json:"diplomacy,omitempty"`
	Learning                        int              `paradox_field:"learning" json:"learning,omitempty"`
	LeadershipTraits                int              `paradox_field:"leadership_traits" json:"leadership_traits,omitempty"`
	AiZeal                          int              `paradox_field:"ai_zeal" json:"ai_zeal,omitempty"`
	VassalOpinion                   int              `paradox_field:"vassal_opinion" json:"vassal_opinion,omitempty"`
	SexAppealOpinion                int              `paradox_field:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	SameOpinion                     int              `paradox_field:"same_opinion" json:"same_opinion,omitempty"`
	AiRationality                   int              `paradox_field:"ai_rationality" json:"ai_rationality,omitempty"`
	DiplomacyPenalty                int              `paradox_field:"diplomacy_penalty" json:"diplomacy_penalty,omitempty"`
	StewardshipPenalty              int              `paradox_field:"stewardship_penalty" json:"stewardship_penalty,omitempty"`
	MartialPenalty                  int              `paradox_field:"martial_penalty" json:"martial_penalty,omitempty"`
	IntriguePenalty                 int              `paradox_field:"intrigue_penalty" json:"intrigue_penalty,omitempty"`
	LearningPenalty                 int              `paradox_field:"learning_penalty" json:"learning_penalty,omitempty"`
	InheritChance                   int              `paradox_field:"inherit_chance" json:"inherit_chance,omitempty"`
	GeneralOpinion                  int              `paradox_field:"general_opinion" json:"general_opinion,omitempty"`
	ChurchOpinion                   int              `paradox_field:"church_opinion" json:"church_opinion,omitempty"`
	SameOpinionIfSameReligion       int              `paradox_field:"same_opinion_if_same_religion" json:"same_opinion_if_same_religion,omitempty"`
	TwinOpinion                     int              `paradox_field:"twin_opinion" json:"twin_opinion,omitempty"`
	SpouseOpinion                   int              `paradox_field:"spouse_opinion" json:"spouse_opinion,omitempty"`
	SameReligionOpinion             int              `paradox_field:"same_religion_opinion" json:"same_religion_opinion,omitempty"`
	DynastyOpinion                  int              `paradox_field:"dynasty_opinion" json:"dynasty_opinion,omitempty"`
	RulerDesignerCost               int              `paradox_field:"ruler_designer_cost" json:"ruler_designer_cost,omitempty"`
	Birth                           int              `paradox_field:"birth" json:"birth,omitempty"`
	BothParentHasTraitInheritChance int              `paradox_field:"both_parent_has_trait_inherit_chance" json:"both_parent_has_trait_inherit_chance,omitempty"`
	TribalOpinion                   int              `paradox_field:"tribal_opinion" json:"tribal_opinion,omitempty"`
	ChristianChurchOpinion          int              `paradox_field:"christian_church_opinion" json:"christian_church_opinion,omitempty"`
	OppositeOpinion                 int              `paradox_field:"opposite_opinion" json:"opposite_opinion,omitempty"`
	AiHonor                         int              `paradox_field:"ai_honor" json:"ai_honor,omitempty"`
	AiGreed                         int              `paradox_field:"ai_greed" json:"ai_greed,omitempty"`
	AiAmbition                      int              `paradox_field:"ai_ambition" json:"ai_ambition,omitempty"`
	LiegeOpinion                    int              `paradox_field:"liege_opinion" json:"liege_opinion,omitempty"`
	AmbitionOpinion                 int              `paradox_field:"ambition_opinion" json:"ambition_opinion,omitempty"`
	InfidelOpinion                  int              `paradox_field:"infidel_opinion" json:"infidel_opinion,omitempty"`
	MuslimOpinion                   int              `paradox_field:"muslim_opinion" json:"muslim_opinion,omitempty"`
	ZoroastrianOpinion              int              `paradox_field:"zoroastrian_opinion" json:"zoroastrian_opinion,omitempty"`
	NorsePaganOpinion               int              `paradox_field:"norse_pagan_opinion" json:"norse_pagan_opinion,omitempty"`
	NorsePaganReformedOpinion       int              `paradox_field:"norse_pagan_reformed_opinion" json:"norse_pagan_reformed_opinion,omitempty"`
	CasteTier                       int              `paradox_field:"caste_tier" json:"caste_tier,omitempty"`
	PaganGroupOpinion               int              `paradox_field:"pagan_group_opinion" json:"pagan_group_opinion,omitempty"`
	TaoistOpinion                   int              `paradox_field:"taoist_opinion" json:"taoist_opinion,omitempty"`
	DaysOfSupply                    int              `paradox_field:"days_of_supply" json:"days_of_supply,omitempty"`
	MonthlyGrace                    int              `paradox_field:"monthly_grace" json:"monthly_grace,omitempty"`
	ChristianOpinion                int              `paradox_field:"christian_opinion" json:"christian_opinion,omitempty"`
	TraitEffectCaptureCommanders    int              `paradox_field:"trait_effect_capture_commanders" json:"trait_effect_capture_commanders,omitempty"`
	IndianGroupOpinion              int              `paradox_field:"indian_group_opinion" json:"indian_group_opinion,omitempty"`
	JewishGroupOpinion              int              `paradox_field:"jewish_group_opinion" json:"jewish_group_opinion,omitempty"`
	ZoroastrianGroupOpinion         int              `paradox_field:"zoroastrian_group_opinion" json:"zoroastrian_group_opinion,omitempty"`
	CastleOpinion                   int              `paradox_field:"castle_opinion" json:"castle_opinion,omitempty"`
	TownOpinion                     int              `paradox_field:"town_opinion" json:"town_opinion,omitempty"`
	Opposites                       []string         `paradox_field:"opposites" paradox_type:"field_list" json:"opposites,omitempty"`
	CommandModifier                 *CommandModifier `paradox_field:"command_modifier" json:"command_modifier,omitempty"`
}

func LoadAllTraits(path string) map[string]*Trait {

	translations := localisation.LoadAllTranslations(path)
	religionPath := filepath.Join(path, "common", "traits")
	files, err := os.ReadDir(religionPath)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*Trait)

	traitCount := 0

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(religionPath, filename)

			content, ok := golangutils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*Trait](content)

				if o {
					for k, v := range *ts {
						v.ID = traitCount + v.ID
						result[k] = v
					}
					traitCount += len(*ts)
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
