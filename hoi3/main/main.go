package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/hoi3"
	"github.com/thalesfu/paradoxtools/hoi3/svgmap"
)

func main() {
	outputSVG := "/Users/thalesfu/Documents/My Games/HOI3/province-map-01.svg"
	saveFilePath := "/Users/thalesfu/Documents/Paradox Interactive/Hearts of Iron III/ChineseLeadersEN/save games/Nationalist_China1935_08_28_20.hoi3"

	world := hoi3.NewWorld(saveFilePath)

	err := svgmap.GenerateCountrySVG(world, outputSVG, 8, "CHI", "JAP")
	if err != nil {
		fmt.Println("生成SVG错误:", err)
		return
	}

	fmt.Println("SVG文件已生成:", outputSVG)
}
