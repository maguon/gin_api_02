package service

import (
	"context"
	"gin_api_02/global"
	"gin_api_02/utils"
)

type JwtService struct{}

//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.SYS_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.SYS_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
