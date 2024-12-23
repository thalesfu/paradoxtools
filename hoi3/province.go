package hoi3

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/thalesfu/paradoxtools/hoi3/save"
	"github.com/thalesfu/paradoxtools/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"image"
	"image/color"
	"path"
	"strconv"
	"strings"
)

type Province struct {
	ID             string
	Color          color.Color
	Polygons       [][]*Point
	Name           string
	Terrain        string
	Owner          string
	Controller     string
	Core           map[string]bool
	Point          int
	Manpower       float32
	Leadership     float32
	Energy         float32
	Metal          float32
	RareMaterials  float32
	CrudeOil       float32
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
	Supplies       float32
	Fuel           float32
}

type Region struct {
	Pixels   []Point
	Color    color.Color
	Polygons [][]Point
}

// 八连通方向，包括水平、垂直、对角线
var directions = []Point{
	{0, -1},  // 上
	{-1, 0},  // 左
	{1, 0},   // 右
	{0, 1},   // 下
	{-1, -1}, // 左上
	{1, -1},  // 右上
	{-1, 1},  // 左下
	{1, 1},   // 右下
}

func LoadProvinces(fileLocation *FileLocation) (map[string]*Province, int, int) {
	provinces := loadProvincesFromMap(fileLocation)
	fmt.Println("加载省份定义完成，共", len(provinces), "个省份")
	provinces, width, height := loadProvincePolygons(fileLocation, provinces)
	fmt.Println("加载省份边界完成")

	provinceNames := loadProvincesName(fileLocation)
	for _, province := range provinces {
		province.Name = provinceNames[province.ID]
	}
	fmt.Println("加载省份名称完成")

	LoadProvincesFromSave(fileLocation, provinces)
	fmt.Println("加载省份数据完成")

	return provinces, width, height
}

func loadProvincesFromMap(fileLocation *FileLocation) map[string]*Province {
	const MapDefinePath = "map/definition.csv"
	content, ok := utils.LoadContent(path.Join(fileLocation.ModDirectory, MapDefinePath))
	if ok {
		return generateProvinceFromMapDefineContent(content)
	}

	content, ok = utils.LoadContent(path.Join(fileLocation.BaseDirectory, MapDefinePath))
	if ok {
		return generateProvinceFromMapDefineContent(content)
	}

	return make(map[string]*Province)
}

func generateProvinceFromMapDefineContent(content string) map[string]*Province {
	provinces := make(map[string]*Province)

	lines := strings.Split(content, "\r\n")
	if len(lines) == 1 {
		lines = strings.Split(content, "\n")
	}

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
			ID:    fields[0],
			Color: color.RGBA{R: uint8(rVal), G: uint8(gVal), B: uint8(bVal), A: 255},
		}
	}

	return provinces
}

func loadProvincesName(fileLocation *FileLocation) map[string]string {
	const ProvinceNamesPath = "localisation/province_names.csv"
	content, ok := utils.LoadContentWithEncoding(path.Join(fileLocation.ModDirectory, ProvinceNamesPath), simplifiedchinese.GB18030)
	if ok {
		return generateProvinceNameFromContent(content)
	}

	content, ok = utils.LoadContentWithEncoding(path.Join(fileLocation.BaseDirectory, ProvinceNamesPath), simplifiedchinese.GB18030)
	if ok {
		return generateProvinceNameFromContent(content)
	}

	return make(map[string]string)
}

func generateProvinceNameFromContent(content string) map[string]string {
	provinceNames := make(map[string]string)

	lines := strings.Split(content, "\r\n")
	if len(lines) == 1 {
		lines = strings.Split(content, "\n")
	}

	for _, line := range lines[1:] {
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			continue
		}

		if strings.HasPrefix(fields[0], "PROV") {
			fields[0] = fields[0][4:]
			provinceNames[fields[0]] = fields[1]
		}
	}

	return provinceNames
}

func LoadProvincesFromSave(fileLocation *FileLocation, provinces map[string]*Province) {
	saveFile, _, err := save.LoadSave(fileLocation.SaveFile)
	if err != nil {
		return
	}

	for _, province := range provinces {
		if sp, ok := saveFile.Province[province.ID]; ok {
			if len(sp.Infra) > 0 {
				province.Infra = int(sp.Infra[1])
			}
		}
	}
}

func loadProvincePolygons(fileLocation *FileLocation, provinces map[string]*Province) (map[string]*Province, int, int) {
	mapBmp := loadProvinceMap(fileLocation)
	if mapBmp == nil {
		return provinces, 0, 0
	}

	mapBmp = utils.VerticalFlip(mapBmp)

	bounds := mapBmp.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 构建颜色到像素点的映射
	colorMap := buildColorMap(mapBmp)

	provinceIndex := 0
	for _, p := range provinces {
		provinceIndex++
		points := colorMap[p.Color]
		if len(points) == 0 {
			continue
		}

		//调试断点
		if p.ID == "11585" {
			fmt.Println("Debug", p.ID, p.Name)
		}

		regs := findColorRegions(points)
		for i := range regs {
			regs[i].Color = p.Color
			p.Polygons = append(p.Polygons, extractPixelBoundary(regs[i], width, height)...)
		}

		fmt.Printf("加载省份边界 省份ID:%s 边界数量 %d  %d-%d \n", p.ID, len(p.Polygons), provinceIndex, len(provinces))
	}

	return provinces, width, height
}

func loadProvinceMap(fileLocation *FileLocation) image.Image {
	const ProvinceMapPath = "map/provinces.bmp"
	bmp, ok := utils.LoadBmpFile(path.Join(fileLocation.ModDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	bmp, ok = utils.LoadBmpFile(path.Join(fileLocation.BaseDirectory, ProvinceMapPath))
	if ok {
		return bmp
	}

	return nil
}

// 构建颜色到像素点的映射
func buildColorMap(img image.Image) map[color.Color][]Point {
	bounds := img.Bounds()
	colorMap := make(map[color.Color][]Point)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//if x == 4728 && y == 874 {
			//	fmt.Println(x, y)
			//}
			c := img.At(x, y)
			colorMap[c] = append(colorMap[c], Point{x, y})
		}
	}
	return colorMap
}

// 根据给定的一组同色像素点划分连通域(8连通)
func findColorRegions(points []Point) []Region {
	if len(points) == 0 {
		return nil
	}

	visited := make(map[Point]bool, len(points))
	pointSet := make(map[Point]bool, len(points))
	for _, p := range points {
		pointSet[p] = true
	}

	var regions []Region

	for _, start := range points {
		if visited[start] {
			continue
		}
		queue := []Point{start}
		visited[start] = true
		var regionPixels []Point
		regionPixels = append(regionPixels, start)

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				np := Point{cur.X + d.X, cur.Y + d.Y}
				if pointSet[np] && !visited[np] {
					visited[np] = true
					queue = append(queue, np)
					regionPixels = append(regionPixels, np)
				}
			}
		}

		regions = append(regions, Region{Pixels: regionPixels})
	}

	return regions
}

// 提取像素级边界(在8连通下，边界仍然同样处理，只是填充区域可能更完整)
func extractPixelBoundary(region Region, width, height int) [][]*Point {
	filled := buildFilledMap(region, width, height)

	type Edge struct {
		Start, End Point
	}

	var edges []Edge

	for _, p := range region.Pixels {
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

func buildFilledMap(region Region, width, height int) [][]bool {
	filled := make([][]bool, height)
	for i := range filled {
		filled[i] = make([]bool, width)
	}
	for _, p := range region.Pixels {
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
