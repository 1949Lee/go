/**
定义了用到的结构体
*/
package main

// SiteConfig 站点配置
type SiteConfig struct {
	Title    string
	Subtitle string
	Theme    string
	Lang     string
	URL      string
}

// Author 博客作者
type Author struct {
	ID    int
	Name  string `yaml:"name"` // 解析的tag，不加这个不区分大小写，虽然yaml区分大小写。
	Intro string
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	Site    SiteConfig        // 站点配置
	Authors map[string]Author // 作者们
}
