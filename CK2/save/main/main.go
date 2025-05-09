package main

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// 指定要扫描的目录
	dir := "R:\\Thales\\Game\\SteamLibrary\\steamaps\\common\\Crusader Kings II\\localisation"

	// 读取目录中的所有文件和子目录
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	regex := regexp.MustCompile(`"key":\s*"((fp)_[^"]*)"`)
	fieldsContent := ""

	// 循环遍历文件
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".csv.json") {
			filepath := filepath.Join(dir, filename)

			content, ok := golangutils.LoadContent(filepath)

			if ok {
				lines := strings.Split(content, "\n")

				for _, line := range lines {
					result := regex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]
						if strings.HasSuffix(strings.ToLower(w), "desc") {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						fieldsContent += fmt.Sprintf("%s\tpserialize.PBool `paradox_field:\"%s\" paradox_type:\"field\" json:\"%s,omitempty\"`\n", wfn, w, w)
					}
				}
			}
		}
	}

	golangutils.WriteContent("logs/fields.txt", fieldsContent)
}
