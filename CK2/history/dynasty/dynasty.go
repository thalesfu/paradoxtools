package dynasty

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type HistoryDynasty struct {
	ID       int    `paradox_type:"map_key" json:"id,omitempty"`
	Religion string `paradox_field:"religion" json:"religion,omitempty"`
	Culture  string `paradox_field:"culture" json:"culture,omitempty"`
}

func LoadAllHistoryDynasty(path string) map[int]*HistoryDynasty {
	characterPath := filepath.Join(path, "common", "dynasties")
	files, err := os.ReadDir(characterPath)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[int]*HistoryDynasty)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(characterPath, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := pserialize.UnmarshalP[map[int]*HistoryDynasty](content)

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

	return result
}
