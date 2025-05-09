package hoi3

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetLocalisation(location *FileLocation) map[string]string {
	const localisationPath = "localisation"
	localisation := make(map[string]string)
	LoadLocalisationFromDirect(filepath.Join(location.BaseDirectory, localisationPath), localisation)
	LoadLocalisationFromDirect(filepath.Join(location.DLCDirectory, localisationPath), localisation)
	LoadLocalisationFromDirect(filepath.Join(location.ModDirectory, localisationPath), localisation)

	return localisation
}

func LoadLocalisationFromDirect(localisationPath string, localisation map[string]string) {
	err := filepath.Walk(localisationPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("无法访问路径 %s: %v\n", path, err)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".csv" {
			return nil
		}

		content, ok := golangutils.LoadContentWithEncoding(path, simplifiedchinese.GB18030)
		if ok {
			loadLocalisationFromContent(content, localisation)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
	}
}

var localisationNewLineRegex = regexp.MustCompile(`\r\n?|\n`)

func loadLocalisationFromContent(content string, localisation map[string]string) {
	lines := localisationNewLineRegex.Split(content, -1)

	for _, line := range lines[1:] {
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			continue
		}

		localisation[fields[0]] = fields[1]
	}
}
