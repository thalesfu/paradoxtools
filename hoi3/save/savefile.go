package save

import (
	"errors"
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type SaveFile struct {
	Player    string                     `paradox_field:"player" json:"player,omitempty"`
	Provinces map[string]*ProvinceInSave `paradox_type:"map" paradox_map_key_pattern:"\\d+"`
	Countries map[string]*CountryInSave  `paradox_type:"map" paradox_map_key_pattern:"^[a-zA-Z]{3}$"`
}

func (s *SaveFile) GetProvinceUnitCount() map[string]map[string]map[string]*UnitCount {
	allUnitCount := make(map[string]map[string]map[string]*UnitCount)

	for _, c := range s.Countries {
		cNavyUnitCount := c.GetProvincesNavyCount()
		for province, unitCount := range cNavyUnitCount {
			if _, ok := allUnitCount[province]; !ok {
				allUnitCount[province] = make(map[string]map[string]*UnitCount)
			}
			allUnitCount[province][c.ID] = unitCount
		}
		cAirUnitCount := c.GetProvincesAirCount()
		for province, unitCount := range cAirUnitCount {
			if _, ok := allUnitCount[province]; !ok {
				allUnitCount[province] = make(map[string]map[string]*UnitCount)
			}
			allUnitCount[province][c.ID] = unitCount
		}
	}

	return allUnitCount
}

func (s *SaveFile) GetProvinceStrength() map[string]map[string]*UnitStrength {
	allStrength := make(map[string]map[string]*UnitStrength)

	for _, c := range s.Countries {
		cStrength := c.GetProvincesStrength()
		for province, strength := range cStrength {
			if _, ok := allStrength[province]; !ok {
				allStrength[province] = make(map[string]*UnitStrength)
			}
			allStrength[province][c.ID] = strength
		}
	}

	return allStrength
}

func LoadSave(savePath string) (*SaveFile, bool, error) {
	var f string
	f = savePath

	content, ok := golangutils.LoadContentWithEncoding(f, simplifiedchinese.GB18030)

	if !ok {
		return nil, false, errors.New("cannot load save file")
	}

	saveFile, ok := pserialize.UnmarshalP[SaveFile](content)

	if !ok {
		return nil, false, errors.New("cannot unmarshal save file")
	}

	fmt.Println("GER")
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["ENG"].GetStrength()))
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["ENG"].GetNavyCount()))
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["ENG"].GetAirCount()))

	fmt.Println("FRA")
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["USA"].GetStrength()))
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["USA"].GetNavyCount()))
	fmt.Println(golangutils.MarshalJSON(saveFile.Countries["USA"].GetAirCount()))

	return saveFile, true, nil
}
