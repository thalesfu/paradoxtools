package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/utils"
)

func printHex(str string) {
	for _, c := range []byte(str) {
		fmt.Printf("%02x ", c)
	}
	fmt.Println()
}

func printHexW(wideText []uint16) {
	for _, w := range wideText {
		fmt.Printf("%04x ", w)
	}
	fmt.Println()
}

func main() {
	content, ok := utils.LoadContent("T:\\codes\\github.com\\thalesfu\\cplusplus\\name2.txt")

	if ok {
		fmt.Println("paradox", []byte(content))
	}

	fileText, err := utils.DecodeEscapedText([]byte(content))

	fmt.Println("来自文件：", fileText)

	utf8Str := "斯来"
	fmt.Print("Original string in hex: ")
	printHex(utf8Str)
	wideStr, err := utils.MultiByteToWideChar(utf8Str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Go UTF-16 hex: ")
	printHexW(wideStr)

	text, err := utils.ConvertWideTextToEscapedText(wideStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Wide String:", text)
}
