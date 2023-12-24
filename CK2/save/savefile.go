package save

import (
	"github.com/thalesfu/paradoxtools/utils/pserialize"
)

type SaveFile struct {
	Version        string                                `paradox_field:"version" paradox_type:"field" json:"version,omitempty"`
	Date           pserialize.Year                       `paradox_field:"date" paradox_type:"field" json:"date,omitempty"`
	Player         *Player                               `paradox_field:"player" paradox_type:"field" json:"player,omitempty"`
	PlayerRealm    string                                `paradox_field:"player_realm" paradox_type:"field" json:"player_realm,omitempty"`
	PlayerName     string                                `paradox_field:"player_name" paradox_type:"field" json:"player_name,omitempty"`
	PlayerAge      int                                   `paradox_field:"player_age" paradox_type:"field" json:"player_age,omitempty"`
	PlayerPortrait *PlayerPortrait                       `paradox_field:"player_portrait" paradox_type:"field" json:"player_portrait,omitempty"`
	IsZeusSave     pserialize.PBool                      `paradox_field:"is_zeus_save" paradox_type:"field" json:"is_zeus_save,omitempty"`
	GameRules      *GameRules                            `paradox_field:"game_rules" paradox_type:"field" json:"game_rules,omitempty"`
	GameSpeed      int                                   `paradox_field:"game_speed" paradox_type:"field" json:"game_speed,omitempty"`
	MapMode        int                                   `paradox_field:"mapmode" paradox_type:"field" json:"mapmode,omitempty"`
	PlayThroughID  int                                   `paradox_field:"playthrough_id" paradox_type:"field" json:"playthrough_id,omitempty"`
	DynastyTitle   []*DynastyTitle                       `paradox_field:"dyn_title" paradox_type:"list" json:"dyn_title,omitempty"`
	Unit           int                                   `paradox_field:"unit" paradox_type:"field" json:"unit,omitempty"`
	SubUnit        int                                   `paradox_field:"sub_unit" paradox_type:"field" json:"sub_unit,omitempty"`
	StartDate      pserialize.Year                       `paradox_field:"start_date" paradox_type:"field" json:"start_date,omitempty"`
	Flags          *Flag                                 `paradox_field:"flags" paradox_type:"field" json:"flags,omitempty"`
	Vars           *Vars                                 `paradox_field:"vars" paradox_type:"field" json:"vars,omitempty"`
	Dynasties      map[int]*Dynasty                      `paradox_field:"dynasties" paradox_type:"field" json:"dynasties,omitempty"`
	Titles         map[string]*Title                     `paradox_field:"title" paradox_type:"field" json:"title,omitempty"`
	Characters     map[int]*Character                    `paradox_field:"character" paradox_type:"field" json:"character,omitempty"`
	DelayedEvents  []*DelayedEvent                       `paradox_field:"delayed_event" paradox_type:"list" json:"delayed_event,omitempty"`
	Relations      map[string]map[string]*PersonRelation `paradox_field:"relation" paradox_type:"field" json:"relation,omitempty"`
	ID             *IDEntity                             `paradox_field:"id" paradox_type:"field" json:"id,omitempty"`
	Religions      map[string]*Religion                  `paradox_field:"religion" paradox_type:"field" json:"religion,omitempty"`
}

type IDEntity struct {
	ID   int `paradox_field:"id" paradox_type:"field" json:"id,omitempty"`
	Type int `paradox_field:"type" paradox_type:"field" json:"type,omitempty"`
}

type DelayedEventScope struct {
	SavedEventTargets []*SavedEventTarget `paradox_field:"saved_event_target" paradox_type:"list" json:"saved_event_target,omitempty"`
	Seed              int                 `paradox_field:"seed" paradox_type:"field" json:"seed,omitempty"`
	Char              int                 `paradox_field:"char" paradox_type:"field" json:"char,omitempty"`
	From              *DelayedEventScope  `paradox_field:"from" paradox_type:"field" json:"from,omitempty"`
	Root              *DelayedEventScope  `paradox_field:"root" paradox_type:"field" json:"root,omitempty"`
}

