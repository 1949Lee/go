package main

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	GetConfig(filepath.Join(".", "config.yml"), true)
}

// GetConfig 读取博客网站的yml配置
func GetConfig(configPath string, develop bool) {
	ParseGlobalConfig(configPath, develop)
}

// ParseGlobalConfig 获取全局配置
func ParseGlobalConfig(configPath string, develop bool) *GlobalConfig {
	var config *GlobalConfig
	// Parse Global Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		Fatal(err.Error())
	}
	return config
}
