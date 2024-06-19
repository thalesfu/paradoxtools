package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConvertWIKIToMarkdown(t *testing.T) {
	Convey("ConvertWIKIToMarkdown", t, func() {
		Convey("Test ConvertWIKIToMarkdown", func() {
			ConvertWIKIToMarkdown("temps/wikimarkdownsource.md", "temps/wikimarkdowntarget.md")
		})
	})
}
