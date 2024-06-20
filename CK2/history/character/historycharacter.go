package character

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type HistoryCharacter struct {
	ID          int              `paradox_type:"map_key" json:"id,omitempty"`
	Female      pserialize.PBool `paradox_field:"female" json:"female,omitempty"`
	Name        string           `paradox_field:"name" json:"name,omitempty"`
	Religion    string           `paradox_field:"religion" json:"religion,omitempty"`
	Culture     string           `paradox_field:"culture" json:"culture,omitempty"`
	Trait       string           `paradox_field:"trait" json:"trait,omitempty"`
	Dna         string           `paradox_field:"dna" json:"dna,omitempty"`
	Fertility   float32          `paradox_field:"fertility" json:"fertility,omitempty"`
	Martial     int              `paradox_field:"martial" json:"martial,omitempty"`
	Diplomacy   int              `paradox_field:"diplomacy" json:"diplomacy,omitempty"`
	Intrigue    int              `paradox_field:"intrigue" json:"intrigue,omitempty"`
	Stewardship int              `paradox_field:"stewardship" json:"stewardship,omitempty"`
	Learning    int              `paradox_field:"learning" json:"learning,omitempty"`
	Health      int              `paradox_field:"health" json:"health,omitempty"`
}

func LoadAllHistoryCharacter(path string) map[int]*HistoryCharacter {

	translations := localisation.LoadAllTranslations(path)
	characterPath := filepath.Join(path, "history", "characters")
	files, err := os.ReadDir(characterPath)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[int]*HistoryCharacter)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			//if strings.HasSuffix(filename, "russian.txt") {
			filepath := filepath.Join(characterPath, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[int]*HistoryCharacter](content)

				if o {
					for k, v := range *ts {
						result[k] = v
					}
				} else {
					log.Printf("Error unmarshalling %s\n", filepath)
				}
			}
		}
	}

	for _, c := range result {
		if nt, ok := translations[fmt.Sprintf("%d#|name", c.ID)]; ok {
			c.Name = nt
		}
	}

	return result
}
