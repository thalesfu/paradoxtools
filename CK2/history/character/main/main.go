package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools"
	"github.com/thalesfu/paradoxtools/CK2/history/character"
	"github.com/thalesfu/paradoxtools/utils"
	"html/template"
	"os"
	"path"
)

var historyPeoplePath = "../CK2Commands/historypeople"

func main() {
	//utils.BuildFields(filepath.Join(paradoxtools.Ck2path, "history", "characters"), 1)

	historyCharacter := character.LoadAllHistoryCharacter(paradoxtools.Ck2path)

	//fmt.Println(utils.MarshalJSON(all))

	buildHistoryPeopleFile(historyCharacter)
}

func buildHistoryPeopleFile(historyCharacter map[int]*character.HistoryCharacter) {
	if _, err := os.Stat(historyPeoplePath); os.IsNotExist(err) {
		err = os.Mkdir(historyPeoplePath, 0755)
		if err != nil {
			panic(err)
		}
	}

	characters := make(map[int]map[int]*character.HistoryCharacter)

	for k, v := range historyCharacter {
		key := k % 10000
		if _, ok := characters[key]; !ok {
			characters[key] = make(map[int]*character.HistoryCharacter)
		}

		characters[key][k] = v
	}

	peopleTemplate, err := template.New("HistoryPeopleTemplate.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/history/character/main/HistoryPeopleTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(historyPeoplePath, "people.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = peopleTemplate.Execute(f, characters)
	if err != nil {
		panic(err)
	}

	peopleTemplate2, err := template.New("HistoryPeopleTemplate2.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/history/character/main/HistoryPeopleTemplate2.txt")
	if err != nil {
		panic(err)
	}

	for k, v := range characters {
		f, err := os.OpenFile(path.Join(historyPeoplePath, fmt.Sprintf("people_%d.go", k)), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		subCharacter := make(map[int]map[int]*character.HistoryCharacter)
		subCharacter[k] = v

		err = peopleTemplate2.Execute(f, subCharacter)
		if err != nil {
			panic(err)
		}
	}
}
