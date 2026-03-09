package utils

import (
	rawDataModels "processBlueprint/internal/model/rawData"
	"strings"
)

func FlattenToString(v interface{}) string {
	var result []string
	flatten(v, &result)
	return strings.Join(result, ",")
}

func flatten(v interface{}, result *[]string) {
	switch val := v.(type) {
	case string:
		*result = append(*result, val)

	case []interface{}:
		for _, item := range val {
			flatten(item, result)
		}
	}
}

// 输出配置允许的产所
// 0-太空 pressure=0 gravity=0 magnetic-field=0
// 1-新地星 pressure=1000 gravity=10 magnetic-field=90
// 2-祝融星 pressure=4000 gravity=40 magnetic-field=25
// 3-雷神星 pressure=800 gravity=8 magnetic-field=99
// 4-句芒星 pressure=2000 gravity=20 magnetic-field=10
// 5-玄冥星 pressure=300 gravity=15 magnetic-field=10
// 如何任意地点可制作,则为 111111,仅太空则为 100000
// 六个地点属性
var surfaces = []map[string]float64{
	{"pressure": 0, "gravity": 0, "magnetic-field": 0},      // 0 太空
	{"pressure": 1000, "gravity": 10, "magnetic-field": 90}, // 1 新地星
	{"pressure": 4000, "gravity": 40, "magnetic-field": 25}, // 2 祝融星
	{"pressure": 800, "gravity": 8, "magnetic-field": 99},   // 3 雷神星
	{"pressure": 2000, "gravity": 20, "magnetic-field": 10}, // 4 句芒星
	{"pressure": 300, "gravity": 15, "magnetic-field": 10},  // 5 玄冥星
}

func GetAllowedSurface(surface []rawDataModels.SurfaceConditions) int {
	// 没条件 = 全允许
	if len(surface) == 0 {
		return SurfaceMask("111111")
	}

	result := make([]byte, 6)

	for i, s := range surfaces {
		allowed := true

		for _, cond := range surface {
			val, ok := s[cond.Property]
			if !ok {
				allowed = false
				break
			}

			if cond.Min != nil && val < *cond.Min {
				allowed = false
				break
			}

			if cond.Max != nil && val > *cond.Max {
				allowed = false
				break
			}
		}

		if allowed {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}

	return SurfaceMask(string(result))
}
