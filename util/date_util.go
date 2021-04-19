package util

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func GetYearMonth() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	date, err := strconv.Atoi(fmt.Sprintf("%d%d", year, month))
	if err != nil {
		log.Printf("format error")
	}
	return int64(date)
}

func GetMonth() int64 {
	return int64(time.Now().Month())
}

func GetDay() int64 {
	return int64(time.Now().Day())
}

func YearMonthDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
