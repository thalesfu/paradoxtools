package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/traderoute"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {
	all := traderoute.LoadAllTradeRoutes("/Volumes/[C] Windows 11.hidden/Program Files (x86)/Steam/steamapps/common/Crusader Kings II")

	for _, v := range all {
		fmt.Println(utils.MarshalJSON(v))
	}
}
