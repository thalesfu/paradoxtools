package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {

	cultureGroups := culture.LoadAllCultures("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cultureGroups))
}
