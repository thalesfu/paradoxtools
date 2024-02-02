package save

import (
	"errors"
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SaveFile struct {
	Version         string                                `paradox_field:"version" json:"version,omitempty"`
	Date            pserialize.Year                       `paradox_field:"date" json:"date,omitempty"`
	Player          *Player                               `paradox_field:"player" json:"player,omitempty"`
	PlayerRealm     string                                `paradox_field:"player_realm" json:"player_realm,omitempty"`
	PlayerName      string                                `paradox_field:"player_name" paradox_text:"escaped" json:"player_name,omitempty"`
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
	FilePath        string                                `json:"file_name,omitempty"`
	FileHash        string                                `json:"file_hash,omitempty"`
	FileUpdateTime  time.Time                             `json:"file_update_time,omitempty"`
}

func LoadSave(path string, savePath string) (*SaveFile, bool, error) {
	var f string

	if utils.IsCompressedFile(savePath) {
		dir := filepath.Join(os.TempDir(), "ck2", "unzipsavefile", strings.TrimSuffix(filepath.Base(savePath), filepath.Ext(savePath)))

		err := utils.Unzip(savePath, dir)

		if err != nil {
			return nil, false, err
		}

		defer func() {
			err := os.RemoveAll(dir)
			if err != nil {
				fmt.Println(err)
			}
		}()

		files, err := os.ReadDir(dir)
		if err != nil {
			return nil, false, err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if strings.HasSuffix(file.Name(), ".ck2") {
				f = filepath.Join(dir, file.Name())
				break
			}
		}

		if f == "" {
			return nil, false, errors.New(fmt.Sprintf("cannot find ck2 file in unzip file %s from %s", savePath, dir))
		}
	} else {
		f = savePath
	}

	content, ok := utils.LoadContent(f)

	if !ok {
		return nil, false, errors.New("cannot load save file")
	}

	saveFile, ok := pserialize.UnmarshalP[SaveFile](content)

	if !ok {
		return nil, false, errors.New("cannot unmarshal save file")
	}

	saveFile.FilePath = strings.ReplaceAll(savePath, "\\", "/")
	hash, err := utils.GetFileHash(savePath)
	if err != nil {
		return nil, false, err
	}
	saveFile.FileHash = hash
	info, err := os.Stat(savePath)
	if err != nil {
		return nil, false, err
	}
	saveFile.FileUpdateTime = info.ModTime()

	translations := localisation.LoadAllTranslations(path)

	processTitles(saveFile, translations)
	processProvinces(saveFile, translations)
	processDynasties(saveFile, translations)
	processCharacters(saveFile, translations)

	return saveFile, true, nil
}

type IDEntity struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
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
	Weak    pserialize.PBool `paradox_field:"weak" json:"weak,omitempty"`
}

type Modifier struct {
	Modifier string           `paradox_field:"modifier" json:"modifier,omitempty"`
	Date     pserialize.Year  `paradox_field:"date" json:"date,omitempty"`
	Hidden   pserialize.PBool `paradox_field:"hidden" json:"hidden,omitempty"`
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
