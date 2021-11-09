package dao

import (
	"context"
	"leeBlogCli/config"
	"log"
)

var ctx = context.Background()

// GetImageTypeNeedConvert 获取需要渐进式转化的图片的类型
func (s *DBServer) GetImageTypeNeedConvert() string {
	if config.ENV == "dev" {
		return config.ImageTypeNeedConvert
	} else {
		val, err := s.Redis.Get(ctx, config.RedisKeyImageTypeNeedConvert).Result()
		if err != nil {
			log.Printf("dao.GetImageTypeNeedConvert 获取需要渐进式转化的图片的类型错误 error：%v", err)
			return ""
		}
		return val
	}
}

func (s *DBServer) SetImageTypeNeedConvert() error {
	if config.ENV != "dev" {
		err := s.Redis.Set(ctx, config.RedisKeyImageTypeNeedConvert, config.ImageTypeNeedConvert, 0).Err()
		if err != nil {
			log.Printf("dao.GetImageTypeNeedConvert 获取需要渐进式转化的图片的类型错误 error：%v", err)
			return err
		}
		return nil
	} else {
		return nil
	}
}

func (s *DBServer) InitRedis() error {
	if err := s.SetImageTypeNeedConvert(); err != nil {
		return err
	}
	return nil
}
