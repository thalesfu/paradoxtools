package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {

	cultureGroups := culture.LoadAllCultures("/Volumes/[C] Windows 11.hidden/Program Files (x86)/Steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cultureGroups))
}
