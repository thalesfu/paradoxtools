package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
)

func main() {
	ok := utils.IsCompressedFile("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉796_12_01.ck2")

	fmt.Println("Is compressed file:", ok)

	ok = utils.IsCompressedFile("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14dd.ck2")

	fmt.Println("Is compressed file:", ok)

	//err := utils.Unzip("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉796_12_01.ck2", "logs/ck2/save/unzip")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
