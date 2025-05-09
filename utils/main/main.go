package main

import (
	"fmt"
	"github.com/thalesfu/golangutils"
)

func main() {
	ok := golangutils.IsCompressedFile("/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉796_12_01.ck2")

	fmt.Println("Is compressed file:", ok)

	ok = golangutils.IsCompressedFile("/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉771_02_14dd.ck2")

	fmt.Println("Is compressed file:", ok)

	//err := utils.Unzip("/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games/酒泉796_12_01.ck2", "logs/ck2/save/unzip")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
