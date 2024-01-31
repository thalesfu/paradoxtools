package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCK2EscapedText(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Ck2 Escaped", t, func() {

		Convey("编码测试", func() {
			data, err := EncodeEscapedText("武毕楼储殷张袁林")

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data, string(data))

			WriteContent("temps/escapedtext.txt", string(data))
		})

		Convey("解码测试", func() {
			data, err := EncodeEscapedText("可生")
			s, err := DecodeEscapedText(data)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data, s)
		})
	})
}
