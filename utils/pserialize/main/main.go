package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"regexp"
)

func main() {

	jcontent := "{\"version\":\"3.3.5.1\",\"dyn_title\":[{\"title\":\"k_dyn_146214\"},{\"title\":\"e_dyn_74012\"}],\"dynasties\":{\"1\":{\"coat_of_arms\":{\"data\":[0,0,0,0,0,0,0,0,0,0,0,0,0,0],\"religion\":\"buddhism\"}}}}"

	jsaveFile, ok := utils.UnmarshalJSON[save.SaveFile](jcontent)

	if ok {
		fmt.Println("json", jsaveFile)
	}

	content, ok := utils.LoadContent("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14.ck2")

	saveFile, ok := pserialize.UnmarshalP[save.SaveFile](content)

	if ok {
		fmt.Println("paradox", utils.MarshalJSON(saveFile.Combat.LandCombat[0]))
		fmt.Println("paradox", len(saveFile.Combat.LandCombat))
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
