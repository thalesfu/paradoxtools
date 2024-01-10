package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {

	cultureGroups := religion.LoadAllReligions("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II\\common\\religions")

	fmt.Println(utils.MarshalJSON(cultureGroups))
}
