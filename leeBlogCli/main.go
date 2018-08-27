package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/russross/blackfriday.v2"

	yaml "gopkg.in/yaml.v2"
)

// Config 全局配置
var Config *GlobalConfig
var RootDir string

func main() {

	RootDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // 得到go程序入口的对应路径
	if err != nil {
		log.Fatal(err)
	}
	Config = GetConfig(filepath.Join(RootDir, "config.yml"), true)
	http.HandleFunc("/yml", func(w http.ResponseWriter, req *http.Request) {
		// str, err := json.MarshalIndent(Config, "", "    ")
		// w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// if err != nil {
		// 	// err.Error()
		// 	io.WriteString(w, "未找到对应的YAML文件")
		// } else {
		// 	fmt.Println(string(str))
		// 	io.WriteString(w, fmt.Sprintf("<pre>%s</pre>", string(str)))
		// }
		str := ParseMarkDown("ink-blog-tool.md")
		// html.EscapeString(string(str)) // 将HTML转换义方便存到数据库
		io.WriteString(w, fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<code>%s</code>
</body>
</html>
		        `, html.UnescapeString(string(str)))) // 将HTML字符串反转义为HTML代码，渲染到网页
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

// ParseMarkDown MarkDown转换成HTML
func ParseMarkDown(fileName string) []byte {
	data, err := ioutil.ReadFile(filepath.Join(RootDir, Config.Site.Theme, fileName))
	if err != nil {
		Fatal(err.Error())
	}
	// Split config and markdown
	contentStr := string(data)
	markdownStr := strings.SplitN(contentStr, "---", 2)
	out := blackfriday.Run([]byte(markdownStr[1]))
	return out
}
