package utils

import (
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"strings"
	"unicode/utf16"
)

func UCS2ToCP1252(cp uint16) uint16 {
	cp1252Map := map[uint16]uint16{
		0x20AC: 0x80, 0x201A: 0x82, 0x0192: 0x83, 0x201E: 0x84, 0x2026: 0x85,
		0x2020: 0x86, 0x2021: 0x87, 0x02C6: 0x88, 0x2030: 0x89, 0x0160: 0x8A,
		0x2039: 0x8B, 0x0152: 0x8C, 0x017D: 0x8E, 0x2018: 0x91, 0x2019: 0x92,
		0x201C: 0x93, 0x201D: 0x94, 0x2022: 0x95, 0x2013: 0x96, 0x2014: 0x97,
		0x02DC: 0x98, 0x2122: 0x99, 0x0161: 0x9A, 0x203A: 0x9B, 0x0153: 0x9C,
		0x017E: 0x9E, 0x0178: 0x9F,
	}

	if val, ok := cp1252Map[cp]; ok {
		return val
	}
	return cp
}

func CP1252ToUCS2(cp uint16) uint16 {
	// 反向映射表
	ucs2Map := map[uint16]uint16{
		0x80: 0x20AC, 0x82: 0x201A, 0x83: 0x0192, 0x84: 0x201E, 0x85: 0x2026,
		0x86: 0x2020, 0x87: 0x2021, 0x88: 0x02C6, 0x89: 0x2030, 0x8A: 0x0160,
		0x8B: 0x2039, 0x8C: 0x0152, 0x8E: 0x017D, 0x91: 0x2018, 0x92: 0x2019,
		0x93: 0x201C, 0x94: 0x201D, 0x95: 0x2022, 0x96: 0x2013, 0x97: 0x2014,
		0x98: 0x02DC, 0x99: 0x2122, 0x9A: 0x0161, 0x9B: 0x203A, 0x9C: 0x0153,
		0x9E: 0x017E, 0x9F: 0x0178,
	}

	// 进行反向映射
	if val, ok := ucs2Map[cp]; ok {
		return val
	}
	return cp
}

func ConvertWideTextToEscapedText(from []uint16) ([]byte, error) {
	if from == nil {
		return nil, fmt.Errorf("input string is nil")
	}

	size := len(from)
	to := make([]byte, 0, size*6+10)

	for _, cp := range from {
		if cp == 0 {
			break
		}
		originalCP := cp
		cp = UCS2ToCP1252(cp)

		if originalCP != cp {
			to = append(to, byte(cp))
			continue
		}

		if cp > 0x100 && cp < 0xA00 {
			cp += 0xE000
		}

		high := byte((cp >> 8) & 0xFF)
		low := byte(cp & 0xFF)

		escapeChr := byte(0x10)

		if high != 0 {
			switch high {
			case 0xA4, 0xA3, 0xA7, 0x24, 0x5B, 0x00, 0x5C, 0x20, 0x0D, 0x0A,
				0x22, 0x7B, 0x7D, 0x40, 0x80, 0x7E, 0x2F, 0xBD, 0x3B, 0x5D,
				0x5F, 0x3D, 0x23:

				escapeChr += 2
			}

			switch low {
			case 0xA4, 0xA3, 0xA7, 0x24, 0x5B, 0x00, 0x5C, 0x20, 0x0D, 0x0A,
				0x22, 0x7B, 0x7D, 0x40, 0x80, 0x7E, 0x2F, 0xBD, 0x3B, 0x5D,
				0x5F, 0x3D, 0x23:

				escapeChr++
			}

			switch escapeChr {
			case 0x11:
				low += 15
			case 0x12:
				high -= 9
			case 0x13:
				low += 15
				high -= 9
			}

			to = append(to, escapeChr, low, high)
		} else {
			to = append(to, byte(cp))
		}
	}

	return to, nil
}

func ConvertEscapedTextToWideText(escapedText []byte) ([]uint16, error) {
	var result []uint16
	var i int
	var cp uint16

	for i < len(escapedText) {
		firstPart := escapedText[i]

		if firstPart == 0x10 || firstPart == 0x11 || firstPart == 0x12 || firstPart == 0x13 {
			escapeChr := firstPart
			i++
			low := escapedText[i]
			i++
			high := escapedText[i]

			switch escapeChr {
			case 0x11:
				low -= 15
			case 0x12:
				high += 9
			case 0x13:
				low -= 15
				high += 9
			}

			cp = combineHighLow(high, low)
		} else {
			cp = uint16(firstPart)
		}

		if cp >= 0xE100 && cp < 0xEA00 {
			cp -= 0xE000
		}
		cp = CP1252ToUCS2(cp) // 使用逆映射函数
		result = append(result, cp)

		i++
	}

	if result[len(result)-1] != 0 {
		result = append(result, 0)
	}

	return result, nil
}

func combineHighLow(high, low byte) uint16 {
	return uint16(high)<<8 | uint16(low)
}

func DecodeEscapedText(escapedText []byte) (string, error) {
	text, err := ConvertEscapedTextToWideText(escapedText)

	if err != nil {
		return string(escapedText), err
	}

	utf8Str, err := WideCharToMultiByte(text)

	if err != nil {
		return string(escapedText), err
	}

	utf8Str = strings.TrimRight(utf8Str, "\u0000")

	return utf8Str, nil
}

func EncodeEscapedText(text string) ([]byte, error) {
	wideChar, err := MultiByteToWideChar(text)

	if err != nil {
		return nil, err
	}

	escapedTextBytes, err := ConvertWideTextToEscapedText(wideChar)

	if err != nil {
		return nil, err
	}

	return escapedTextBytes, nil
}

// MultiByteToWideChar converts a string from UTF-8 to UTF-16.
func MultiByteToWideChar(utf8Str string) ([]uint16, error) {
	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16Bytes, _, err := transform.Bytes(encoder, []byte(utf8Str))
	if err != nil {
		return nil, fmt.Errorf("failed to encode UTF-8 to UTF-16: %v", err)
	}

	// Convert byte slice to uint16 slice
	utf16Ints := make([]uint16, 0, len(utf16Bytes)/2)
	for i := 0; i < len(utf16Bytes); i += 2 {
		utf16Ints = append(utf16Ints, uint16(utf16Bytes[i])+(uint16(utf16Bytes[i+1])<<8))
	}

	return utf16Ints, nil
}

// WideCharToMultiByte converts a string from UTF-16 to UTF-8.
func WideCharToMultiByte(wideStr []uint16) (string, error) {
	// Decode UTF-16 to UTF-8
	runes := utf16.Decode(wideStr)
	utf8Str := string(runes)

	return utf8Str, nil
}
