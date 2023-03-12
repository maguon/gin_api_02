package utils

import (
	"fmt"
	"gin_api_02/global"
	req "gin_api_02/model/request"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*req.CustomClaims, error) {
	token := c.Request.Header.Get("auth-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.SYS_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetAdminClaims(c *gin.Context) (*req.AdminCustomClaims, error) {
	token := c.Request.Header.Get("auth-token")
	fmt.Println(token)
	j := NewJWT()
	claims, err := j.ParseAdminToken(token)
	if err != nil {
		global.SYS_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*req.CustomClaims)
		return waitUse.ID
	}
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetAdminID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetAdminClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*req.AdminCustomClaims)
		return waitUse.ID
	}
}
