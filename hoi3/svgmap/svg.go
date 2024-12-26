package svgmap

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/samber/lo"
	"github.com/thalesfu/paradoxtools/hoi3"
	"math"
	"os"
	"strings"
)

type ProvincePosition struct {
	X, Y     int
	Province *hoi3.Province
}

func GenerateCountrySVG(world *hoi3.World, outputPath string, scale int, countries ...string) error {
	polygons := make([][]*hoi3.Point, 0)

	for _, c := range countries {
		if country, ok := world.Countries[c]; ok {
			if len(country.Provinces) > 0 {
				for _, p := range country.Provinces {
					if province, ok := world.Provinces[p]; ok {
						polygons = append(polygons, province.Pixels)
					}
				}
			}
		}
	}

	xmin, ymin, xmax, ymax, width, height := getPolygonSharp(polygons, 50, world.Width, world.Height)

	newWorld := &hoi3.World{
		Provinces: map[string]*hoi3.Province{},
		Regions:   map[string]*hoi3.Region{},
		Countries: map[string]*hoi3.Country{},
		Width:     width,
		Height:    height,
	}

	for _, province := range world.Provinces {
		pixels := make([]*hoi3.Point, 0)
		for _, p := range province.Pixels {
			if p.X >= xmin && p.X < xmax && p.Y >= ymin && p.Y < ymax {
				pixels = append(pixels, &hoi3.Point{
					X: p.X - xmin,
					Y: p.Y - ymin,
				})
			}
		}
		if len(pixels) > 0 {
			newWorld.Provinces[province.ID] = &hoi3.Province{
				ID:             province.ID,
				Name:           province.Name,
				Color:          province.Color,
				Terrain:        province.Terrain,
				Pixels:         pixels,
				Controller:     province.Controller,
				Core:           province.Core,
				Point:          province.Point,
				Manpower:       province.Manpower,
				Leadership:     province.Leadership,
				Energy:         province.Energy,
				Metal:          province.Metal,
				RareMaterials:  province.RareMaterials,
				CrudeOil:       province.CrudeOil,
				AirBase:        province.AirBase,
				NavalBase:      province.NavalBase,
				LandFort:       province.LandFort,
				CoastalFort:    province.CoastalFort,
				AntiAir:        province.AntiAir,
				RadarStation:   province.RadarStation,
				NuclearReactor: province.NuclearReactor,
				RocketTest:     province.RocketTest,
				Industry:       province.Industry,
				Infra:          province.Infra,
				Supplies:       province.Supplies,
				Fuel:           province.Fuel,
				Strength:       province.Strength,
			}
		}
	}

	for _, region := range world.Regions {
		provinces := make([]string, 0)
		for _, p := range region.Provinces {
			if _, ok := newWorld.Provinces[p]; ok {
				provinces = append(provinces, p)
			}
		}
		if len(provinces) > 0 {
			newWorld.Regions[region.ID] = &hoi3.Region{
				ID:        region.ID,
				Name:      region.Name,
				Provinces: provinces,
			}
		}
	}

	for _, country := range world.Countries {
		provinces := make([]string, 0)
		for _, p := range country.Provinces {
			if _, ok := newWorld.Provinces[p]; ok {
				provinces = append(provinces, p)
			}
		}
		if len(provinces) > 0 {
			newWorld.Countries[country.ID] = &hoi3.Country{
				ID:        country.ID,
				Provinces: provinces,
			}
		}
	}

	return GenerateDetailSVG(newWorld, outputPath, scale)
}

