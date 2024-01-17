package culture

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type CultureGroup struct {
	Code     string              `paradox_type:"map_key" json:"code,omitempty"`
	Name     string              `json:"name,omitempty"`
	Cultures map[string]*Culture `paradox_type:"map" paradox_map_key_pattern:"(norse|swedish|norwegian|danish|german|lombard|old_frankish|suebi|english|saxon|old_saxon|frisian|dutch|frankish|norman|italian|occitan|roman|dalmatian|sardinian|basque|castillan|catalan|portuguese|visigothic|arberian|armenian|greek|alan|georgian|assyrian|crimean_gothic|irish|scottish|pictish|welsh|breton|finnish|lappish|ugricbaltic|komi|khanty|mordvin|meshchera|lettigallish|lithuanian|prussian|karluk|kirghiz|uyghur|mongol|khitan|jurchen|levantine_arabic|egyptian_arabic|andalusian_arabic|russian|ilmenian|severian|volhynian|pommeranian|bohemian|polish|slovieni|croatian|serbian|romanian|bulgarian|bosnian|carantanian|hungarian|persian|sogdian|tocharian|kurdish|saka|ethiopian|somali|nubian|daju|kanuri|hausa|zaghawa|manden|soninke|songhay|nahuatl|ashkenazi|sephardi|nepali|tangut|han|horse|cat|bear|hedgehog_culture|duck_culture|dog_culture|elephant_culture|dragon_culture|red_panda)" json:"cultures,omitempty"`
}

type Culture struct {
	Code                  string           `paradox_type:"map_key" json:"code,omitempty"`
	Name                  string           `json:"name,omitempty"`
	FromDynastyPrefix     string           `paradox_field:"from_dynasty_prefix" json:"from_dynasty_prefix,omitempty"`
	MalePatronym          string           `paradox_field:"male_patronym" json:"male_patronym,omitempty"`
	FemalePatronym        string           `paradox_field:"female_patronym" json:"female_patronym,omitempty"`
	DynastyTitleNames     pserialize.PBool `paradox_field:"dynasty_title_names" json:"dynasty_title_names,omitempty"`
	FounderNamedDynasties pserialize.PBool `paradox_field:"founder_named_dynasties" json:"founder_named_dynasties,omitempty"`
	Castes                pserialize.PBool `paradox_field:"castes" json:"castes,omitempty"`
	DynastyNameFirst      pserialize.PBool `paradox_field:"dynasty_name_first" json:"dynasty_name_first,omitempty"`
	DukesCalledKings      pserialize.PBool `paradox_field:"dukes_called_kings" json:"dukes_called_kings,omitempty"`
	CountTitlesHidden     pserialize.PBool `paradox_field:"count_titles_hidden" json:"count_titles_hidden,omitempty"`
	BaronTitlesHidden     pserialize.PBool `paradox_field:"baron_titles_hidden" json:"baron_titles_hidden,omitempty"`
	AllowLooting          pserialize.PBool `paradox_field:"allow_looting" json:"allow_looting,omitempty"`
}

func LoadAllCultures(path string) map[string]*CultureGroup {

	translations := localisation.LoadAllTranslations(path)
	culturePath := filepath.Join(path, "common", "cultures")
	files, err := os.ReadDir(culturePath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*CultureGroup)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名E
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(culturePath, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*CultureGroup](content)

				if o {
					for k, v := range *ts {
						result[k] = v
					}
				}
			}
		}
	}

	for _, cg := range result {
		cg.Name = translations[cg.Code]
		for _, c := range cg.Cultures {
			c.Name = translations[c.Code]
		}
	}

	return result
}
