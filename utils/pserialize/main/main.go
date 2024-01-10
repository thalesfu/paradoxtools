package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/trait"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"regexp"
)

func main() {

	//content, ok := utils.LoadContent("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14.ck2")
	//
	//saveFile, ok := pserialize.UnmarshalP[save.SaveFile](content)
	//
	//if ok {
	//	fmt.Println(utils.MarshalJSON(saveFile.Dynasties[1000103347]))
	//}

	traitContent, ok := utils.LoadContent("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II\\common\\traits\\00_traits.txt")

	t, ok := pserialize.UnmarshalP[map[string]*trait.Trait](traitContent)

	if ok {
		fmt.Println(utils.MarshalJSON(t))
	}

}

func IsCommonRelation(field string, content string) bool {

	re, err := regexp.Compile(field + `=\s*(\{[^\}]*\})`)
	if err != nil {
		return false
	}

	// 使用正则表达式找到所有匹配项
	matches := re.FindAllStringSubmatch(content, -1)

	// 打印所有匹配项
	for _, match := range matches {
		m, o := pserialize.UnmarshalP[map[string]string](match[1])
		if o {
			mm := *m
			_, existD := mm["d"]
			_, existMultiplier := mm["multiplier"]
			if !((len(mm) == 1 && existD) || (len(mm) == 2 && existD && existMultiplier)) {
				return false
			}
		}
	}

	return true
}