type DelayedEvent struct {
	Event string             `paradox_field:"event" paradox_type:"field" json:"event,omitempty"`
	Days  int                `paradox_field:"days" paradox_type:"field" json:"days,omitempty"`
	Scope *DelayedEventScope `paradox_field:"scope" paradox_type:"field" json:"scope,omitempty"`
}

type Title struct {
	ID string `paradox_type:"map_key" json:"id,omitempty"`
}

type Character struct {
	ID                  int                  `paradox_type:"map_key" json:"id,omitempty"`
	Name                string               `paradox_field:"bn" paradox_type:"field" json:"bn,omitempty"`
	Hist                pserialize.PBool     `paradox_field:"hist" paradox_type:"field" json:"hist,omitempty"`
	Pi                  pserialize.PBool     `paradox_field:"pi" paradox_type:"field" json:"pi,omitempty"`
	Birthday            pserialize.Year      `paradox_field:"b_d" paradox_type:"field" json:"b_d,omitempty"`
	DeathDay            pserialize.Year      `paradox_field:"d_d" paradox_type:"field" json:"d_d,omitempty"`
	Father              int                  `paradox_field:"fat" paradox_type:"field" json:"fat,omitempty"`
	Spouse              int                  `paradox_field:"spouse" paradox_type:"field" json:"spouse,omitempty"`
	Attributes          []int                `paradox_field:"att" paradox_type:"field_list" json:"att,omitempty"`
	Traits              []int                `paradox_field:"tr" paradox_type:"field_list" json:"tr,omitempty"`
	Dynasty             int                  `paradox_field:"dnt" paradox_type:"field" json:"dnt,omitempty"`
	DNA                 string               `paradox_field:"dna" paradox_type:"field" json:"dna,omitempty"`
	Government          string               `paradox_field:"gov" paradox_type:"field" json:"gov,omitempty"`
	Fertility           float32              `paradox_field:"fer" paradox_type:"field" json:"fer,omitempty"`
	Health              float32              `paradox_field:"health" paradox_type:"field" json:"health,omitempty"`
	Consort             []int                `paradox_field:"consort" paradox_type:"list" json:"consort,omitempty"`
	Prs                 float32              `paradox_field:"prs" paradox_type:"field" json:"prs,omitempty"`
	Piety               float32              `paradox_field:"piety" paradox_type:"field" json:"piety,omitempty"`
	Wealth              float32              `paradox_field:"wealth" paradox_type:"field" json:"wealth,omitempty"`
	Emp                 int                  `paradox_field:"emp" paradox_type:"field" json:"emp,omitempty"`
	Host                int                  `paradox_field:"host" paradox_type:"field" json:"host,omitempty"`
	Curinc              float32              `paradox_field:"curinc" paradox_type:"field" json:"curinc,omitempty"`
	Emi                 float32              `paradox_field:"emi" paradox_type:"field" json:"emi,omitempty"`
	Eme                 float32              `paradox_field:"eme" paradox_type:"field" json:"eme,omitempty"`
	Eyi                 float32              `paradox_field:"eyi" paradox_type:"field" json:"eyi,omitempty"`
	Eypi                float32              `paradox_field:"eypi" paradox_type:"field" json:"eypi,omitempty"`
	Action              string               `paradox_field:"action" paradox_type:"field" json:"action,omitempty"`
	ACDate              pserialize.Year      `paradox_field:"acdate" paradox_type:"field" json:"acdate,omitempty"`
	ACLoc               int                  `paradox_field:"acloc" paradox_type:"field" json:"acloc,omitempty"`
	AmbitionDate        pserialize.Year      `paradox_field:"ambition_date" paradox_type:"field" json:"ambition_date,omitempty"`
	FocusDate           pserialize.Year      `paradox_field:"focus_date" paradox_type:"field" json:"focus_date,omitempty"`
	CPos                string               `paradox_field:"cpos" paradox_type:"field" json:"cpos,omitempty"`
	CPosStart           int                  `paradox_field:"cpos_start" paradox_type:"field" json:"cpos_start,omitempty"`
	Vars                *Vars                `paradox_field:"vars" paradox_type:"field" json:"vars,omitempty"`
	Flags               *Flag                `paradox_field:"flags" paradox_type:"field" json:"flags,omitempty"`
	Modifier            *Modifier            `paradox_field:"md" paradox_type:"field" json:"md,omitempty"`
	Claim               *Claim               `paradox_field:"claim" paradox_type:"field" json:"claim,omitempty"`
	Lgr                 *Lgr                 `paradox_field:"lgr" paradox_type:"field" json:"lgr,omitempty"`
	SavedEventTarget    *SavedEventTarget    `paradox_field:"saved_event_target" paradox_type:"field" json:"saved_event_target,omitempty"`
	OffMapCurrencies    map[int]float32      `paradox_field:"offmap_currencies" paradox_type:"field" json:"offmap_currencies,omitempty"`
	DMN                 *DMN                 `paradox_field:"dmn" paradox_type:"field" json:"dmn,omitempty"`
	CharacterPlayerData *CharacterPlayerData `paradox_field:"character_player_data" paradox_type:"field" json:"character_player_data,omitempty"`
}

