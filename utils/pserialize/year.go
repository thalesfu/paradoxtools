package pserialize

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Year time.Time

func (y *Year) UnmarshalP(data []byte) error {
	s, ok := unquoteBytes(data)
	if !ok {
		log.Fatalf("pserialize: invalid year %q", data)
	}
	dateString := string(s)

	// 将字符串分割为年、月、日
	parts := strings.Split(dateString, ".")
	if len(parts) != 3 {
		panic("Invalid date format")
	}

	// 分别解析年、月、日
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	// 创建 time.Time 对象
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	*y = Year(date)
	return nil
}

func (y Year) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", y.SimpleDateStr())), nil
}

func (y Year) SimpleDateStr() string {
	return time.Time(y).Format("2006-01-02")
}

func (y Year) PDateStr() string {
	t := time.Time(y)
	year, month, day := t.Date()
	return fmt.Sprintf("%d.%d.%d", year, month, day)
}
