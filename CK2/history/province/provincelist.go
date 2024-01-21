package province

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/ck2utils"
	"github.com/thalesfu/paradoxtools/CK2/config"
	"github.com/thalesfu/paradoxtools/CK2/translations"
	"github.com/thalesfu/paradoxtools/segments"
	"github.com/thalesfu/paradoxtools/utils"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var ProvinceCodeList = make(map[string]Province)
var ProvinceIdList = make(map[int]Province)

func init() {
	err := filepath.Walk(config.CK2ProvinceFold, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if !strings.HasSuffix(p, ".txt") {
			return nil
		}

		id, titleCode, ok := getIDAndTitle(p, segments.LoadSegments(p))

		if ok {
			p := &BaseProvince{
				ID:        id,
				Title:     translations.GetFeudName(titleCode),
				TitleCode: titleCode,
			}
			ProvinceCodeList[titleCode] = p
			i, err := strconv.Atoi(id)
			if err == nil {
				ProvinceIdList[i] = p
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("error:", err)
	}
}

var rgxProvince = regexp.MustCompile(`^(\d+)\s*-\s*.*$`)

func getIDAndTitle(p string, segments []*segments.Segment) (string, string, bool) {

	var id, titleCode string
	var ok bool

	fileName := utils.GetFileNameWithoutExtension(p)

	matches := rgxProvince.FindStringSubmatch(fileName)
	if len(matches) >= 2 {
		id = matches[1]
	}

	for _, segment := range segments {
		if segment.Name == "title" {
			if ck2utils.IsCountyString(segment.Value) {
				titleCode = segment.Value
				ok = true
			}
		}
	}

	return id, titleCode, ok
}
