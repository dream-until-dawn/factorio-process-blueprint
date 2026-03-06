package utils

import (
	"strconv"
	"time"
)

func StringToFloat64(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return val
}

func StringToInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func TimestampToString(timestamp int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(timestamp, 0).In(loc).Format("2006-01-02")
}
