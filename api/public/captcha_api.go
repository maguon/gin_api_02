package public

import (
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	systemRes "gin_api_02/model/response"
	redis_util "gin_api_02/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = redis_util.NewCaptchaRedisStore()

type PublicApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router    /public/captcha [get]
func (b *PublicApi) Captcha(ctx *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.SYS_CONFIG.Captcha.ImgHeight, global.SYS_CONFIG.Captcha.ImgWidth, global.SYS_CONFIG.Captcha.KeyLong, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx)) // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx.Request.Context()))
	id, b64s, err := cp.Generate()
	if err != nil {
		global.SYS_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", ctx)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId: id,
		Img:       b64s,
	}, "验证码获取成功", ctx)
}
