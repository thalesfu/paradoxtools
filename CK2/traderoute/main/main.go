package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/traderoute"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {
	all := traderoute.LoadAllTradeRoutes("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II")

	for _, v := range all {
		fmt.Println(utils.MarshalJSON(v))
	}
}
