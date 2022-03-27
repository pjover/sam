package model

import (
	"fmt"
	"time"
)

const YearMonthLayout = "2006-01"

type YearMonth struct {
	year  int
	month time.Month
}

func (ym YearMonth) Year() int {
	return ym.year
}

func (ym YearMonth) Month() time.Month {
	return ym.month
}

func NewYearMonth(year int, month time.Month) YearMonth {
	return YearMonth{
		year:  year,
		month: month,
	}
}

func StringToYearMonth(yearMonth string) (YearMonth, error) {
	tm, err := time.Parse(YearMonthLayout, yearMonth)
	if err != nil {
		return YearMonth{}, fmt.Errorf("error al introduïr el mes i any: %s", err.Error())
	}
	return YearMonth{tm.Year(), tm.Month()}, nil
}

func TimeToYearMonth(tm time.Time) YearMonth {
	return YearMonth{tm.Year(), tm.Month()}
}

func (ym YearMonth) String() string {
	return fmt.Sprintf("%d-%02d", ym.year, ym.month)
}
