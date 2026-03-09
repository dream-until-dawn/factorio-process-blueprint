package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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

// 特供: 编译正则表达式：匹配最后一个 / 和 .png 之间的内容
func ExtractNameFromPath(path string) string {
	re := regexp.MustCompile(`/([^/]+)\.png$`)
	matches := re.FindStringSubmatch(path)

	if len(matches) > 1 {
		return strings.ReplaceAll(Capitalize(matches[1]), "-", "_")
	}
	return ""
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	// 将字符串转换为 rune 切片以正确处理 Unicode
	r := []rune(s)
	// 将第一个字符转换为大写
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
