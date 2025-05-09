package utils

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func BuildFields(path string, indent int) {

	files, err := os.ReadDir(path)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	boolRegex := regexp.MustCompile(fmt.Sprintf(`^\t{%d}(\w+)\s*=\s*(yes|no)`, indent))
	float32Regex := regexp.MustCompile(fmt.Sprintf(`^\t{%d}(\w+)\s*=\s*-?\d+\.\d*`, indent))
	stringRegex := regexp.MustCompile(fmt.Sprintf(`^\t{%d}(\w+)\s*=\s*([a-zA-Z]\w*)`, indent))
	intRegex := regexp.MustCompile(fmt.Sprintf(`^\t{%d}(\w+)\s*=\s*-?\d+\.*\d*`, indent))
	structRegex := regexp.MustCompile(fmt.Sprintf(`^\t{%d}(\w+)\s*=\s*\{`, indent))
	boolFieldMap := make(map[string]bool)
	boolFieldsContent := ""
	float32FieldMap := make(map[string]bool)
	float32FieldsContent := ""
	stringFieldMap := make(map[string]bool)
	stringFieldsContent := ""
	intFieldMap := make(map[string]bool)
	intFieldsContent := ""
	structFieldMap := make(map[string]bool)
	structFieldsContent := ""

	// 循环遍历文件
	for _, file := range files {
		if file.IsDir() || strings.HasSuffix(file.Name(), ".json") {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		filepath := filepath.Join(path, filename)

		content, ok := golangutils.LoadContent(filepath)

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

	golangutils.WriteContent("temps/fields.txt", boolFieldsContent+stringFieldsContent+float32FieldsContent+intFieldsContent+structFieldsContent)
}
