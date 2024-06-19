package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/utils"
	"html/template"
	"os"
	"path"
)

var culturePath = "../CK2Commands/culture"

func main() {

	cultureGroups := culture.LoadAllCultures("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cultureGroups))

	buildCultureFile(cultureGroups)
}

func buildCultureFile(cultureGroups map[string]*culture.CultureGroup) {
	if _, err := os.Stat(culturePath); os.IsNotExist(err) {
		err = os.Mkdir(culturePath, 0755)
		if err != nil {
			panic(err)
		}
	}

	cultureTemplate, err := template.New("CultureTemplate.txt").Funcs(template.FuncMap{"RP": utils.ReplaceTemplateSpecialWords, "ES": utils.EscapeTemplateSpecialWords}).ParseFiles("Ck2/culture/main/CultureTemplate.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path.Join(culturePath, "cultures.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = cultureTemplate.Execute(f, cultureGroups)
	if err != nil {
		panic(err)
	}
}
