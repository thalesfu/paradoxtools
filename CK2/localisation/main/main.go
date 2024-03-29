package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
)

func main() {

	translations, repeated := localisation.LoadAllTranslationsDetail("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(len(translations))

	for k, v := range repeated {
		fmt.Printf("%s: %d\n", k, len(v))
		for _, t := range v {
			fmt.Printf("\t%s, %s\n", t.Translation, t.File)
		}
	}

	fmt.Println(len(repeated))

	ts := localisation.LoadAllTranslations("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	fmt.Println(len(ts))
}
