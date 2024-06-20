package main

import (
	"github.com/thalesfu/paradoxtools"
	"github.com/thalesfu/paradoxtools/CK2/history/dynasty"
	"github.com/thalesfu/paradoxtools/utils"
	"html/template"
	"os"
	"path"
)

var historyDynastyPath = "../CK2Commands/historydynasty"

func main() {
	dynasties := dynasty.LoadAllHistoryDynasty(paradoxtools.Ck2path)

	//fmt.Println(utils.MarshalJSON(dynasties))

	buildHistoryDynastyFile(dynasties)
}

func buildHistoryDynastyFile(historyDynasty map[int]*dynasty.HistoryDynasty) {
	if _, err := os.Stat(historyDynastyPath); os.IsNotExist(err) {
		err = os.Mkdir(historyDynastyPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	t, err := template.New("HistoryDynastyTemplate.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/history/dynasty/main/HistoryDynastyTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(historyDynastyPath, "dynasty.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, historyDynasty)
	if err != nil {
		panic(err)
	}
}
