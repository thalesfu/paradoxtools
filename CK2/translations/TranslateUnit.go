package translations

import (
	"encoding/json"
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/ck2utils"
	"github.com/thalesfu/paradoxtools/CK2/config"
	"github.com/thalesfu/paradoxtools/utils"
	"os"
	"path/filepath"
	"strings"
)

type Unit struct {
	Key         string `json:"key"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
	Stage       int    `json:"stage"`
}

var AllTranslations = map[string]map[string]Unit{}
var FeudTranslations = map[string]Unit{}
var FeudAdjTranslations = map[string]Unit{}

func init() {
	err := filepath.Walk(config.CKIIPath, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if strings.HasSuffix(p, ".csv.json") || strings.HasSuffix(p, ".txt.json") {
			bytes, err := os.ReadFile(p)
			if err != nil {
				fmt.Println("error:", err)
				return err
			}

			var units []Unit
			err = json.Unmarshal(bytes, &units)
			if err != nil {
				fmt.Println("error:", err)
				return err
			}

			subMap := map[string]Unit{}

			for _, unit := range units {
				subMap[unit.Key] = unit

				if ck2utils.IsFeudString(unit.Key) {
					FeudTranslations[unit.Key] = unit
				}

				if ck2utils.IsFeudAdjString(unit.Key) {
					FeudAdjTranslations[unit.Key] = unit
				}
			}

			AllTranslations[fi.Name()] = subMap
		}
		return nil
	})
	if err != nil {
		fmt.Println("error:", err)
	}
}

func GetFeudName(name string) string {
	if u, ok := FeudTranslations[name]; ok {
		return utils.ReplaceSpecialChars(u.Translation)
	}

	if u, ok := FeudAdjTranslations[name+"_adj"]; ok {
		return utils.ReplaceSpecialChars(u.Translation)
	}

	fmt.Println("No translation for", name)
	return name
}
