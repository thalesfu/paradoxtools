package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"github.com/thalesfu/paradoxtools/utils/pserialize"
)

func main() {
	jcontent := "{\"version\":\"3.3.5.1\",\"dyn_title\":[{\"title\":\"k_dyn_146214\"},{\"title\":\"e_dyn_74012\"}]}"

	jsaveFile, ok := utils.UnmarshalJSON[save.SaveFile](jcontent)

	if ok {
		fmt.Println(jsaveFile.Version)
	}

	content, ok := utils.LoadContent("T:\\OneDrive\\fu.thales@live.com\\OneDrive\\MyDocument\\Paradox Interactive\\Crusader Kings II\\save games\\酒泉771_02_14.ck2")

	saveFile, ok := pserialize.UnmarshalP[save.SaveFile](content)

	if !ok {
		fmt.Println(saveFile.Version)
	}

}
