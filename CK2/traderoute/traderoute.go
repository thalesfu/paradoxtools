package traderoute

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
	"github.com/thalesfu/paradoxtools/CK2/modifier"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"os"
	"path/filepath"
	"strings"
)

type TradeRoute struct {
	Code     string             `paradox_type:"map_key" json:"code,omitempty"`
	Modifier *modifier.Modifier `paradox_field:"modifier" json:"modifier,omitempty"`
}

func LoadAllTradeRoutes(path string) map[string]*TradeRoute {
	translations := localisation.LoadAllTranslations(path)
	tradeRoutePath := filepath.Join(path, "common", "trade_routes")
	files, err := os.ReadDir(tradeRoutePath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]*TradeRoute)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(tradeRoutePath, filename)

			content, ok := golangutils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[string]*TradeRoute](content)

				if o {
					for k, v := range *ts {
						result[k] = v
					}
				}
			}
		}
	}

	for _, m := range result {
		m.Modifier.Code = m.Code
		m.Modifier.Name = translations[m.Code]
		m.Modifier.Description = translations[m.Code+"_desc"]
	}

	return result
}
