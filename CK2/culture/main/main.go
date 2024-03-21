package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {

	cultureGroups := culture.LoadAllCultures("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cultureGroups))
}