type ChronicleEntry struct {
	Text            string `paradox_field:"text" paradox_type:"field" json:"text,omitempty"`
	Picture         string `paradox_field:"picture" paradox_type:"field" json:"picture,omitempty"`
	Portrait        int    `paradox_field:"portrait" paradox_type:"field" json:"portrait,omitempty"`
	PortraitCulture string `paradox_field:"portrait_culture" paradox_type:"field" json:"portrait_culture,omitempty"`
	Type            string `paradox_field:"type" paradox_type:"field" json:"type,omitempty"`
}

type ChronicleChapter struct {
	Chronicles []*ChronicleEntry `paradox_field:"chronicle_entry" paradox_type:"list" json:"chronicle_entry,omitempty"`
	Year       int               `paradox_field:"year" paradox_type:"field" json:"year,omitempty"`
}

type Chronicle struct {
	ChronicleChapters []*ChronicleChapter `paradox_field:"chronicle_chapter" paradox_type:"list" json:"chronicle_chapter,omitempty"`
	Character         int                 `paradox_field:"character" paradox_type:"field" json:"character,omitempty"`
}

type ChronicleCollection struct {
	Chronicle         *Chronicle       `paradox_field:"chronicle" paradox_type:"field" json:"chronicle,omitempty"`
	ChroniclePosition int              `paradox_field:"chronicle_position" paradox_type:"field" json:"chronicle_position,omitempty"`
	ChapterPosition   int              `paradox_field:"chapter_position" paradox_type:"field" json:"chapter_position,omitempty"`
	EntryPosition     int              `paradox_field:"entry_position" paradox_type:"field" json:"entry_position,omitempty"`
	ChronicleIconLit  pserialize.PBool `paradox_field:"chronicle_icon_lit" paradox_type:"field" json:"chronicle_icon_lit,omitempty"`
}

type CharacterPlayerData struct {
	ChronicleCollection         *ChronicleCollection `paradox_field:"chronicle_collection" paradox_type:"field" json:"chronicle_collection,omitempty"`
	SocietyShowInterestCooldown pserialize.Year      `paradox_field:"society_show_interest_cooldown" paradox_type:"field" json:"society_show_interest_cooldown,omitempty"`
	JoinSocietyCooldown         pserialize.Year      `paradox_field:"join_society_cooldown" paradox_type:"field" json:"join_society_cooldown,omitempty"`
	Telws                       int                  `paradox_field:"telws" paradox_type:"field" json:"telws,omitempty"`
	Telbc                       int                  `paradox_field:"telbc" paradox_type:"field" json:"telbc,omitempty"`
	Telde                       int                  `paradox_field:"telde" paradox_type:"field" json:"telde,omitempty"`
	Telld                       pserialize.PBool     `paradox_field:"telld" paradox_type:"field" json:"telld,omitempty"`
	Telsc                       pserialize.PBool     `paradox_field:"telsc" paradox_type:"field" json:"telsc,omitempty"`
}