func GenerateDetailSVG(world *hoi3.World, outputPath string, scale int) error {

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(world.Width*scale, world.Height*scale)
	canvas.Rect(0, 0, world.Width, world.Height, "fill:white")

	positions := make([]ProvincePosition, 0)

	provinceIndex := 0
	for _, province := range world.Provinces {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}
		provinceIndex++

		provincePolygons := world.GetProvincePolygons(province.ID)
		polygons := lo.Filter(provincePolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		regionxs := make([]int, 0)
		regionys := make([]int, 0)

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)
			for _, p := range polygonsMap[i] {
				regionxs = append(regionxs, p.X*scale)
				regionys = append(regionys, p.Y*scale)
			}

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
					for _, p := range polygonsMap[c] {
						regionxs = append(regionxs, p.X*scale)
						regionys = append(regionys, p.Y*scale)
					}
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:%s;stroke:#464646;stroke-width:1;fill-rule:evenodd\"", province.ID, getTerrainColor(province)))
		}

		centerX, centerY := getPolygonCenter(regionxs, regionys)

		positions = append(positions, ProvincePosition{
			X:        centerX,
			Y:        centerY,
			Province: province,
		})

		canvas.Gend()
	}

	//区域边界
	canvas.Group()
	regionIndex := 0
	for _, region := range world.Regions {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}

		regionIndex++

		regionPolygons := world.GetRegionsPolygons(region.ID)
		polygons := lo.Filter(regionPolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:none;stroke:#323232;stroke-width:2;fill-rule:evenodd\"", region.ID))
		}

		canvas.Gend()
	}
	canvas.Gend()

	//区域边界
	canvas.Group()
	countryIndex := 0
	for _, country := range world.Countries {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}
		countryIndex++

		countryPolygons := world.GetCountryPolygons(country.ID)
		polygons := lo.Filter(countryPolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:none;stroke:#74fff5;stroke-width:3;fill-rule:evenodd;stroke-dasharray:8,8;\"", country.ID))
		}

		canvas.Gend()
	}
	canvas.Gend()

	//省名
	canvas.Group()
	fontSize := 8
	fontPadding := int(math.Ceil(float64(fontSize) / 4))
	for _, position := range positions {
		baseBuilder := strings.Builder{}
		if position.Province.Infra < 2 {
			baseBuilder.WriteString("基:X")
		} else {
			baseBuilder.WriteString(fmt.Sprintf("基:%d", position.Province.Infra))
		}

		if position.Province.Industry > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 厂:%d", position.Province.Industry))
		}

		if position.Province.RadarStation > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 雷:%d", position.Province.RadarStation))
		}

		if position.Province.RocketTest > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 射:%d", position.Province.RocketTest))
		}

		if position.Province.NuclearReactor > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 核:%d", position.Province.NuclearReactor))
		}

		if position.Province.AirBase > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 空:%d", position.Province.AirBase))
		}

		if position.Province.NavalBase > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 海:%d", position.Province.NavalBase))
		}

		if position.Province.LandFort > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 堡:%d", position.Province.LandFort))
		}

		if position.Province.CoastalFort > 0 {
			baseBuilder.WriteString(fmt.Sprintf(" 岸:%d", position.Province.CoastalFort))
		}

		countrySimpleName := getCountrySimpleName(position.Province.Controller)
		name := position.Province.Name
		if countrySimpleName != "" {
			name = fmt.Sprintf("[%s] %s", countrySimpleName, position.Province.Name)
		}

		maxWordCount := len(name)
		if len(baseBuilder.String()) > maxWordCount {
			maxWordCount = len(baseBuilder.String())
		}

		// 计算背景矩形的宽度和高度
		textWidth := maxWordCount * fontSize / 2 // 简单估算宽度（可以根据具体字体调整）
		textHeight := fontSize

		// 绘制背景矩形
		rectWidth := textWidth + int(math.Ceil(float64(fontPadding)/2))
		rectHeight := textHeight*2 + 6*fontPadding
		rectX := position.X - int(math.Ceil(float64(rectWidth)/2))
		rectY := position.Y - int(math.Ceil(float64(rectHeight)/2))

		canvas.Rect(rectX, rectY, rectWidth, rectHeight, fmt.Sprintf("fill:%s;rx:16;ry:16", getInfraColor(position.Province))) // 添加圆角 rx 和 ry 可选

		canvas.Text(position.X, position.Y-textHeight+fontPadding*2, name, fmt.Sprintf("font-family:PingFangSC-Medium-GBpc-EUC-H; font-size:%dpx; fill:white; text-anchor:middle; alignment-baseline:middle", fontSize))
		canvas.Text(position.X, position.Y+textHeight-int(math.Ceil(float64(fontPadding)/2)), baseBuilder.String(), fmt.Sprintf("font-family:PingFangSC-Medium-GBpc-EUC-H; font-size:%dpx; fill:white; text-anchor:middle; alignment-baseline:middle", fontSize))

		if len(position.Province.Strength) > 0 {
			strengthFontSize := 12
			strengthBuilder := strings.Builder{}
			si := 0
			for c, s := range position.Province.Strength {
				if si > 0 {
					strengthBuilder.WriteString(" ")
				}
				strengthBuilder.WriteString(fmt.Sprintf("%s:%.2f,%.2f", getCountrySimpleName(c), s.Strength, s.GetAverageOrganisation()))
				si++
			}

			canvas.Text(position.X, rectY+rectHeight+strengthFontSize, strengthBuilder.String(), fmt.Sprintf("font-family:PingFangSC-Medium-GBpc-EUC-H; font-size:%dpx; fill:#FF40FF; font-weight:bold; text-anchor:middle; alignment-baseline:middle", strengthFontSize))
		}
	}
	canvas.Gend()
	canvas.End()
	return nil
}

