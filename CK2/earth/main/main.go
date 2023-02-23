package main

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"os"
	"path"
	"sync"
)

var earthFolder = "../CK2Commands/earth"

var wg sync.WaitGroup

func main() {
	CreateEarthFile(feud.AllEmpires)

	for _, empire := range feud.AllEmpires {
		wg.Add(1)
		go CreateEmpireFile(empire)
	}

	wg.Wait()
}

func CreateEarthFile(empires map[string]feud.Empire) {
	if _, err := os.Stat(earthFolder); os.IsNotExist(err) {
		err = os.Mkdir(earthFolder, 0755)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(path.Join(earthFolder, "earth.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = EarthTemplate.Execute(f, empires)
	if err != nil {
		panic(err)
	}
}

func CreateBaronyFile(barony feud.Barony) {
	defer wg.Done()
	p := path.Join(earthFolder, barony.GetParent().GetParent().GetParent().GetParent().GetTitle(), barony.GetParent().GetParent().GetParent().GetTitle(), barony.GetParent().GetParent().GetTitle(), barony.GetParent().GetTitle())

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.Mkdir(p, 0755)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(path.Join(p, filterFileName(barony.GetTitleCode())+".go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = BaronyTemplate.Execute(f, barony)
	if err != nil {
		panic(err)
	}
}

func CreateCountyFile(county feud.County) {
	defer wg.Done()
	p := path.Join(earthFolder, county.GetParent().GetParent().GetParent().GetTitle(), county.GetParent().GetParent().GetTitle(), county.GetParent().GetTitle(), county.GetTitle())

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.Mkdir(p, 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, barony := range county.GetChildren() {
		wg.Add(1)
		go CreateBaronyFile(barony)
	}

	f, err := os.OpenFile(path.Join(p, filterFileName(county.GetTitleCode())+".go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = CountyTemplate.Execute(f, county)
	if err != nil {
		panic(err)
	}
}

func CreateDukeFile(duke feud.Duke) {
	defer wg.Done()
	p := path.Join(earthFolder, duke.GetParent().GetParent().GetTitle(), duke.GetParent().GetTitle(), duke.GetTitle())

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.Mkdir(p, 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, county := range duke.GetChildren() {
		wg.Add(1)
		go CreateCountyFile(county)
	}

	f, err := os.OpenFile(path.Join(p, filterFileName(duke.GetTitleCode())+".go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = DukeTemplate.Execute(f, duke)
	if err != nil {
		panic(err)
	}
}

func CreateKingdomFile(kingdom feud.Kingdom) {
	defer wg.Done()
	p := path.Join(earthFolder, kingdom.GetParent().GetTitle(), kingdom.GetTitle())

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.Mkdir(p, 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, duke := range kingdom.GetChildren() {
		wg.Add(1)
		go CreateDukeFile(duke)
	}

	f, err := os.OpenFile(path.Join(p, filterFileName(kingdom.GetTitleCode())+".go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = KingdomTemplate.Execute(f, kingdom)
	if err != nil {
		panic(err)
	}
}

func CreateEmpireFile(empire feud.Empire) {
	defer wg.Done()
	p := path.Join(earthFolder, empire.GetTitle())

	if _, err := os.Stat(p); os.IsNotExist(err) {
		err = os.Mkdir(p, 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, kingdom := range empire.GetChildren() {
		wg.Add(1)
		go CreateKingdomFile(kingdom)
	}

	f, err := os.OpenFile(path.Join(p, filterFileName(empire.GetTitleCode())+".go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = EmpireTemplate.Execute(f, empire)
	if err != nil {
		panic(err)
	}
}

func filterFileName(s string) string {
	switch {
	case s == "b_aix":
		return "b_aix_modify"
	default:
		return s
	}
}
