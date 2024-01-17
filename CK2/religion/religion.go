package religion

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type ReligionGroup struct {
	Code      string               `paradox_type:"map_key" json:"code,omitempty"`
	Name      string               `json:"name,omitempty"`
	Religions map[string]*Religion `paradox_type:"map" paradox_map_key_pattern:"(aztec|aztec_reformed|baltic_pagan|baltic_pagan_reformed|bogomilist|bon|bon_reformed|buddhist|cathar|catholic|druze|finnish_pagan|finnish_pagan_reformed|fraticelli|hellenic_pagan|hellenic_pagan_reformed|hindu|hurufi|ibadi|iconoclast|jain|jewish|karaite|kharijite|khurmazta|lollard|manichean|mazdaki|messalian|miaphysite|monophysite|monothelite|nestorian|norse_pagan|norse_pagan_reformed|orthodox|paulician|qarmatian|samaritan|shiite|slavic_pagan|slavic_pagan_reformed|sunni|taoist|waldensian|west_african_pagan|west_african_pagan_reformed|yazidi|zikri|zoroastrian|zun_pagan|zun_pagan_reformed)" json:"religions,omitempty"`
}

type Religion struct {
	Code string `paradox_type:"map_key" json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func LoadAllReligions(path string) map[string]*ReligionGroup {

	translations := localisation.LoadAllTranslations(path)
	religionPath := filepath.Join(path, "common", "religions")
	files, err := os.ReadDir(religionPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*ReligionGroup)

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
				ts, o := pserialize.UnmarshalP[map[string]*ReligionGroup](content)

				if o {
					for k, v := range *ts {
						result[k] = v
					}
				}
			}
		}
	}

	for _, rg := range result {
		rg.Name = translations[rg.Code]
		for _, r := range rg.Religions {
			r.Name = translations[r.Code]
		}
	}

	return result
}
