package main

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/CK2/trait"
	"github.com/thalesfu/paradoxtools/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var traitPath = "../CK2Commands/trait"

func main() {
	//buildCommandModifierField("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II/common/traits")

	all := trait.LoadAllTraits("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	for _, v := range all {
		fmt.Println(utils.MarshalJSON(v))
	}

	buildTraitFile(all)
}

func buildTraitFile(traits map[string]*trait.Trait) {
	if _, err := os.Stat(traitPath); os.IsNotExist(err) {
		err = os.Mkdir(traitPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	traitTemplate, err := template.New("TraitTemplate.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/trait/main/TraitTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(traitPath, "traits.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = traitTemplate.Execute(f, traits)
	if err != nil {
		panic(err)
	}
}

func buildCommonField(dir string) {
	// 读取目录中的所有文件和子目录
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	boolRegex := regexp.MustCompile(`^\t(\w+)\s*=\s(yes|no)`)
	boolFieldMap := make(map[string]bool)
	boolFieldsContent := ""
	float32Regex := regexp.MustCompile(`^\t(\w+)\s*=\s-?\d+\.\d*`)
	float32FieldMap := make(map[string]bool)
	float32FieldsContent := ""
	stringRegex := regexp.MustCompile(`\t(\w+)\s*=\s([a-zA-Z]\w*)`)
	stringFieldMap := make(map[string]bool)
	stringFieldsContent := ""
	intRegex := regexp.MustCompile(`^\t(\w+)\s*=\s-?\d+\.*\d*`)
	intFieldMap := make(map[string]bool)
	intFieldsContent := ""
	structRegex := regexp.MustCompile(`^\t(\w+)\s*=\s\{`)
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
			filepath := filepath.Join(dir, filename)

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
	}

	golangutils.WriteContent("logs/fields.txt", boolFieldsContent+stringFieldsContent+float32FieldsContent+intFieldsContent+structFieldsContent)
}

func buildCommandModifierField(dir string) {
	// 读取目录中的所有文件和子目录
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	boolRegex := regexp.MustCompile(`^(\w+)\s*=\s(yes|no)`)
	boolFieldMap := make(map[string]bool)
	boolFieldsContent := ""
	float32Regex := regexp.MustCompile(`(\w+)\s*=\s-?\d+\.\d*`)
	float32FieldMap := make(map[string]bool)
	float32FieldsContent := ""
	stringRegex := regexp.MustCompile(`(\w+)\s*=\s([a-zA-Z]\w*)`)
	stringFieldMap := make(map[string]bool)
	stringFieldsContent := ""
	intRegex := regexp.MustCompile(`(\w+)\s*=\s-?\d+\.*\d*`)
	intFieldMap := make(map[string]bool)
	intFieldsContent := ""
	structRegex := regexp.MustCompile(`(\w+)\s*=\s\{`)
	structFieldMap := make(map[string]bool)
	structFieldsContent := ""

	commandModifierRegex := regexp.MustCompile(`\tcommand_modifier\s*=\s*\{\s*([^\}]*)\}`)

	// 循环遍历文件
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		// 获取文件名
		filename := file.Name()

		// 检查文件后缀是否为.csv或.json
		if strings.HasSuffix(filename, ".txt") {
			filepath := filepath.Join(dir, filename)

			content, ok := golangutils.LoadContent(filepath)

			if ok {
				ms := commandModifierRegex.FindAllStringSubmatch(content, -1)

				if len(ms) > 0 {

					for _, m := range ms {

						lines := strings.Split(m[1], "\n")

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
		}
	}

	golangutils.WriteContent("logs/fields.txt", boolFieldsContent+stringFieldsContent+float32FieldsContent+intFieldsContent+structFieldsContent)
}
