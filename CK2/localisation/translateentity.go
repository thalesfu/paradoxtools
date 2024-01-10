package localisation

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
	"os"
	"path/filepath"
	"strings"
)

type TranslateEntity struct {
	Key         string ` json:"key,omitempty"`
	Original    string ` json:"original,omitempty"`
	Translation string ` json:"translation,omitempty"`
	Stage       int    ` json:"stage,omitempty"`
}

func LoadAllTranslations(path string) map[string]string {
	// 读取目录中的所有文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	result := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".csv.json") {
			filepath := filepath.Join(path, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				ts, o := utils.UnmarshalJSON[[]*TranslateEntity](content)

				if o {
					for _, t := range *ts {
						result[t.Key] = t.Translation
					}
				}
			}
		}
	}

	return result
}
