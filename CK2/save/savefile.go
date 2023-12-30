package save

import (
	"github.com/thalesfu/paradoxtools/utils/pserialize"
)

type SaveFile struct {
	Version         string                                `paradox_field:"version" json:"version,omitempty"`
	Date            pserialize.Year                       `paradox_field:"date" json:"date,omitempty"`
	Player          *Player                               `paradox_field:"player" json:"player,omitempty"`
	PlayerRealm     string                                `paradox_field:"player_realm" json:"player_realm,omitempty"`
	PlayerName      string                                `paradox_field:"player_name" json:"player_name,omitempty"`
	PlayerAge       int                                   `paradox_field:"player_age" json:"player_age,omitempty"`
	PlayerPortrait  *PlayerPortrait                       `paradox_field:"player_portrait" json:"player_portrait,omitempty"`
	IsZeusSave      pserialize.PBool                      `paradox_field:"is_zeus_save" json:"is_zeus_save,omitempty"`
	GameRules       *GameRules                            `paradox_field:"game_rules" json:"game_rules,omitempty"`
	GameSpeed       int                                   `paradox_field:"game_speed" json:"game_speed,omitempty"`
	MapMode         int                                   `paradox_field:"mapmode" json:"mapmode,omitempty"`
	PlayThroughID   int                                   `paradox_field:"playthrough_id" json:"playthrough_id,omitempty"`
	DynastyTitle    []*Title                              `paradox_field:"dyn_title" paradox_type:"list" json:"dyn_title,omitempty"`
	Unit            int                                   `paradox_field:"unit" json:"unit,omitempty"`
	SubUnit         int                                   `paradox_field:"sub_unit" json:"sub_unit,omitempty"`
	StartDate       pserialize.Year                       `paradox_field:"start_date" json:"start_date,omitempty"`
	Flags           *Flag                                 `paradox_field:"flags" json:"flags,omitempty"`
	Vars            *Vars                                 `paradox_field:"vars" json:"vars,omitempty"`
	Dynasties       map[int]*Dynasty                      `paradox_field:"dynasties" json:"dynasties,omitempty"`
	Titles          map[string]*Title                     `paradox_field:"title" json:"title,omitempty"`
	Characters      map[int]*Character                    `paradox_field:"character" json:"character,omitempty"`
	DelayedEvents   []*Event                              `paradox_field:"delayed_event" paradox_type:"list" json:"delayed_event,omitempty"`
	Relations       map[string]map[string]*PersonRelation `paradox_field:"relation" json:"relation,omitempty"`
	ID              *IDEntity                             `paradox_field:"id" json:"id,omitempty"`
	Religions       map[string]*Religion                  `paradox_field:"religion" json:"religion,omitempty"`
	Provinces       map[int]*Province                     `paradox_field:"provinces" json:"provinces,omitempty"`
	Bloodline       map[int]*Bloodline                    `paradox_field:"bloodline" json:"bloodline,omitempty"`
	ActiveAmbitions []*ActiveEvent                        `paradox_field:"active_ambition" paradox_type:"list" json:"active_ambition,omitempty"`
	ActiveFocuses   []*ActiveEvent                        `paradox_field:"active_focus" paradox_type:"list" json:"active_focus,omitempty"`
	Nomads          map[string]*Nomad                     `paradox_field:"nomad" json:"nomad,omitempty"`
	Combat          *Combat                               `paradox_field:"combat" json:"combat,omitempty"`
	War             *War                                  `paradox_field:"war" json:"war,omitempty"`
}

type IDEntity struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}

type ChronicleEntry struct {
	Text            string `paradox_field:"text" json:"text,omitempty"`
	Picture         string `paradox_field:"picture" json:"picture,omitempty"`
	Portrait        int    `paradox_field:"portrait" json:"portrait,omitempty"`
	PortraitCulture string `paradox_field:"portrait_culture" json:"portrait_culture,omitempty"`
	Type            string `paradox_field:"type" json:"type,omitempty"`
}

