package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Config 全局配置
var Config *GlobalConfig

func main() {
	Config = GetConfig(filepath.Join(".", "config.yml"), true)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		str, err := json.MarshalIndent(Config, "", "    ")
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		if err != nil {
			io.WriteString(w, err.Error())
		} else {
			fmt.Println(string(str))
			io.WriteString(w, fmt.Sprintf("<pre>%s</pre>", string(str)))
		}

	})
	http.ListenAndServe(":8800", nil)
}

// GetConfig 读取博客网站的yml配置
func GetConfig(configPath string, develop bool) *GlobalConfig {
	return ParseGlobalConfig(configPath, develop)
}

// ParseGlobalConfig 获取全局配置
func ParseGlobalConfig(configPath string, develop bool) *GlobalConfig {
	var config *GlobalConfig
	// 读取yml文件
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		Fatal(err.Error())
	}
	return config
}