func getCountrySimpleName(country string) string {
	country = strings.ToUpper(country)
	switch country {
	case "AFG":
		return "阿"
	case "AUS":
		return "澳"
	case "CHC":
		return "共"
	case "CHI":
		return "国"
	case "CSX":
		return "晋"
	case "ENG":
		return "英"
	case "FRA":
		return "法"
	case "GER":
		return "德"
	case "HOL":
		return "荷"
	case "ITA":
		return "意"
	case "JAP":
		return "日"
	case "MAN":
		return "满"
	case "PHI":
		return "菲"
	case "SIA":
		return "泰"
	case "SOV":
		return "苏"
	case "USA":
		return "美"
	}
	return country
}

func GenerateSVG(world *hoi3.World, outputPath string, scale int) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(world.Width*scale, world.Height*scale)
	canvas.Rect(0, 0, world.Width, world.Height, "fill:white")

	positions := make([]ProvincePosition, 0)

	provinceIndex := 0
	for _, province := range world.Provinces {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}
		provinceIndex++

		provincePolygons := world.GetProvincePolygons(province.ID)
		polygons := lo.Filter(provincePolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		regionxs := make([]int, 0)
		regionys := make([]int, 0)

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)
			for _, p := range polygonsMap[i] {
				regionxs = append(regionxs, p.X*scale)
				regionys = append(regionys, p.Y*scale)
			}

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
					for _, p := range polygonsMap[c] {
						regionxs = append(regionxs, p.X*scale)
						regionys = append(regionys, p.Y*scale)
					}
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:%s;stroke:#464646;stroke-width:1;fill-rule:evenodd\"", province.ID, getTerrainColor(province)))
		}

		centerX, centerY := getPolygonCenter(regionxs, regionys)

		positions = append(positions, ProvincePosition{
			X: centerX,
			Y: centerY,
		})

		canvas.Gend()
	}

	//区域边界
	canvas.Group()
	regionIndex := 0
	for _, region := range world.Regions {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}

		regionIndex++

		regionPolygons := world.GetRegionsPolygons(region.ID)
		polygons := lo.Filter(regionPolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:none;stroke:#323232;stroke-width:2;fill-rule:evenodd\"", region.ID))
		}

		canvas.Gend()
	}
	canvas.Gend()

	//区域边界
	canvas.Group()
	countryIndex := 0
	for _, country := range world.Countries {
		//调试断点
		//if province.ID == "11585" {
		//	fmt.Println("Debug generate svg", province.ID, province.Name)
		//}
		countryIndex++

		countryPolygons := world.GetCountryPolygons(country.ID)
		polygons := lo.Filter(countryPolygons, func(p []*hoi3.Point, _ int) bool {
			return len(p) >= 3
		})

		if len(polygons) == 0 {
			continue
		}

		containedPolygons := findContainedPolygons(polygons)
		polygonsMap := make(map[int][]*hoi3.Point)
		for i, p := range polygons {
			polygonsMap[i] = p
		}

		canvas.Group()

		for i, contained := range containedPolygons {
			path := polygonToPath(polygonsMap[i], false, scale)

			if len(contained) > 0 {
				for _, c := range contained {
					path = path + " " + polygonToPath(polygonsMap[c], true, scale)
				}
			}

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:none;stroke:#7F2D2C;stroke-width:3;fill-rule:evenodd;stroke-dasharray:5,5;\"", country.ID))
		}

		canvas.Gend()
	}
	canvas.Gend()

	//省名
	canvas.Group()
	for _, position := range positions {
		canvas.Text(position.X, position.Y, position.Province.Name, "font-family:PingFangSC-Medium-GBpc-EUC-H; font-size:8pt; fill:white; text-anchor:middle; alignment-baseline:middle")
	}
	canvas.Gend()
	canvas.End()
	return nil
}

func getTerrainName(terrain string) string {
	switch terrain {
	case "mountain":
		return "山脉"
	case "forest":
		return "森林"
	case "woods":
		return "林地"
	case "marsh":
		return "沼泽"
	case "plains":
		return "平原"
	case "urban":
		return "城市"
	case "hills":
		return "丘陵"
	case "jungle":
		return "丛林"
	case "desert":
		return "沙漠"
	case "arctic":
		return "北极"
	default:
		return "海"
	}
}

func getDefaultColor(province *hoi3.Province) string {
	r, g, b, _ := province.Color.RGBA()
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	return fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)
}

