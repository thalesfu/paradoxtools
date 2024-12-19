package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ajstarks/svgo"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/draw"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Region struct {
	ID      string
	Name    string
	Pixels  []Point
	Color   color.Color
	Polygon []Point
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

type Province struct {
	ID    string
	Color color.Color
	Name  string
}

// 从 CSV 读取颜色列表
func readProvincesFromCSV(csvPath string) ([]*Province, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var redIdx, greenIdx, blueIdx int = -1, -1, -1
	header := records[0]
	for i, col := range header {
		lc := strings.ToLower(col)
		if lc == "red" {
			redIdx = i
		} else if lc == "green" {
			greenIdx = i
		} else if lc == "blue" {
			blueIdx = i
		}
	}

	if redIdx == -1 || greenIdx == -1 || blueIdx == -1 {
		return nil, fmt.Errorf("CSV缺少red/green/blue列")
	}

	var colorsList []*Province
	for _, rec := range records[1:] {
		rVal, err1 := strconv.Atoi(rec[redIdx])
		gVal, err2 := strconv.Atoi(rec[greenIdx])
		bVal, err3 := strconv.Atoi(rec[blueIdx])
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		colorsList = append(colorsList, &Province{
			ID:    rec[0],
			Color: color.RGBA{R: uint8(rVal), G: uint8(gVal), B: uint8(bVal), A: 255},
			Name:  rec[4],
		})
	}

	return colorsList, nil
}

func colorsEqual(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

// 在图像中寻找指定颜色的所有像素点
func findAllPixelsOfColor(img image.Image, target color.Color) []Point {
	bounds := img.Bounds()
	var points []Point
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			if colorsEqual(c, target) {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
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

// 提取像素级边界(在8连通下，边界仍然同样处理，只是填充区域可能更完整)
func extractPixelBoundary(region Region, width, height int) []Point {
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

	var startPoint Point
	for k, v := range adj {
		if len(v) > 0 {
			startPoint = k
			break
		}
	}

	if (startPoint == Point{}) {
		return nil
	}

	polygon := []Point{startPoint}
	current := startPoint
	var prev Point

	for {
		nextList := adj[current]
		var next Point
		found := false
		for _, candidate := range nextList {
			if candidate != prev {
				next = candidate
				found = true
				break
			}
		}
		if !found {
			break
		}
		if next == startPoint && len(polygon) > 2 {
			break
		}
		polygon = append(polygon, next)
		prev = current
		current = next
		if len(polygon) > len(edges)+1 {
			break
		}
	}

	return polygon
}

// 生成 SVG
func generateSVG(allRegions []Region, width, height int, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")

	for i, region := range allRegions {
		if len(region.Polygon) < 3 {
			continue
		}
		r, g, b, _ := region.Color.RGBA()
		r8 := uint8(r >> 8)
		g8 := uint8(g >> 8)
		b8 := uint8(b >> 8)
		colorHex := fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)

		xs := make([]int, len(region.Polygon))
		ys := make([]int, len(region.Polygon))
		for i, p := range region.Polygon {
			xs[i] = p.X
			ys[i] = p.Y
		}

		canvas.Polygon(xs, ys, fmt.Sprintf("fill:%s;stroke:none", colorHex))
		fmt.Printf("绘制第%d个连通域,共有%d个连通域\n", i+1, len(allRegions))
	}

	canvas.End()
	return nil
}

// 将一个 image.Image 转为 RGBA 以方便操作
func toRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}

// 对图像进行180度旋转
func rotate180(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(bounds)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			srcX := w - x - 1
			srcY := h - y - 1
			dst.Set(x, y, img.At(srcX, srcY))
		}
	}
	return dst
}

// 水平翻转图像
func flipHorizontal(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(bounds)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			srcX := w - x - 1
			dst.Set(x, y, img.At(srcX, y))
		}
	}
	return dst
}

// 将图像先180度旋转，再水平镜像
func transformImage(img image.Image) *image.RGBA {
	rgba := toRGBA(img)
	r180 := rotate180(rgba)
	flipped := flipHorizontal(r180)
	return flipped
}

// 构建颜色到像素点的映射
func buildColorMap(img image.Image) map[color.Color][]Point {
	bounds := img.Bounds()
	colorMap := make(map[color.Color][]Point)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if x == 4728 && y == 874 {
				fmt.Println(x, y)
			}
			c := img.At(x, y)
			colorMap[c] = append(colorMap[c], Point{x, y})
		}
	}
	return colorMap
}

func main() {
	bmpPath := "/Users/thalesfu/Documents/My Games/HOI3/map.bmp"
	csvPath := "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3/tfh/map/definition.csv"
	outputSVG := "/Users/thalesfu/Documents/My Games/HOI3/output-03.svg"

	provinceList, err := readProvincesFromCSV(csvPath)
	if err != nil {
		fmt.Println("读取CSV错误:", err)
		return
	}

	bmpFile, err := os.Open(bmpPath)
	if err != nil {
		fmt.Println("无法打开BMP文件:", err)
		return
	}
	defer bmpFile.Close()

	img, err := bmp.Decode(bmpFile)
	if err != nil {
		fmt.Println("BMP解码错误:", err)
		return
	}

	//// 对图像先旋转180度，再水平翻转
	//img = transformImage(img)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 3. 构建颜色到像素点的映射
	colorMap := buildColorMap(img)

	var allRegions []Region
	for i, p := range provinceList {
		fmt.Printf("正在处理 %d/%d, %s\n", i+1, len(provinceList), p.Name)
		points := colorMap[p.Color]
		fmt.Printf("\t找到 %d 个像素点\n", len(points))
		if len(points) == 0 {
			continue
		}
		regs := findColorRegions(points)
		fmt.Printf("\t找到 %d 个连通域\n", len(regs))
		for i := range regs {
			regs[i].ID = p.ID
			regs[i].Name = p.Name
			regs[i].Color = p.Color
			regs[i].Polygon = extractPixelBoundary(regs[i], width, height)
			fmt.Printf("\t\t连通域 %d 的边界点数: %d\n", i+1, len(regs[i].Polygon))
			allRegions = append(allRegions, regs[i])
		}
	}

	err = generateSVG(allRegions, width, height, outputSVG)
	if err != nil {
		fmt.Println("生成SVG错误:", err)
		return
	}

	fmt.Println("SVG文件已生成:", outputSVG)
}
