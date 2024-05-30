package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/utils"
	"html/template"
	"os"
	"path"
	"strings"
)

var religionPath = "../CK2Commands/religion"

func main() {
	religionGroups := religion.LoadAllReligions("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(religionGroups))

	buildReligionFile(religionGroups)
}

func buildReligionFile(religionGroups map[string]*religion.ReligionGroup) {
	if _, err := os.Stat(religionPath); os.IsNotExist(err) {
		err = os.Mkdir(religionPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	religionTemplate, err := template.New("ReligionTemplate.txt").Funcs(template.FuncMap{"RP": replaceSpecialWords, "ES": escapeSpecialWords}).ParseFiles("Ck2/religion/main/ReligionTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(religionPath, "religion.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = religionTemplate.Execute(f, religionGroups)
	if err != nil {
		panic(err)
	}
}

func replaceSpecialWords(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "－", "_")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "’", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, "“", "")
	s = strings.ReplaceAll(s, "”", "")
	s = strings.ReplaceAll(s, "‘", "")
	s = strings.ReplaceAll(s, "’", "")
	return s
}

func escapeSpecialWords(s string) string {
	if s == "" {
		return s
	}

	text, err := utils.DecodeEscapedText([]byte(s))
	if err != nil {
		fmt.Println(err)
		return s
	}
	return text
}
