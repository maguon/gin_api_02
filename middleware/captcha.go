package middleware

import (
	"gin_api_02/model/common/response"
	req "gin_api_02/model/request"
	redis_util "gin_api_02/utils"

	"github.com/gin-gonic/gin"
)

var store = redis_util.NewCaptchaRedisStore()

func CheckCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l req.ReqCaptcha
		err := c.ShouldBindJSON(&l)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "Captcha is Error", c)
			c.Abort()
			return
		}
		if store.Verify(l.CaptchaId, l.Captcha, true) {
			c.Next()
		} else {
			c.Abort()
			return
		}
	}
}