type ChronicleChapter struct {
	Chronicles []*ChronicleEntry `paradox_field:"chronicle_entry" paradox_type:"list" json:"chronicle_entry,omitempty"`
	Year       int               `paradox_field:"year" json:"year,omitempty"`
}

type Chronicle struct {
	ChronicleChapters []*ChronicleChapter `paradox_field:"chronicle_chapter" paradox_type:"list" json:"chronicle_chapter,omitempty"`
	Character         int                 `paradox_field:"character" json:"character,omitempty"`
}

type ChronicleCollection struct {
	Chronicle         *Chronicle       `paradox_field:"chronicle" json:"chronicle,omitempty"`
	ChroniclePosition int              `paradox_field:"chronicle_position" json:"chronicle_position,omitempty"`
	ChapterPosition   int              `paradox_field:"chapter_position" json:"chapter_position,omitempty"`
	EntryPosition     int              `paradox_field:"entry_position" json:"entry_position,omitempty"`
	ChronicleIconLit  pserialize.PBool `paradox_field:"chronicle_icon_lit" json:"chronicle_icon_lit,omitempty"`
}

type CharacterPlayerData struct {
	ChronicleCollection         *ChronicleCollection `paradox_field:"chronicle_collection" json:"chronicle_collection,omitempty"`
	SocietyShowInterestCooldown pserialize.Year      `paradox_field:"society_show_interest_cooldown" json:"society_show_interest_cooldown,omitempty"`
	JoinSocietyCooldown         pserialize.Year      `paradox_field:"join_society_cooldown" json:"join_society_cooldown,omitempty"`
	Telws                       int                  `paradox_field:"telws" json:"telws,omitempty"`
	Telbc                       int                  `paradox_field:"telbc" json:"telbc,omitempty"`
	Telde                       int                  `paradox_field:"telde" json:"telde,omitempty"`
	Telld                       pserialize.PBool     `paradox_field:"telld" json:"telld,omitempty"`
	Telsc                       pserialize.PBool     `paradox_field:"telsc" json:"telsc,omitempty"`
}

type DMN struct {
	Capital                 string      `paradox_field:"capital" json:"capital,omitempty"`
	Primary                 *Title      `paradox_field:"primary" paradox_type:"map_value" paradox_map_name:"title" json:"primary,omitempty"`
	LiegeTroops             *ArmyDetail `paradox_field:"liege_troops" json:"liege_troops,omitempty"`
	RaisedLiegeTroops       []int       `paradox_field:"raised_liege_troops" paradox_type:"field_list" json:"raised_liege_troops,omitempty"`
	MyLiegelevyContribution int         `paradox_field:"my_liegelevy_contribution" json:"my_liegelevy_contribution,omitempty"`
	PeaceMonths             int         `paradox_field:"peace_months" json:"peace_months,omitempty"`
}

type Lgr struct {
	LastMonthIncomeTable  []float32 `paradox_field:"lastmonthincometable" paradox_type:"field_list" json:"lastmonthincometable,omitempty"`
	LastMonthExpenseTable []float32 `paradox_field:"lastmonthexpensetable" paradox_type:"field_list" json:"lastmonthexpensetable,omitempty"`
	LastMonthIncome       float32   `paradox_field:"lastmonthincome" json:"lastmonthincome,omitempty"`
	LastMonthExpense      float32   `paradox_field:"lastmonthexpense" json:"lastmonthexpense,omitempty"`
}

type Claim struct {
	Title   *Title           `paradox_field:"title" paradox_type:"map_value" paradox_map_name:"title" json:"title,omitempty"`
	Pressed pserialize.PBool `paradox_field:"pressed" json:"pressed,omitempty"`
}