type LiegeTroops struct {
	Li []float32 `paradox_field:"li" paradox_type:"field_list" json:"li,omitempty"`
	Hi []float32 `paradox_field:"hi" paradox_type:"field_list" json:"hi,omitempty"`
	Pi []float32 `paradox_field:"pi" paradox_type:"field_list" json:"pi,omitempty"`
	Lc []float32 `paradox_field:"lc" paradox_type:"field_list" json:"lc,omitempty"`
	Ar []float32 `paradox_field:"ar" paradox_type:"field_list" json:"ar,omitempty"`
}

type DMN struct {
	Capital                 string       `paradox_field:"capital" paradox_type:"field" json:"capital,omitempty"`
	Primary                 *Title       `paradox_field:"primary" paradox_type:"map_value" paradox_map_name:"title" json:"primary,omitempty"`
	LiegeTroops             *LiegeTroops `paradox_field:"liege_troops" paradox_type:"field" json:"liege_troops,omitempty"`
	RaisedLiegeTroops       []int        `paradox_field:"raised_liege_troops" paradox_type:"field_list" json:"raised_liege_troops,omitempty"`
	MyLiegelevyContribution int          `paradox_field:"my_liegelevy_contribution" paradox_type:"field" json:"my_liegelevy_contribution,omitempty"`
	PeaceMonths             int          `paradox_field:"peace_months" paradox_type:"field" json:"peace_months,omitempty"`
}

type SavedEventTarget struct {
	Name string `paradox_field:"name" paradox_type:"field" json:"name,omitempty"`
	Char int    `paradox_field:"char" paradox_type:"field" json:"char,omitempty"`
}

type Lgr struct {
	LastMonthIncomeTable  []float32 `paradox_field:"lastmonthincometable" paradox_type:"field_list" json:"lastmonthincometable,omitempty"`
	LastMonthExpenseTable []float32 `paradox_field:"lastmonthexpensetable" paradox_type:"field_list" json:"lastmonthexpensetable,omitempty"`
	LastMonthIncome       float32   `paradox_field:"lastmonthincome" paradox_type:"field" json:"lastmonthincome,omitempty"`
	LastMonthExpense      float32   `paradox_field:"lastmonthexpense" paradox_type:"field" json:"lastmonthexpense,omitempty"`
}

type Claim struct {
	Title   *Title           `paradox_field:"title" paradox_type:"map_value" paradox_map_name:"title" json:"title,omitempty"`
	Pressed pserialize.PBool `paradox_field:"pressed" paradox_type:"field" json:"pressed,omitempty"`
}

type Modifier struct {
	Modifier string           `paradox_field:"modifier" paradox_type:"field" json:"modifier,omitempty"`
	Date     pserialize.Year  `paradox_field:"date" paradox_type:"field" json:"date,omitempty"`
	Hidden   pserialize.PBool `paradox_field:"hidden" paradox_type:"field" json:"hidden,omitempty"`
}

type CoatOfArms struct {
	Data     []int  `paradox_field:"data" paradox_type:"field_list" json:"data,omitempty"`
	Religion string `paradox_field:"religion" paradox_type:"field" json:"religion,omitempty"`
}

type Dynasty struct {
	ID            int              `paradox_type:"map_key" json:"id,omitempty"`
	Name          string           `paradox_field:"name" paradox_type:"field" json:"name,omitempty"`
	Culture       string           `paradox_field:"culture" paradox_type:"field" json:"culture,omitempty"`
	Religion      string           `paradox_field:"religion" paradox_type:"field" json:"religion,omitempty"`
	CoatOfArms    *CoatOfArms      `paradox_field:"coat_of_arms" paradox_type:"field" json:"coat_of_arms,omitempty"`
	SetCoatOfArms pserialize.PBool `paradox_field:"set_coat_of_arms" paradox_type:"field" json:"set_coat_of_arms,omitempty"`
}

