package hoi3

import "fmt"

type World struct {
	Provinces map[string]*Province
	Regions   map[string]*Region
	Countries map[string]*Country
	Width     int
	Height    int
}

func NewWorld(saveFilePath string) *World {
	fileLocation := NewFileLocationWithDefault(saveFilePath)

	localisation := GetLocalisation(fileLocation)
	fmt.Println("多语言已经加载")

	regions := LoadRegion(fileLocation, localisation)
	fmt.Println("地区已加载", len(regions))

	provinceList, width, height := LoadProvinces(fileLocation, localisation)

	world := &World{
		Provinces: provinceList,
		Width:     width,
		Height:    height,
		Regions:   regions,
		Countries: make(map[string]*Country),
	}

	for _, province := range provinceList {
		if _, ok := world.Countries[province.Controller]; !ok {
			world.Countries[province.Controller] = &Country{
				ID:        province.Controller,
				Provinces: []string{},
			}
		}
		world.Countries[province.Controller].Provinces = append(world.Countries[province.Controller].Provinces, province.ID)
	}

	return world
}

func (w *World) GetRegionsPolygons(id string) [][]*Point {
	return w.GetProvincesPolygons(w.Regions[id].Provinces)
}

func (w *World) GetCountryPolygons(id string) [][]*Point {
	return w.GetProvincesPolygons(w.Countries[id].Provinces)
}

func (w *World) GetProvincePolygons(id string) [][]*Point {
	return extractPixelBoundary(w.Provinces[id].Pixels, w.Width, w.Height)
}

func (w *World) GetProvincesPolygons(provinceIds []string) [][]*Point {
	pixels := make([]*Point, 0)

	for _, id := range provinceIds {
		if _, ok := w.Provinces[id]; ok {
			pixels = append(pixels, w.Provinces[id].Pixels...)
		}
	}

	return extractPixelBoundary(pixels, w.Width, w.Height)
}
