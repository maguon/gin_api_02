package public

import (
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	req "gin_api_02/model/request"
	res "gin_api_02/model/response"
	"gin_api_02/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Login
// @Tags     User
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      req.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /public/login [post]
func (b *PublicApi) Login(c *gin.Context) {
	var l req.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &res.UserInfo{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.SYS_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

// Login
// @Tags     Admin
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      req.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.AdminLoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /public/adminLogin [post]
func (b *PublicApi) AdminLogin(c *gin.Context) {
	var l req.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &res.AdminInfo{Username: l.Username, Password: l.Password}
		user, err := adminService.AdminLogin(u)
		if err != nil {
			global.SYS_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		b.AdminTokenNext(c, *user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (b *PublicApi) TokenNext(c *gin.Context, user res.UserInfo) {
	j := &utils.JWT{SigningKey: []byte(global.SYS_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(req.BaseClaims{
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.SYS_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.SYS_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if _, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.SYS_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.SYS_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// TokenNext 登录以后签发jwt
func (b *PublicApi) AdminTokenNext(c *gin.Context, user res.AdminInfo) {
	j := &utils.JWT{SigningKey: []byte(global.SYS_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateAdminClaims(req.AdminBaseClaims{
		ID:       user.ID,
		Username: user.Username,
		Type:     user.Type,
	})
	token, err := j.CreateAdminToken(claims)
	if err != nil {
		global.SYS_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.SYS_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(res.AdminLoginResponse{
			Admin:     user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if _, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.SYS_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(res.AdminLoginResponse{
			Admin:     user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.SYS_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		response.OkWithDetailed(res.AdminLoginResponse{
			Admin:     user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