type Modifier struct {
	Modifier string           `paradox_field:"modifier" json:"modifier,omitempty"`
	Date     pserialize.Year  `paradox_field:"date" json:"date,omitempty"`
	Hidden   pserialize.PBool `paradox_field:"hidden" json:"hidden,omitempty"`
}

type Dynasty struct {
	ID            int              `paradox_type:"map_key" json:"id,omitempty"`
	Name          string           `paradox_field:"name" json:"name,omitempty"`
	Culture       string           `paradox_field:"culture" json:"culture,omitempty"`
	Religion      string           `paradox_field:"religion" json:"religion,omitempty"`
	CoatOfArms    *CoatOfArms      `paradox_field:"coat_of_arms" json:"coat_of_arms,omitempty"`
	SetCoatOfArms pserialize.PBool `paradox_field:"set_coat_of_arms" json:"set_coat_of_arms,omitempty"`
}

type Vars struct {
	GlobalAmountOfTimesChinaFamine             float32 `paradox_field:"global_amount_of_times_china_famine" json:"global_amount_of_times_china_famine,omitempty"`
	GlobalDeathsInBattle                       float32 `paradox_field:"global_deaths_in_battle" json:"global_deaths_in_battle,omitempty"`
	GlobalRaidingAdventurerSpawnByDisplacement float32 `paradox_field:"global_raiding_adventurer_spawn_by_displacement" json:"global_raiding_adventurer_spawn_by_displacement,omitempty"`
	GlobalTempPolicyYears                      float32 `paradox_field:"global_temp_policy_years" json:"global_temp_policy_years,omitempty"`
	PhysiqueVariable                           float32 `paradox_field:"physique_variable" json:"physique_variable,omitempty"`
}

type Flag struct {
	AchievementGetMarried        pserialize.Year `paradox_field:"achievement_get_married" json:"achievement_get_married,omitempty"`
	AvarKhaganateRenamed         pserialize.Year `paradox_field:"avar_khaganate_renamed" json:"avar_khaganate_renamed,omitempty"`
	BerniciaRenamed              pserialize.Year `paradox_field:"bernicia_renamed" json:"bernicia_renamed,omitempty"`
	CGStudyMilitary              pserialize.Year `paradox_field:"cg_study_military" json:"cg_study_military,omitempty"`
	FlagCrownsDelivered          pserialize.Year `paradox_field:"flag_crowns_delivered" json:"flag_crowns_delivered,omitempty"`
	GameStartCharlemagne         pserialize.Year `paradox_field:"game_start_charlemagne" json:"game_start_charlemagne,omitempty"`
	HadEvent1013                 pserialize.Year `paradox_field:"had_event_1013" json:"had_event_1013,omitempty"`
	RecievedRoyalMarriageAidDuty pserialize.Year `paradox_field:"recieved_royal_marriage_aid_duty" json:"recieved_royal_marriage_aid_duty,omitempty"`
}

type Player struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}

type PlayerPortrait struct {
	IsFemale   pserialize.PBool `paradox_field:"fem" json:"fem,omitempty"`
	DNA        string           `paradox_field:"dna" json:"dna,omitempty"`
	Properties string           `paradox_field:"properties" json:"properties,omitempty"`
	Religion   string           `paradox_field:"religion" json:"religion,omitempty"`
	Culture    string           `paradox_field:"culture" json:"culture,omitempty"`
	Government string           `paradox_field:"government" json:"government,omitempty"`
	Dynasty    int              `paradox_field:"dynasty" json:"dynasty,omitempty"`
}

type GameRules struct {
	AlterNateStart           pserialize.PBool `paradox_field:"alternate_start" json:"alternate_start,omitempty"`
	Epidemics                string           `paradox_field:"epidemics" json:"epidemics,omitempty"`
	DejureDriftDuration      string           `paradox_field:"dejure_drift_duration" json:"dejure_drift_duration,omitempty"`
	ConclaveEducationFocuses string           `paradox_field:"conclave_education_focuses" json:"conclave_education_focuses,omitempty"`
}
