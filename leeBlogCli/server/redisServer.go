package server

import (
	"leeBlogCli/config"
	"leeBlogCli/definition"
)

func (b *Blog) GetRedisValueByKey(param *definition.GetRedisValueByKeyParam) string {
	if param.Key == config.RedisKeyImageTypeNeedConvert {
		return b.Dao.GetImageTypeNeedConvert()
	} else {
		return ""
	}
}
func (b *Blog) GetImageTypeNeedConvert() string {
	return b.Dao.GetImageTypeNeedConvert()
}

func (b *Blog) InitRedis() error {
	return b.Dao.InitRedis()
}
