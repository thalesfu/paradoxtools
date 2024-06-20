package main

import (
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

	religionTemplate, err := template.New("HistoryPeopleTemplate.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/history/character/main/HistoryPeopleTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(historyPeoplePath, "people.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = religionTemplate.Execute(f, historyCharacter)
	if err != nil {
		panic(err)
	}
}
