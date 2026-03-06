package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SaveToJSON 将任意数据保存为 JSON 文件
func SaveToJSON(filename string, data interface{}, pretty bool) error {
	var jsonData []byte
	var err error

	if pretty {
		// 带缩进的格式化 JSON
		jsonData, err = json.MarshalIndent(data, "", "  ")
	} else {
		// 紧凑格式（无额外空格）
		jsonData, err = json.Marshal(data)
	}

	if err != nil {
		return fmt.Errorf("JSON 编码失败: %w", err)
	}

	// 确保目录存在
	dir := filepath.Dir(filename)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录失败: %w", err)
		}
	}

	// 写入文件
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}
