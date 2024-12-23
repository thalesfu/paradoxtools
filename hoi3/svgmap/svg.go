package svgmap

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/samber/lo"
	"github.com/thalesfu/paradoxtools/hoi3"
	"os"
)

type NameText struct {
	Name string
	X, Y int
}

func GenerateSVG(allProvince map[string]*hoi3.Province, width, height int, outputPath string, scale int) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(width*scale, height*scale)
	canvas.Rect(0, 0, width, height, "fill:white")

	names := make([]NameText, 0)

	provinceIndex := 0
	for _, province := range allProvince {
		//调试断点
		if province.ID == "11585" {
			fmt.Println("Debug", province.ID, province.Name)
		}

		provinceIndex++
		polygons := lo.Filter(province.Polygons, func(p []*hoi3.Point, _ int) bool {
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

			canvas.Path(path, fmt.Sprintf("id=\"%s\" style=\"fill:%s;stroke:#464646;stroke-width:1;fill-rule:evenodd\"", province.ID, getDefaultColor(province)))
		}

		centerX, centerY := getPolygonCenter(regionxs, regionys)

		names = append(names, NameText{
			Name: province.ID,
			X:    centerX,
			Y:    centerY,
		})

		canvas.Gend()
		fmt.Printf("绘制 %d-%d %s\n", provinceIndex, len(allProvince), province.Name)
	}

	canvas.Group()
	for _, name := range names {
		canvas.Text(name.X, name.Y, name.Name, "font-family:PingFangSC-Medium-GBpc-EUC-H; font-size:6pt; fill:white; text-anchor:middle; alignment-baseline:middle")
	}
	canvas.Gend()

	canvas.End()
	return nil
}

func getDefaultColor(province *hoi3.Province) string {
	r, g, b, _ := province.Color.RGBA()
	r8 := uint8(r >> 8)
	g8 := uint8(g >> 8)
	b8 := uint8(b >> 8)
	return fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)
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
