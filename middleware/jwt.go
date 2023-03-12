package middleware

import (
	"fmt"
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	"gin_api_02/service"
	"gin_api_02/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("auth-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.SYS_CONFIG.System.UseMultipoint {
				if err != nil {
					global.SYS_LOG.Error("get redis jwt failed", zap.Error(err))
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func JWTAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("auth-token")
		fmt.Println(token)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseAdminToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateAdminToken(*claims)
			newClaims, _ := j.ParseAdminToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.SYS_CONFIG.System.UseMultipoint {
				if err != nil {
					global.SYS_LOG.Error("get redis jwt failed", zap.Error(err))
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
