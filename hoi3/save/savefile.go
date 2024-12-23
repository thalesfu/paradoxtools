package save

import (
	"errors"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type SaveFile struct {
	Player   string                     `paradox_field:"player" json:"player,omitempty"`
	Province map[string]*ProvinceInSave `paradox_type:"map" paradox_map_key_pattern:"\\d+"`
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

	return saveFile, true, nil
}
