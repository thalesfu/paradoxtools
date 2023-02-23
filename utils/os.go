package utils

import (
	"path/filepath"
	"strings"
)

func GetFileNameWithoutExtension(p string) string {
	fileName := filepath.Base(p)
	fileExt := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, fileExt)
}
