package utils

import (
	"fmt"
	"strings"
)

func ReplaceTemplateSpecialWords(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ReplaceAll(s, "－", "_")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "[", "_")
	s = strings.ReplaceAll(s, "]", "_")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "’", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, "“", "")
	s = strings.ReplaceAll(s, "”", "")
	s = strings.ReplaceAll(s, "‘", "")
	s = strings.ReplaceAll(s, "’", "")
	return s
}

func EscapeTemplateSpecialWords(s string) string {
	if s == "" {
		return s
	}

	text, err := DecodeEscapedText([]byte(s))
	if err != nil {
		fmt.Println(err)
		return s
	}
	return text
}
