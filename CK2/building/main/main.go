package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/building"
	"github.com/thalesfu/paradoxtools/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	//buildBuildingField("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")

	all := building.LoadAllBuildings("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")

	for _, v := range all {
		fmt.Println(v.Code, v.Name, v.Description)
		for _, vv := range v.Buildings {
			fmt.Print(vv.Code)
			fmt.Print("|")
		}
		fmt.Println()
	}
}

func buildBuildingField(dir string) {
	building_dir := filepath.Join(dir, "common", "buildings")
	// 读取目录中的所有文件和子目录
	files, err := os.ReadDir(building_dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	boolRegex := regexp.MustCompile(`^\t\t(\w+)\s*=\s(yes|no)`)
	boolFieldMap := make(map[string]bool)
	boolFieldsContent := ""
	float32Regex := regexp.MustCompile(`^\t\t(\w+)\s*=\s-?\d+\.\d*`)
	float32FieldMap := make(map[string]bool)
	float32FieldsContent := ""
	stringRegex := regexp.MustCompile(`^\t\t(\w+)\s*=\s([a-zA-Z]\w*)`)
	stringFieldMap := make(map[string]bool)
	stringFieldsContent := ""
	intRegex := regexp.MustCompile(`^\t\t(\w+)\s*=\s-?\d+\.*\d*`)
	intFieldMap := make(map[string]bool)
	intFieldsContent := ""
	structRegex := regexp.MustCompile(`^\t\t(\w+)\s*=\s\{`)
	structFieldMap := make(map[string]bool)
	structFieldsContent := ""

	// 循环遍历文件
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(building_dir, filename)

			content, ok := utils.LoadContent(filepath)

			if ok {
				lines := strings.Split(content, "\n")

				for _, line := range lines {
					result := boolRegex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]

						if boolFieldMap[w] {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						boolFieldsContent += fmt.Sprintf("%s\tpserialize.PBool `paradox_field:\"%s\" json:\"%s,omitempty\"`\n", wfn, w, w)
						boolFieldMap[w] = true
						continue
					}

					result = stringRegex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]

						if stringFieldMap[w] {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						stringFieldsContent += fmt.Sprintf("%s\tstring `paradox_field:\"%s\" json:\"%s,omitempty\"`\n", wfn, w, w)
						stringFieldMap[w] = true
						continue
					}

					result = float32Regex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]

						if float32FieldMap[w] {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						float32FieldsContent += fmt.Sprintf("%s\tfloat32 `paradox_field:\"%s\" json:\"%s,omitempty\"`\n", wfn, w, w)
						float32FieldMap[w] = true
						continue
					}

					result = intRegex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]

						if intFieldMap[w] {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						intFieldsContent += fmt.Sprintf("%s\tint `paradox_field:\"%s\" json:\"%s,omitempty\"`\n", wfn, w, w)
						intFieldMap[w] = true
						continue
					}

					result = structRegex.FindStringSubmatch(line)

					if len(result) > 0 {
						w := result[1]

						if structFieldMap[w] {
							continue
						}

						wp := strings.Split(w, "_")
						wfn := ""
						for _, ww := range wp {
							wfn += cases.Title(language.English).String(ww)
						}

						structFieldsContent += fmt.Sprintf("%s\t*%s `paradox_field:\"%s\" json:\"%s,omitempty\"`\n", wfn, wfn, w, w)
						structFieldMap[w] = true
						continue
					}
				}
			}
		}
	}

	utils.WriteContent("logs/building_fields.txt", boolFieldsContent+stringFieldsContent+float32FieldsContent+intFieldsContent+structFieldsContent)
}
