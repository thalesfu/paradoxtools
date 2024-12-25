package hoi3

import (
	"github.com/thalesfu/paradoxtools/utils"
	"path"
	"regexp"
)

type Region struct {
	ID        string
	Name      string
	Provinces []string
}

var regionRegex = regexp.MustCompile(`(?P<id>\w+)\s*=\s*\{\s*(?P<value>((\d*)\s*)*)\}`)
var splitRegex = regexp.MustCompile(`\s+`)
var numRegex = regexp.MustCompile(`\d+`)

func LoadRegion(fileLocation *FileLocation, localisation map[string]string) map[string]*Region {
	const regionPath = "map/region.txt"

	regions := make(map[string]*Region)

	LoadRegionsFromFile(path.Join(fileLocation.BaseDirectory, regionPath), regions)
	LoadRegionsFromFile(path.Join(fileLocation.DLCDirectory, regionPath), regions)
	LoadRegionsFromFile(path.Join(fileLocation.ModDirectory, regionPath), regions)

	for _, region := range regions {
		region.Name = localisation[region.ID]
	}

	return regions
}

func LoadRegionsFromFile(regionPath string, regions map[string]*Region) {
	content, ok := utils.LoadContent(regionPath)
	if ok {
		// 查找匹配项
		matches := regionRegex.FindAllStringSubmatch(content, -1)

		for _, match := range matches {
			if len(match) > 0 {
				// 提取命名组
				result := make(map[string]string)
				for i, name := range regionRegex.SubexpNames() {
					if i > 0 && name != "" {
						result[name] = match[i]
					}
				}

				region := &Region{
					ID: result["id"],
				}

				regions[region.ID] = region

				// 处理 value（转换为切片）
				value := result["value"]
				if value != "" {
					values := splitRegex.Split(value, -1)
					for _, v := range values {
						if numRegex.MatchString(v) {
							region.Provinces = append(region.Provinces, v)
						}
					}
				}
			}
		}
	}
}
