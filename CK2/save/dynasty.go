package save

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"time"
)

type Dynasty struct {
	ID            int              `paradox_type:"map_key" json:"id,omitempty"`
	Name          string           `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	Culture       string           `paradox_field:"culture" json:"culture,omitempty"`
	Religion      string           `paradox_field:"religion" json:"religion,omitempty"`
	CoatOfArms    *CoatOfArms      `paradox_field:"coat_of_arms" json:"coat_of_arms,omitempty"`
	SetCoatOfArms pserialize.PBool `paradox_field:"set_coat_of_arms" json:"set_coat_of_arms,omitempty"`
	PlayID        int              `description:"game play id" json:"play_id,omitempty"`
	PlayDate      time.Time        `description:"game play date" json:"play_date,omitempty"`
}

func processDynasties(saveFile *SaveFile, translations map[string]string) {
	for _, d := range saveFile.Dynasties {
		d.PlayID = saveFile.PlayThroughID
		d.PlayDate = time.Time(saveFile.Date)

		if d.Name == "" {
			k := fmt.Sprintf("%d#|name", d.ID)
			d.Name = translations[k]
		}
	}
}
