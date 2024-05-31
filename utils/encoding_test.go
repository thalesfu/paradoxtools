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
			data, err := EncodeEscapedText("彦伟")

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data, string(data))

			WriteContent("temps/escapedtext.txt", string(data))
		})

		Convey("解码测试", func() {
			//data, err := EncodeEscapedText("可生")
			data := []byte{16, 133, 80, 16, 239, 83, 16, 31, 117, 16, 199, 150, 16, 99, 79, 16, 155, 81}
			s, err := DecodeEscapedText(data)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data, s)

			content, ok := LoadContent("temps/escapedtext.txt")

			if ok {

				text, err := DecodeEscapedText([]byte(content))

				if err == nil {
					fmt.Println(text)
				}

			}

		})
	})
}
