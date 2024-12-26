package save

import (
	"errors"
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type SaveFile struct {
	Player    string                     `paradox_field:"player" json:"player,omitempty"`
	Provinces map[string]*ProvinceInSave `paradox_type:"map" paradox_map_key_pattern:"\\d+"`
	Countries map[string]*CountryInSave  `paradox_type:"map" paradox_map_key_pattern:"^[a-zA-Z]{3}$"`
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

	content, ok := utils.LoadContentWithEncoding(f, simplifiedchinese.GB18030)

	if !ok {
		return nil, false, errors.New("cannot load save file")
	}

	saveFile, ok := pserialize.UnmarshalP[SaveFile](content)

	if !ok {
		return nil, false, errors.New("cannot unmarshal save file")
	}

	fmt.Println(utils.MarshalJSON(saveFile.Countries["JAP"].GetStrength()))
	fmt.Println(utils.MarshalJSON(saveFile.Countries["CHI"].GetStrength()))

	return saveFile, true, nil
}