type Vars struct {
	GlobalAmountOfTimesChinaFamine             float32 `paradox_field:"global_amount_of_times_china_famine" paradox_type:"field" json:"global_amount_of_times_china_famine,omitempty"`
	GlobalDeathsInBattle                       float32 `paradox_field:"global_deaths_in_battle" paradox_type:"field" json:"global_deaths_in_battle,omitempty"`
	GlobalRaidingAdventurerSpawnByDisplacement float32 `paradox_field:"global_raiding_adventurer_spawn_by_displacement" paradox_type:"field" json:"global_raiding_adventurer_spawn_by_displacement,omitempty"`
	GlobalTempPolicyYears                      float32 `paradox_field:"global_temp_policy_years" paradox_type:"field" json:"global_temp_policy_years,omitempty"`
	PhysiqueVariable                           float32 `paradox_field:"physique_variable" paradox_type:"field" json:"physique_variable,omitempty"`
}

type Flag struct {
	AchievementGetMarried        pserialize.Year `paradox_field:"achievement_get_married" paradox_type:"field" json:"achievement_get_married,omitempty"`
	AvarKhaganateRenamed         pserialize.Year `paradox_field:"avar_khaganate_renamed" paradox_type:"field" json:"avar_khaganate_renamed,omitempty"`
	BerniciaRenamed              pserialize.Year `paradox_field:"bernicia_renamed" paradox_type:"field" json:"bernicia_renamed,omitempty"`
	CGStudyMilitary              pserialize.Year `paradox_field:"cg_study_military" paradox_type:"field" json:"cg_study_military,omitempty"`
	FlagCrownsDelivered          pserialize.Year `paradox_field:"flag_crowns_delivered" paradox_type:"field" json:"flag_crowns_delivered,omitempty"`
	GameStartCharlemagne         pserialize.Year `paradox_field:"game_start_charlemagne" paradox_type:"field" json:"game_start_charlemagne,omitempty"`
	HadEvent1013                 pserialize.Year `paradox_field:"had_event_1013" paradox_type:"field" json:"had_event_1013,omitempty"`
	RecievedRoyalMarriageAidDuty pserialize.Year `paradox_field:"recieved_royal_marriage_aid_duty" paradox_type:"field" json:"recieved_royal_marriage_aid_duty,omitempty"`
}

type DynastyTitle struct {
	Title     string           `paradox_field:"title" paradox_type:"field" json:"title,omitempty"`
	BaseTitle string           `paradox_field:"base_title" paradox_type:"field" json:"base_title,omitempty"`
	IsCustom  pserialize.PBool `paradox_field:"is_custom" paradox_type:"field" json:"is_custom,omitempty"`
	IsDynamic pserialize.PBool `paradox_field:"is_dynamic" paradox_type:"field" json:"is_dynamic,omitempty"`
}

type Player struct {
	ID   int `paradox_field:"id" paradox_type:"field" json:"id,omitempty"`
	Type int `paradox_field:"type" paradox_type:"field" json:"type,omitempty"`
}

type PlayerPortrait struct {
	IsFemale   pserialize.PBool `paradox_field:"fem" paradox_type:"field" json:"fem,omitempty"`
	DNA        string           `paradox_field:"dna" paradox_type:"field" json:"dna,omitempty"`
	Properties string           `paradox_field:"properties" paradox_type:"field" json:"properties,omitempty"`
	Religion   string           `paradox_field:"religion" paradox_type:"field" json:"religion,omitempty"`
	Culture    string           `paradox_field:"culture" paradox_type:"field" json:"culture,omitempty"`
	Government string           `paradox_field:"government" paradox_type:"field" json:"government,omitempty"`
	Dynasty    int              `paradox_field:"dynasty" paradox_type:"field" json:"dynasty,omitempty"`
}

type GameRules struct {
	AlterNateStart           pserialize.PBool `paradox_field:"alternate_start" paradox_type:"field" json:"alternate_start,omitempty"`
	Epidemics                string           `paradox_field:"epidemics" paradox_type:"field" json:"epidemics,omitempty"`
	DejureDriftDuration      string           `paradox_field:"dejure_drift_duration" paradox_type:"field" json:"dejure_drift_duration,omitempty"`
	ConclaveEducationFocuses string           `paradox_field:"conclave_education_focuses" paradox_type:"field" json:"conclave_education_focuses,omitempty"`
}
