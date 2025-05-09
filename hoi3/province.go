package hoi3

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/hoi3/save"
	"image"
	"image/color"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Province struct {
	ID             string
	Color          color.Color
	Pixels         []*Point
	Name           string
	Terrain        string
	Owner          string
	Controller     string
	Core           map[string]bool
	Point          int
	Manpower       float64
	Leadership     float64
	Energy         float64
	Metal          float64
	RareMaterials  float64
	CrudeOil       float64
	AirBase        int
	NavalBase      int
	LandFort       int
	CoastalFort    int
	AntiAir        int
	RadarStation   int
	NuclearReactor int
	RocketTest     int
	Industry       int
	Infra          int
	Supplies       float64
	Fuel           float64
	Strength       map[string]*save.UnitStrength
	UnitCounts     map[string]map[string]*save.UnitCount
}

func LoadProvinces(fileLocation *FileLocation, localisation map[string]string) (map[string]*Province, int, int) {
	provinces := loadProvincesFromMap(fileLocation)
	fmt.Println("加载省份定义完成，共", len(provinces), "个省份")

	for _, province := range provinces {
		province.Name = localisation[fmt.Sprintf("PROV%s", province.ID)]
	}
	fmt.Println("加载省份名称完成")

	provinces, width, height := loadProvincePolygons(fileLocation, provinces)
	fmt.Println("加载省份边界完成")

	loadProvinceTerrain(fileLocation, provinces)

	loadProvinceData(fileLocation, provinces)
	fmt.Println("加载省份数据完成")

	return provinces, width, height
}

func loadProvinceTerrain(location *FileLocation, provinces map[string]*Province) {
	terrainMap := loadTerrainMap(location)

	terrainMap = golangutils.VerticalFlip(terrainMap)

	for _, province := range provinces {
		loadTerrainFromMap(province, terrainMap)
	}

	fmt.Println("加载省份地形完成")
}

