package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/localisation"
)

func main() {

	translations := localisation.LoadAllTranslations("R:\\Thales\\Game\\SteamLibrary\\steamapps\\common\\Crusader Kings II\\localisation")

	fmt.Println(len(translations))
}
