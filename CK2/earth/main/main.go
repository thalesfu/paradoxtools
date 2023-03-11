package main

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"os"
	"path"
	"sync"
)

var earthFolder = "../CK2Commands/earth"

var wg sync.WaitGroup

const (
	BuildEarthFile   = 0b000001
	BuildEmpireFile  = 0b000010
	BuildKingdomFile = 0b000100
	BuildDukeFile    = 0b001000
	BuildCountyFile  = 0b010000
	BuildBaronyFile  = 0b100000
)

var buildFlags byte

func main() {

	//buildFlags = BuildEarthFile
	buildFlags = BuildEarthFile | BuildEmpireFile | BuildKingdomFile | BuildDukeFile | BuildCountyFile | BuildBaronyFile

	CreateEarthFile(AllEmpires)

	for _, empire := range AllEmpires {
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

	if canBuildEarth() {
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
}

func canBuildEarth() bool {
	return (buildFlags & BuildEarthFile) == BuildEarthFile
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

	if canBuildBarony() {
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

	if canBuildCounty() {
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

	if canBuildDuke() {
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

	if canBuildKingdom() {
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

	if canBuildEmpire() {
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
}

func filterFileName(s string) string {
	switch {
	case s == "b_aix":
		return "b_aix_modify"
	default:
		return s
	}
}

func canBuildBarony() bool {

	if buildFlags&BuildEmpireFile != BuildEmpireFile {
		return false
	}

	if buildFlags&BuildKingdomFile != BuildKingdomFile {
		return false

	}

	if buildFlags&BuildDukeFile != BuildDukeFile {
		return false
	}

	if buildFlags&BuildCountyFile != BuildCountyFile {
		return false
	}

	return (buildFlags & BuildBaronyFile) == BuildBaronyFile
}

func canBuildCounty() bool {

	if buildFlags&BuildEmpireFile != BuildEmpireFile {
		return false
	}

	if buildFlags&BuildKingdomFile != BuildKingdomFile {
		return false

	}

	if buildFlags&BuildDukeFile != BuildDukeFile {
		return false
	}

	return (buildFlags & BuildCountyFile) == BuildCountyFile
}

func canBuildDuke() bool {

	if buildFlags&BuildEmpireFile != BuildEmpireFile {
		return false
	}

	if buildFlags&BuildKingdomFile != BuildKingdomFile {
		return false

	}

	return (buildFlags & BuildDukeFile) == BuildDukeFile
}

func canBuildKingdom() bool {
	if buildFlags&BuildEmpireFile != BuildEmpireFile {
		return false
	}

	return (buildFlags & BuildKingdomFile) == BuildKingdomFile
}

func canBuildEmpire() bool {
	return (buildFlags & BuildEmpireFile) == BuildEmpireFile
}
