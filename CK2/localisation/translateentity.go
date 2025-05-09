package localisation

import (
	"github.com/thalesfu/golangutils"
	"os"
	"path/filepath"
	"strings"
)

type TranslateEntity struct {
	Key         string ` json:"key,omitempty"`
	Original    string ` json:"original,omitempty"`
	Translation string ` json:"translation,omitempty"`
	Stage       int    ` json:"stage,omitempty"`
	File        string `json:"file,omitempty"`
}

func LoadAllTranslations(path string) map[string]string {
	result := make(map[string]string)

	translations, _ := LoadAllTranslationsDetail(path)

	for _, t := range translations {
		result[t.Key] = t.Translation
	}

	return result
}

func LoadAllTranslationsDetail(path string) (map[string]*TranslateEntity, map[string][]*TranslateEntity) {
	result := make(map[string]*TranslateEntity)
	repeated := make(map[string][]*TranslateEntity)

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// 获取文件名
		filename := info.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".csv.json") || strings.HasSuffix(filename, ".txt.json") {
			content, ok := golangutils.LoadContent(path)

			if ok {
				ts, o := golangutils.UnmarshalJSON[[]*TranslateEntity](content)

				if o {
					for _, t := range *ts {
						t.File = path
						if existed, ok := result[t.Key]; ok {
							if existed.Translation != t.Translation {
								if existedRepeated, ok := repeated[t.Key]; ok {
									repeated[t.Key] = append(existedRepeated, t)
								} else {
									repeated[t.Key] = []*TranslateEntity{existed, t}
								}
							}
						}

						result[t.Key] = t
					}
				}
			}
		}

		return nil
	})

	return result, repeated
}
