package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var globalConfig *Config

// ================================
// Config Struct
// ================================
type Config struct {
	Environment string `yaml:"environment"`
	// log config
	LogLevel string `yaml:"log_level"`
}

// ================================
// Load Config
// ================================
func Load() {
	abs := ""
	data := []byte{}
	abs1, err1 := filepath.Abs("env/config.yaml")
	if err1 != nil {
		log.Fatalf("\033[31m[error & exit] \033[0m无法解析配置文件路径: %v", err1)
	}

	abs2, err1 := filepath.Abs("configs/config.yaml")
	if err1 != nil {
		log.Fatalf("\033[31m[error & exit] \033[0m无法解析配置文件路径: %v", err1)
	}

	data1, err := ioutil.ReadFile(abs1)
	if err != nil {
		data2, err := ioutil.ReadFile(abs2)
		if err != nil {
			log.Fatalf("\033[31m[error & exit] \033[0m配置文件 %s 不存在: %v", abs2, err)
		}
		abs = abs2
		data = data2
	} else {
		abs = abs1
		data = data1
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalf("\033[31m[error & exit] \033[0m配置文件解析失败: %v", err)
	}

	globalConfig = cfg
	log.Printf("\033[32m[suc] \033[0m配置加载完成: %s", abs)
}

// ================================
// Get Global Config
// ================================
func Get() *Config {
	if globalConfig == nil {
		Load()
	}
	return globalConfig
}
