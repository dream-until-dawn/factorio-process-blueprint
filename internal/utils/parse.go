package utils

import (
	"fmt"
	"strings"
)

func FlattenToString(v interface{}) string {
	var result []string
	fmt.Printf("%+v\n", v)
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