func loadTerrainFromMap(province *Province, terrainMap image.Image) {
	if len(province.Pixels) == 0 {
		return
	}

	terrainDetail := make(map[string]int)
	maxCount := 0
	var maxCountTerrain string

	for _, p := range province.Pixels {
		r, g, b, _ := terrainMap.At(p.X, p.Y).RGBA()
		terrain := getTerrainFromColor(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		terrainDetail[terrain]++
		if terrainDetail[terrain] > maxCount {
			maxCount = terrainDetail[terrain]
			maxCountTerrain = terrain
		}
	}

	province.Terrain = maxCountTerrain
}

func getTerrainFromColor(countColor color.RGBA) string {
	switch countColor {
	case color.RGBA{255, 241, 123, 255}:
		return "desert"
	case color.RGBA{255, 236, 71, 255}:
		return "desert"
	case color.RGBA{255, 234, 0, 255}:
		return "desert"
	case color.RGBA{241, 210, 151, 255}:
		return "plains"
	case color.RGBA{236, 236, 236, 255}:
		return "mountain"
	case color.RGBA{235, 179, 233, 255}:
		return "urban"
	case color.RGBA{231, 32, 55, 255}:
		return "mountain"
	case color.RGBA{225, 192, 130, 255}:
		return "plains"
	case color.RGBA{224, 184, 184, 255}:
		return "hills"
	case color.RGBA{218, 195, 0, 255}:
		return "desert"
	case color.RGBA{218, 164, 177, 255}:
		return "hills"
	case color.RGBA{213, 144, 199, 255}:
		return "urban"
	case color.RGBA{210, 210, 210, 255}:
		return "mountain"
	case color.RGBA{208, 101, 69, 255}:
		return "hills"
	case color.RGBA{206, 169, 99, 255}:
		return "plains"
	case color.RGBA{204, 225, 189, 255}:
		return "forest"
	case color.RGBA{203, 164, 164, 255}:
		return "hills"
	case color.RGBA{202, 153, 164, 255}:
		return "hills"
	case color.RGBA{199, 144, 117, 255}:
		return "hills"
	case color.RGBA{195, 187, 95, 255}:
		return "desert"
	case color.RGBA{192, 90, 117, 255}:
		return "mountain"
	case color.RGBA{190, 181, 193, 255}:
		return "marsh"
	case color.RGBA{188, 170, 24, 255}:
		return "desert"
	case color.RGBA{187, 147, 147, 255}:
		return "hills"
	case color.RGBA{187, 97, 0, 255}:
		return "marsh"
	case color.RGBA{187, 90, 62, 255}:
		return "hills"
	case color.RGBA{185, 164, 102, 255}:
		return "desert"
	case color.RGBA{185, 141, 151, 255}:
		return "hills"
	case color.RGBA{181, 200, 167, 255}:
		return "forest"
	case color.RGBA{181, 163, 40, 255}:
		return "desert"
	case color.RGBA{181, 111, 177, 255}:
		return "urban"
	case color.RGBA{180, 130, 106, 255}:
		return "hills"
	case color.RGBA{180, 86, 179, 255}:
		return "urban"
	case color.RGBA{179, 11, 27, 255}:
		return "mountain"
	case color.RGBA{178, 207, 174, 255}:
		return "jungle"
	case color.RGBA{176, 176, 176, 255}:
		return "mountain"
	case color.RGBA{173, 59, 83, 255}:
		return "mountain"
	case color.RGBA{172, 136, 67, 255}:
		return "plains"
	case color.RGBA{171, 89, 0, 255}:
		return "marsh"
	case color.RGBA{167, 80, 54, 255}:
		return "hills"
	case color.RGBA{166, 192, 163, 255}:
		return "jungle"
	case color.RGBA{166, 129, 129, 255}:
		return "hills"
	case color.RGBA{165, 205, 108, 255}:
		return "hills"
	case color.RGBA{165, 154, 168, 255}:
		return "marsh"
	case color.RGBA{164, 125, 134, 255}:
		return "hills"
	case color.RGBA{164, 118, 96, 255}:
		return "hills"
	case color.RGBA{162, 39, 83, 255}:
		return "mountain"
	case color.RGBA{161, 179, 147, 255}:
		return "forest"
	case color.RGBA{160, 212, 220, 255}:
		return "plains"
	case color.RGBA{160, 180, 160, 255}:
		return "jungle"
	case color.RGBA{157, 180, 116, 255}:
		return "hills"
	case color.RGBA{156, 139, 228, 255}:
		return "woods"
	case color.RGBA{153, 79, 0, 255}:
		return "marsh"
	case color.RGBA{152, 211, 131, 255}:
		return "forest"
	case color.RGBA{151, 175, 148, 255}:
		return "jungle"
	case color.RGBA{151, 72, 49, 255}:
		return "hills"
	case color.RGBA{150, 113, 41, 255}:
		return "plains"
	case color.RGBA{149, 107, 87, 255}:
		return "hills"
	case color.RGBA{145, 180, 94, 255}:
		return "hills"
	case color.RGBA{143, 210, 197, 255}:
		return "forest"
	case color.RGBA{143, 172, 200, 255}:
		return "woods"
	case color.RGBA{143, 160, 143, 255}:
		return "jungle"
	case color.RGBA{143, 109, 117, 255}:
		return "hills"
	case color.RGBA{142, 132, 145, 255}:
		return "marsh"
	case color.RGBA{141, 161, 105, 255}:
		return "hills"
	case color.RGBA{140, 140, 140, 255}:
		return "mountain"
	case color.RGBA{139, 154, 128, 255}:
		return "forest"
	case color.RGBA{138, 79, 32, 255}:
		return "mountain"
	case color.RGBA{138, 11, 26, 255}:
		return "mountain"
	case color.RGBA{136, 119, 210, 255}:
		return "woods"
	case color.RGBA{135, 70, 0, 255}:
		return "marsh"
	case color.RGBA{134, 191, 92, 255}:
		return "forest"
	case color.RGBA{133, 160, 186, 255}:
		return "woods"
	case color.RGBA{132, 193, 181, 255}:
		return "forest"
	case color.RGBA{132, 154, 129, 255}:
		return "jungle"
	case color.RGBA{132, 95, 77, 255}:
		return "hills"
	case color.RGBA{128, 160, 83, 255}:
		return "hills"
	case color.RGBA{128, 104, 104, 255}:
		return "hills"
	case color.RGBA{128, 62, 42, 255}:
		return "hills"
	case color.RGBA{127, 24, 60, 255}:
		return "mountain"
	case color.RGBA{125, 140, 125, 255}:
		return "jungle"
	case color.RGBA{124, 66, 14, 255}:
		return "mountain"
	case color.RGBA{123, 90, 27, 255}:
		return "plains"
	case color.RGBA{122, 140, 91, 255}:
		return "hills"
	case color.RGBA{122, 136, 113, 255}:
		return "forest"
	case color.RGBA{122, 92, 99, 255}:
		return "hills"
	case color.RGBA{121, 59, 2, 255}:
		return "mountain"
	case color.RGBA{120, 180, 202, 255}:
		return "plains"
	case color.RGBA{120, 144, 167, 255}:
		return "woods"
	case color.RGBA{119, 150, 138, 255}:
		return "arctic"
	case color.RGBA{118, 245, 217, 255}:
		return "hills"
	case color.RGBA{118, 61, 0, 255}:
		return "marsh"
	case color.RGBA{117, 108, 119, 255}:
		return "marsh"
	case color.RGBA{117, 99, 194, 255}:
		return "woods"
	case color.RGBA{117, 11, 16, 255}:
		return "mountain"
	case color.RGBA{116, 66, 21, 255}:
		return "mountain"
	case color.RGBA{114, 169, 158, 255}:
		return "forest"
	case color.RGBA{114, 142, 74, 255}:
		return "hills"
	case color.RGBA{112, 130, 109, 255}:
		return "jungle"
	case color.RGBA{112, 112, 112, 255}:
		return "mountain"
	case color.RGBA{111, 162, 57, 255}:
		return "forest"
	case color.RGBA{111, 102, 74, 255}:
		return "desert"
	case color.RGBA{111, 80, 65, 255}:
		return "hills"
	case color.RGBA{107, 128, 149, 255}:
		return "woods"
	case color.RGBA{107, 121, 107, 255}:
		return "jungle"
	case color.RGBA{106, 51, 35, 255}:
		return "hills"
	case color.RGBA{105, 54, 0, 255}:
		return "marsh"
	case color.RGBA{104, 119, 77, 255}:
		return "hills"
	case color.RGBA{103, 114, 94, 255}:
		return "forest"
	case color.RGBA{102, 80, 75, 255}:
		return "plains"
	case color.RGBA{101, 150, 140, 255}:
		return "forest"
	case color.RGBA{101, 128, 118, 255}:
		return "arctic"
	case color.RGBA{101, 126, 65, 255}:
		return "hills"
	case color.RGBA{101, 76, 82, 255}:
		return "hills"
	case color.RGBA{101, 71, 15, 255}:
		return "plains"
	case color.RGBA{98, 79, 79, 255}:
		return "hills"
	case color.RGBA{97, 220, 193, 255}:
		return "hills"
	case color.RGBA{97, 209, 66, 255}:
		return "forest"
	case color.RGBA{95, 87, 62, 255}:
		return "desert"
	case color.RGBA{95, 69, 56, 255}:
		return "hills"
	case color.RGBA{94, 113, 131, 255}:
		return "woods"
	case color.RGBA{91, 194, 62, 255}:
		return "forest"
	case color.RGBA{91, 107, 89, 255}:
		return "jungle"
	case color.RGBA{91, 103, 91, 255}:
		return "jungle"
	case color.RGBA{91, 67, 62, 255}:
		return "plains"
	case color.RGBA{91, 12, 158, 255}:
		return "mountain"
	case color.RGBA{89, 131, 123, 255}:
		return "forest"
	case color.RGBA{88, 101, 66, 255}:
		return "hills"
	case color.RGBA{87, 109, 56, 255}:
		return "hills"
	case color.RGBA{86, 124, 27, 255}:
		return "forest"
	case color.RGBA{86, 86, 86, 255}:
		return "mountain"
	case color.RGBA{86, 45, 0, 255}:
		return "marsh"
	case color.RGBA{85, 106, 98, 255}:
		return "arctic"
	case color.RGBA{85, 41, 28, 255}:
		return "hills"
	case color.RGBA{84, 180, 58, 255}:
		return "forest"
	case color.RGBA{84, 46, 12, 255}:
		return "mountain"
	case color.RGBA{84, 33, 127, 255}:
		return "mountain"
	case color.RGBA{83, 100, 116, 255}:
		return "woods"
	case color.RGBA{83, 93, 77, 255}:
		return "forest"
	case color.RGBA{81, 60, 65, 255}:
		return "hills"
	case color.RGBA{80, 65, 146, 255}:
		return "woods"
	case color.RGBA{80, 59, 48, 255}:
		return "hills"
	case color.RGBA{78, 168, 53, 255}:
		return "forest"
	case color.RGBA{78, 78, 78, 255}:
		return "mountain"
	case color.RGBA{78, 71, 50, 255}:
		return "desert"
	case color.RGBA{77, 62, 90, 255}:
		return "mountain"
	case color.RGBA{77, 10, 133, 255}:
		return "mountain"
	case color.RGBA{76, 112, 105, 255}:
		return "forest"
	case color.RGBA{76, 86, 4, 255}:
		return "forest"
	case color.RGBA{75, 210, 208, 255}:
		return "arctic"
	case color.RGBA{75, 147, 174, 255}:
		return "plains"
	case color.RGBA{74, 84, 74, 255}:
		return "jungle"
	case color.RGBA{74, 53, 49, 255}:
		return "plains"
	case color.RGBA{73, 91, 47, 255}:
		return "hills"
	case color.RGBA{73, 50, 6, 255}:
		return "plains"
	case color.RGBA{73, 45, 97, 255}:
		return "mountain"
	case color.RGBA{72, 84, 71, 255}:
		return "jungle"
	case color.RGBA{71, 155, 47, 255}:
		return "forest"
	case color.RGBA{71, 85, 99, 255}:
		return "woods"
	case color.RGBA{71, 81, 54, 255}:
		return "hills"
	case color.RGBA{69, 88, 81, 255}:
		return "arctic"
	case color.RGBA{69, 36, 0, 255}:
		return "marsh"
	case color.RGBA{68, 33, 22, 255}:
		return "hills"
	case color.RGBA{67, 192, 191, 255}:
		return "arctic"
	case color.RGBA{65, 143, 44, 255}:
		return "forest"
	case color.RGBA{65, 53, 53, 255}:
		return "hills"
	case color.RGBA{65, 52, 121, 255}:
		return "woods"
	case color.RGBA{64, 97, 12, 255}:
		return "forest"
	case color.RGBA{63, 93, 87, 255}:
		return "forest"
	case color.RGBA{63, 57, 40, 255}:
		return "desert"
	case color.RGBA{62, 70, 57, 255}:
		return "forest"
	case color.RGBA{62, 43, 39, 255}:
		return "plains"
	case color.RGBA{61, 76, 39, 255}:
		return "hills"
	case color.RGBA{60, 42, 26, 255}:
		return "mountain"
	case color.RGBA{59, 70, 81, 255}:
		return "woods"
	case color.RGBA{59, 39, 21, 255}:
		return "mountain"
	case color.RGBA{58, 170, 169, 255}:
		return "arctic"
	case color.RGBA{58, 129, 39, 255}:
		return "forest"
	case color.RGBA{58, 66, 58, 255}:
		return "jungle"
	case color.RGBA{58, 31, 81, 255}:
		return "mountain"
	case color.RGBA{58, 14, 95, 255}:
		return "mountain"
	case color.RGBA{56, 199, 167, 255}:
		return "hills"
	case color.RGBA{56, 56, 56, 255}:
		return "mountain"
	case color.RGBA{55, 63, 42, 255}:
		return "hills"
	case color.RGBA{53, 33, 29, 255}:
		return "plains"
	case color.RGBA{52, 115, 35, 255}:
		return "forest"
	case color.RGBA{52, 47, 32, 255}:
		return "desert"
	case color.RGBA{51, 41, 41, 255}:
		return "hills"
	case color.RGBA{50, 152, 151, 255}:
		return "arctic"
	case color.RGBA{50, 73, 68, 255}:
		return "forest"
	case color.RGBA{49, 57, 48, 255}:
		return "jungle"
	case color.RGBA{48, 175, 147, 255}:
		return "hills"
	case color.RGBA{45, 134, 133, 255}:
		return "arctic"
	case color.RGBA{45, 119, 146, 255}:
		return "plains"
	case color.RGBA{45, 52, 45, 255}:
		return "jungle"
	case color.RGBA{45, 34, 95, 255}:
		return "woods"
	case color.RGBA{41, 37, 25, 255}:
		return "desert"
	case color.RGBA{40, 23, 19, 255}:
		return "plains"
	case color.RGBA{39, 119, 118, 255}:
		return "arctic"
	case color.RGBA{39, 66, 0, 255}:
		return "forest"
	case color.RGBA{38, 44, 28, 255}:
		return "hills"
	case color.RGBA{38, 12, 60, 255}:
		return "mountain"
	case color.RGBA{37, 96, 126, 255}:
		return "plains"
	case color.RGBA{37, 23, 10, 255}:
		return "mountain"
	case color.RGBA{33, 100, 99, 255}:
		return "arctic"
	case color.RGBA{33, 40, 0, 255}:
		return "forest"
	case color.RGBA{31, 154, 127, 255}:
		return "hills"
	case color.RGBA{30, 27, 18, 255}:
		return "desert"
	case color.RGBA{30, 15, 12, 255}:
		return "plains"
	case color.RGBA{27, 84, 83, 255}:
		return "arctic"
	case color.RGBA{26, 17, 67, 255}:
		return "woods"
	case color.RGBA{24, 0, 255, 255}:
		return "plains"
	case color.RGBA{22, 0, 232, 255}:
		return "plains"
	case color.RGBA{19, 0, 205, 255}:
		return "plains"
	case color.RGBA{18, 8, 6, 255}:
		return "plains"
	case color.RGBA{17, 15, 10, 255}:
		return "desert"
	case color.RGBA{16, 122, 99, 255}:
		return "hills"
	case color.RGBA{16, 11, 41, 255}:
		return "woods"
	case color.RGBA{16, 0, 174, 255}:
		return "plains"
	case color.RGBA{15, 63, 90, 255}:
		return "plains"
	case color.RGBA{14, 0, 150, 255}:
		return "plains"
	case color.RGBA{11, 0, 120, 255}:
		return "plains"
	case color.RGBA{9, 0, 96, 255}:
		return "plains"
	case color.RGBA{6, 41, 78, 255}:
		return "plains"
	case color.RGBA{6, 0, 65, 255}:
		return "plains"
	case color.RGBA{2, 94, 74, 255}:
		return "hills"
	case color.RGBA{2, 20, 41, 255}:
		return "plains"
	case color.RGBA{0, 73, 57, 255}:
		return "hills"
	case color.RGBA{0, 0, 0, 255}:
		return "plains"
	default:
		return ""
	}
}

func loadTerrainMap(fileLocation *FileLocation) image.Image {
	const ProvinceMapPath = "map/terrain.bmp"
	bmp, ok := golangutils.LoadBmpFile(path.Join(fileLocation.ModDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	bmp, ok = golangutils.LoadBmpFile(path.Join(fileLocation.DLCDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	bmp, ok = golangutils.LoadBmpFile(path.Join(fileLocation.BaseDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	return nil
}

func loadProvincesFromMap(fileLocation *FileLocation) map[string]*Province {
	const MapDefinePath = "map/definition.csv"
	content, ok := golangutils.LoadContent(path.Join(fileLocation.ModDirectory, MapDefinePath))
	if ok {
		return generateProvinceFromMapDefineContent(content)
	}

	content, ok = golangutils.LoadContent(path.Join(fileLocation.DLCDirectory, MapDefinePath))
	if ok {
		return generateProvinceFromMapDefineContent(content)
	}

	content, ok = golangutils.LoadContent(path.Join(fileLocation.BaseDirectory, MapDefinePath))
	if ok {
		return generateProvinceFromMapDefineContent(content)
	}

	return make(map[string]*Province)
}

var newlineRegex = regexp.MustCompile(`\r\n?|\n`)

func generateProvinceFromMapDefineContent(content string) map[string]*Province {
	provinces := make(map[string]*Province)

	lines := newlineRegex.Split(content, -1)

	for _, line := range lines[1:] {
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			continue
		}

		rVal, err1 := strconv.Atoi(fields[1])
		gVal, err2 := strconv.Atoi(fields[2])
		bVal, err3 := strconv.Atoi(fields[3])
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		provinces[fields[0]] = &Province{
			ID:         fields[0],
			Color:      color.RGBA{R: uint8(rVal), G: uint8(gVal), B: uint8(bVal), A: 255},
			Core:       make(map[string]bool),
			Strength:   make(map[string]*save.UnitStrength),
			UnitCounts: make(map[string]map[string]*save.UnitCount),
		}
	}

	return provinces
}

func loadProvincesFromSave(fileLocation *FileLocation, provinces map[string]*Province) {
	saveFile, _, err := save.LoadSave(fileLocation.SaveFile)
	if err != nil {
		return
	}

	allProvincesStrength := saveFile.GetProvinceStrength()

	allProvincesUnitCount := saveFile.GetProvinceUnitCount()

	for _, province := range provinces {
		if sp, ok := saveFile.Provinces[province.ID]; ok {
			restProvince(province)
			if sp.Owner != "" {
				province.Owner = sp.Owner
			}

			if sp.Controller != "" {
				province.Controller = sp.Controller
			}

			if len(sp.Core) > 0 {
				for _, country := range sp.Core {
					province.Core[country] = true
				}
			}

			if sp.Manpower > 0 {
				province.Manpower = sp.Manpower
			}

			if sp.Leadership > 0 {
				province.Leadership = sp.Leadership
			}

			if sp.CurrentProducing != nil {
				if sp.CurrentProducing.Energy > 0 {
					province.Energy = sp.CurrentProducing.Energy
				}
				if sp.CurrentProducing.Metal > 0 {
					province.Metal = sp.CurrentProducing.Metal
				}
				if sp.CurrentProducing.RareMaterials > 0 {
					province.RareMaterials = sp.CurrentProducing.RareMaterials
				}
				if sp.CurrentProducing.CrudeOil > 0 {
					province.CrudeOil = sp.CurrentProducing.CrudeOil
				}
			}

			if len(sp.AirBase) > 1 {
				province.AirBase = int(sp.AirBase[1])
			}

			if len(sp.NavalBase) > 1 {
				province.NavalBase = int(sp.NavalBase[1])
			}

			if len(sp.LandFort) > 1 {
				province.LandFort = int(sp.LandFort[1])
			}

			if len(sp.CoastalFort) > 1 {
				province.CoastalFort = int(sp.CoastalFort[1])
			}

			if len(sp.AntiAir) > 1 {
				province.AntiAir = int(sp.AntiAir[1])
			}

			if len(sp.RadarStation) > 1 {
				province.RadarStation = int(sp.RadarStation[1])
			}

			if len(sp.NuclearReactor) > 1 {
				province.NuclearReactor = int(sp.NuclearReactor[1])
			}

			if len(sp.RocketTest) > 1 {
				province.RocketTest = int(sp.RocketTest[1])
			}

			if len(sp.Industry) > 1 {
				province.Industry = int(sp.Industry[1])
			}

			if len(sp.Infra) > 1 {
				province.Infra = int(sp.Infra[1])
			}

			if sp.Pool != nil {
				if sp.Pool.Supplies > 0 {
					province.Supplies = sp.Pool.Supplies
				}
				if sp.Pool.Fuel > 0 {
					province.Fuel = sp.Pool.Fuel
				}
			}

			if strength, ok := allProvincesStrength[province.ID]; ok {
				province.Strength = strength
			}

			if unitCount, ok := allProvincesUnitCount[province.ID]; ok {
				province.UnitCounts = unitCount
			}
		}
	}
}

func loadProvincePolygons(fileLocation *FileLocation, provinces map[string]*Province) (map[string]*Province, int, int) {
	mapBmp := loadProvinceMap(fileLocation)
	if mapBmp == nil {
		return provinces, 0, 0
	}

	mapBmp = golangutils.VerticalFlip(mapBmp)

	bounds := mapBmp.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 构建颜色到像素点的映射
	colorMap := buildColorMap(mapBmp)

	provinceIndex := 0
	for _, p := range provinces {
		provinceIndex++
		p.Pixels = colorMap[p.Color]
	}

	return provinces, width, height
}

func loadProvinceMap(fileLocation *FileLocation) image.Image {
	const ProvinceMapPath = "map/provinces.bmp"
	bmp, ok := golangutils.LoadBmpFile(path.Join(fileLocation.ModDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	bmp, ok = golangutils.LoadBmpFile(path.Join(fileLocation.DLCDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	bmp, ok = golangutils.LoadBmpFile(path.Join(fileLocation.BaseDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	return nil
}

// 构建颜色到像素点的映射
func buildColorMap(img image.Image) map[color.Color][]*Point {
	bounds := img.Bounds()
	colorMap := make(map[color.Color][]*Point)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//if x == 4728 && y == 874 {
			//	fmt.Println(x, y)
			//}
			c := img.At(x, y)
			colorMap[c] = append(colorMap[c], &Point{x, y})
		}
	}
	return colorMap
}

// 提取像素级边界(在8连通下，边界仍然同样处理，只是填充区域可能更完整)
func extractPixelBoundary(points []*Point, width, height int) [][]*Point {
	filled := buildFilledMap(points, width, height)

	type Edge struct {
		Start, End Point
	}

	var edges []Edge

	for _, p := range points {
		x, y := p.X, p.Y
		// 与之前相同的逻辑，因为边界定义仍然基于上下左右的空隙
		// 上边界
		if y == 0 || !filled[y-1][x] {
			edges = append(edges, Edge{Start: Point{x, y}, End: Point{x + 1, y}})
		}
		// 下边界
		if y == height-1 || !filled[y+1][x] {
			edges = append(edges, Edge{Start: Point{x, y + 1}, End: Point{x + 1, y + 1}})
		}
		// 左边界
		if x == 0 || !filled[y][x-1] {
			edges = append(edges, Edge{Start: Point{x, y}, End: Point{x, y + 1}})
		}
		// 右边界
		if x == width-1 || !filled[y][x+1] {
			edges = append(edges, Edge{Start: Point{x + 1, y}, End: Point{x + 1, y + 1}})
		}
	}

	adj := make(map[Point][]Point)
	for _, e := range edges {
		adj[e.Start] = append(adj[e.Start], e.End)
		adj[e.End] = append(adj[e.End], e.Start)
	}

	result := make([][]*Point, 0)

	for {
		polygon, rest := generatePolygon(adj)
		if polygon != nil {
			result = append(result, polygon)
		}
		if len(rest) == 0 {
			break
		}
		adj = rest
	}

	return result
}

func buildFilledMap(points []*Point, width, height int) [][]bool {
	filled := make([][]bool, height)
	for i := range filled {
		filled[i] = make([]bool, width)
	}
	for _, p := range points {
		if p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height {
			filled[p.Y][p.X] = true
		}
	}
	return filled
}

func generatePolygon(adj map[Point][]Point) (polygon []*Point, rest map[Point][]Point) {
	var startPoint Point
	var maxEdgeCount int
	for k, v := range adj {
		if len(v) > maxEdgeCount {
			startPoint = k
			maxEdgeCount = len(v)
		}
	}

	if (startPoint == Point{}) {
		return nil, nil
	}

	polygon = []*Point{&startPoint}
	current := startPoint
	var prev Point

	for {
		nextList := adj[current]
		var next Point
		next, found := lo.Find(nextList, func(candidate Point) bool {
			return candidate != prev
		})

		if !found {
			return nil, adj
		}

		i := lo.IndexOf(adj[current], next)
		adj[current] = append(nextList[:i], nextList[i+1:]...)
		if len(adj[current]) == 0 {
			delete(adj, current)
		}

		i = lo.IndexOf(adj[next], current)
		adj[next] = append(adj[next][:i], adj[next][i+1:]...)
		if len(adj[next]) == 0 {
			delete(adj, next)
		}

		if next == startPoint && len(polygon) > 2 {
			break
		}

		polygon = append(polygon, &next)
		prev = current
		current = next
	}

	return polygon, adj
}

func loadProvinceData(fileLocation *FileLocation, provinces map[string]*Province) {
	const ProvinceDataPath = "history/provinces"

	p := path.Join(fileLocation.BaseDirectory, ProvinceDataPath)
	if _, err := os.Stat(p); err == nil {
		fmt.Println("从基础文件中加载省份数据")
		loadProvinceDataFromDir(p, provinces)
	}

	p = path.Join(fileLocation.DLCDirectory, ProvinceDataPath)
	if _, err := os.Stat(p); err == nil {
		fmt.Println("从DLC文件中加载省份数据")
		loadProvinceDataFromDir(p, provinces)
	}

	p = path.Join(fileLocation.ModDirectory, ProvinceDataPath)
	if _, err := os.Stat(p); err == nil {
		fmt.Println("从MOD文件中加载省份数据")
		loadProvinceDataFromDir(p, provinces)
	}

	if _, err := os.Stat(fileLocation.SaveFile); err == nil {
		fmt.Println("从存档文件中加载省份数据")
		loadProvincesFromSave(fileLocation, provinces)
	}
}

var provinceIdRegex = regexp.MustCompile(`^\d+`)

func loadProvinceDataFromDir(path string, provinces map[string]*Province) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("无法访问路径 %s: %v\n", path, err)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		//正则表达式提取文件命中的数字
		provinceID := provinceIdRegex.FindString(info.Name())
		if provinceID == "" {
			return nil
		}

		if province, ok := provinces[provinceID]; ok {
			loadProvinceDataFromSingleFile(path, province)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}

}

func loadProvinceDataFromSingleFile(path string, province *Province) {
	//if province.ID == "9223" {
	//	fmt.Println("Debug load province data from single file", province.ID)
	//}

	rest := false
	content, ok := golangutils.LoadContent(path)
	if !ok {
		return
	}

	lines := newlineRegex.Split(content, -1)

	for _, line := range lines {
		fields := strings.Split(line, "=")
		if len(fields) < 2 {
			continue
		}

		if !rest {
			restProvince(province)
			rest = true
		}

		switch strings.Trim(fields[0], " ") {
		case "owner":
			province.Owner = strings.Trim(fields[1], " ")
		case "controller":
			province.Controller = strings.Trim(fields[1], " ")
		case "add_core":
			province.Core[fields[1]] = true
		case "manpower":
			province.Manpower, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "leadership":
			province.Leadership, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "energy":
			province.Energy, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "metal":
			province.Metal, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "rare_materials":
			province.RareMaterials, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "crude_oil":
			province.CrudeOil, _ = strconv.ParseFloat(strings.Trim(fields[1], " "), 32)
		case "air_base":
			province.AirBase, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "naval_base":
			province.NavalBase, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "land_fort":
			province.LandFort, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "coastal_fort":
			province.CoastalFort, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "anti_air":
			province.AntiAir, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "radar_station":
			province.RadarStation, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "nuclear_reactor":
			province.NuclearReactor, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "rocket_test":
			province.RocketTest, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "industry":
			province.Industry, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "infra":
			province.Infra, _ = strconv.Atoi(strings.Trim(fields[1], " "))
		case "terrain":
			terrain := strings.Trim(fields[1], " ")
			if terrain != "" {
				province.Terrain = terrain
			}
		}
	}
}

func restProvince(province *Province) {
	province.Owner = ""
	province.Controller = ""
	province.Core = make(map[string]bool)
	province.Manpower = 0
	province.Leadership = 0
	province.Energy = 0
	province.Metal = 0
	province.RareMaterials = 0
	province.CrudeOil = 0
	province.AirBase = 0
	province.NavalBase = 0
	province.LandFort = 0
	province.CoastalFort = 0
	province.AntiAir = 0
	province.RadarStation = 0
	province.NuclearReactor = 0
	province.RocketTest = 0
	province.Industry = 0
	province.Infra = 0
}