func getTerrainColor(province *hoi3.Province) string {
	switch province.Terrain {
	case "mountain":
		return "#756c77"
	case "forest":
		return "#5b7b2d"
	case "woods":
		return "#a5cd6c"
	case "marsh":
		return "#4c7069"
	case "plains":
		return "#f1ddb8"
	case "urban":
		return "#8968a5"
	case "hills":
		return "#874600"
	case "jungle":
		return "#209700"
	case "desert":
		return "#dac300"
	case "arctic":
		return "#ebebeb"
	default:
		return "#0000ff"
	}
}

func getInfraColor(province *hoi3.Province) string {
	switch province.Infra {
	case 0:
		return "#010101"
	case 1:
		return "#242424"
	case 2:
		return "#601616"
	case 3:
		return "#924136"
	case 4:
		return "#a55433"
	case 5:
		return "#a8a835"
	case 6:
		return "#636f18"
	case 7:
		return "#52743b"
	case 8:
		return "#31581f"
	case 9:
		return "#3d8433"
	case 10:
		return "#49b048"
	default:
		return "#000000"
	}
}

// 计算多边形的中心点
func getPolygonCenter(xPoints, yPoints []int) (int, int) {
	var sumX, sumY int
	for i := 0; i < len(xPoints); i++ {
		sumX += xPoints[i]
		sumY += yPoints[i]
	}
	centerX := sumX / len(xPoints)
	centerY := sumY / len(yPoints)
	return centerX, centerY
}

// 判断一个点是否在一个多边形内
func isPointInPolygon(point *hoi3.Point, polygon []*hoi3.Point) bool {
	n := len(polygon)
	if n < 3 {
		return false // 一个多边形至少需要三个顶点
	}

	inside := false
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		pi, pj := polygon[i], polygon[j]
		if ((pi.Y > point.Y) != (pj.Y > point.Y)) &&
			(point.X < (pj.X-pi.X)*(point.Y-pi.Y)/(pj.Y-pi.Y)+pi.X) {
			inside = !inside
		}
	}
	return inside
}

// 判断多边形 A 是否包含多边形 B
func isPolygonContained(polygonA, polygonB []*hoi3.Point) bool {
	for _, point := range polygonB {
		if !isPointInPolygon(point, polygonA) {
			return false // 只要有一个点不在 A 内，B 就不被包含
		}
	}
	return true
}

// 判断多个多边形的包含关系
func findContainedPolygons(polygons [][]*hoi3.Point) map[int][]int {
	containedMap := make(map[int][]int)
	containedSlice := make([]int, 0)

	for i, polyA := range polygons {
		containedMap[i] = []int{}
		for j, polyB := range polygons {
			if i != j && isPolygonContained(polyA, polyB) {
				containedMap[i] = append(containedMap[i], j)
				containedSlice = append(containedSlice, j)
			}
		}
	}

	for _, contained := range containedSlice {
		delete(containedMap, contained)
	}

	return containedMap
}

// 将多边形顶点转换为 SVG 路径指令
func polygonToPath(polygon []*hoi3.Point, reverse bool, scale int) string {
	path := ""
	n := len(polygon)

	for i := 0; i < n; i++ {
		point := polygon[i]
		if reverse {
			point = polygon[n-1-i] // 反向绘制
		}

		if i == 0 {
			path += "M " // 移动到起点
		} else {
			path += "L " // 连线到下一个点
		}
		path += fmt.Sprintf("%d %d ", point.X*scale, point.Y*scale)
	}
	path += "Z" // 闭合路径
	return path
}

func getPolygonSharp(polygons [][]*hoi3.Point, padding int, originalWidth int, originalHeight int) (xmin int, ymin int, xmax int, ymax int, width int, height int) {
	if len(polygons) == 0 {
		return
	}
	xmin = math.MaxInt
	ymin = math.MaxInt
	xmax = math.MinInt
	ymax = math.MinInt

	for _, polygon := range polygons {
		for _, p := range polygon {
			if p.X < xmin {
				xmin = p.X
			}
			if p.X > xmax {
				xmax = p.X
			}
			if p.Y < ymin {
				ymin = p.Y
			}
			if p.Y > ymax {
				ymax = p.Y
			}
		}
	}

	xmin -= padding
	if xmin < 0 {
		xmin = 0
	}

	ymin -= padding
	if ymin < 0 {
		ymin = 0
	}

	xmax += padding
	if xmax > originalWidth {
		xmax = originalWidth
	}

	ymax += padding
	if ymax > originalHeight {
		ymax = originalHeight
	}

	width = xmax - xmin
	height = ymax - ymin

	return
}
