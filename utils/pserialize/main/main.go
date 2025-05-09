package main

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/hoi3/save"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
)

func main() {

	content, ok := golangutils.LoadContentWithEncoding("/Users/thalesfu/Documents/Paradox Interactive/Hearts of Iron III/ChineseLeaders/save games/CHI1934_07_13_14.hoi3", simplifiedchinese.GB18030)

	saveFile, ok := pserialize.UnmarshalP[save.SaveFile](content)

	if ok {
		fmt.Println(golangutils.MarshalJSON(saveFile))
	}

	for k, v := range saveFile.Provinces {
		if len(v.Infra) == 0 {
			continue
		}
		if v.Infra[0] != v.Infra[1] {
			fmt.Println(k, v.Infra)
		}
	}

	//traitContent, ok := utils.LoadContent("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II/common/traits/00_traits.txt")
	//
	//t, ok := pserialize.UnmarshalP[map[string]*trait.Trait](traitContent)
	//
	//if ok {
	//	fmt.Println(utils.MarshalJSON(t))
	//}

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
