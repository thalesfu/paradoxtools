package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {

	cultureGroups := religion.LoadAllReligions("/Volumes/[C] Windows 11.hidden/Program Files (x86)/Steam/steamapps/common/Crusader Kings II")

	fmt.Println(utils.MarshalJSON(cultureGroups))
}
