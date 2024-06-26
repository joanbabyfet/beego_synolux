// 父控制器
package controllers

import (
	"encoding/base64"
	"strings"
	dto "synolux/dto"
	"synolux/models"
	"synolux/service"
	"synolux/utils"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	BaseController
}

// 登录
func (c *UserController) Login() {
	username := c.GetString("username") //帐号
	password := c.GetString("password") //密码
	code := c.GetString("code")         //验证码
	key := c.GetString("key")           //验证码key
	login := dto.UserLogin{}            //登录请求格式
	login.Username = username
	login.Password = password
	login.Code = code
	login.Key = key
	login.LoginIp = c.getClientIp()
	enable_captcha, _ := beego.AppConfig.Bool("enable_captcha") //是否启用验证码

	//用户密码解密
	pwd, _ := base64.StdEncoding.DecodeString(login.Password)
	login.Password = string(pwd)

	//参数验证
	valid := validation.Validation{}
	valid.Required(login.Username, "username")
	valid.Required(login.Password, "password")
	if enable_captcha {
		valid.Required(login.Key, "key")
		valid.Required(login.Code, "code")
	}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	//检测验证码
	// if enable_captcha && !cpt.Verify(key, code) {
	// 	c.ErrorJson(-2, "验证码错误", nil)
	// }
	if enable_captcha && !utils.Store.Verify(key, code, true) {
		c.ErrorJson(-2, "验证码错误", nil)
	}

	//登录
	service_user := new(service.UserService)
	stat, user, err := service_user.Login(login)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", user)
}

// 登录退出
func (c *UserController) Logout() {
	//c.DestroySession()
	c.SuccessJson("success", nil)
}

// 修改密码
func (c *UserController) SetPassword() {
	password := c.GetString("password")         //原始密码
	new_password := c.GetString("new_password") //新密码
	re_password := c.GetString("re_password")   //确认密码
	auth := c.Ctx.Input.Header("Authorization")
	kv := strings.Split(auth, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		c.ErrorJson(-1, "未带token", nil)
	}
	token := kv[1]
	payload, err := models.ValidateToken(token)
	if err != nil {
		c.ErrorJson(-2, "未登录或登录超时", nil)
	}
	uid := payload.UserID //取得用户id

	//用户密码解密
	pwd, _ := base64.StdEncoding.DecodeString(password)
	new_pwd, _ := base64.StdEncoding.DecodeString(new_password)
	re_pwd, _ := base64.StdEncoding.DecodeString(re_password)

	dto := dto.Password{}
	dto.Password = string(pwd)
	dto.NewPassword = string(new_pwd)
	dto.RePassword = string(re_pwd)
	dto.Uid = uid

	//参数验证
	valid := validation.Validation{}
	valid.Required(dto.Password, "password")
	valid.Required(dto.NewPassword, "new_password")
	valid.Required(dto.RePassword, "re_password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-3, err.Key+err.Error(), nil)
		}
	}
	//检测输入密码是否一致
	if dto.RePassword != dto.NewPassword {
		c.ErrorJson(-4, "确认密码不一样", nil)
	}

	//修改密码
	service_user := new(service.UserService)
	stat, err := service_user.SetPassword(dto)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 获取用户信息
func (c *UserController) GetUserInfo() {
	auth := c.Ctx.Input.Header("Authorization")
	kv := strings.Split(auth, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		c.ErrorJson(-1, "未带token", nil)
	}
	token := kv[1]
	payload, err := models.ValidateToken(token)
	if err != nil {
		c.ErrorJson(-2, "未登录或登录超时", nil)
	}
	uid := payload.UserID //取得用户id

	//获取用户信息
	service_user := new(service.UserService)
	info, err := service_user.GetById(uid)
	if err != nil {
		c.ErrorJson(-3, err.Error(), nil)
	}
	c.SuccessJson("success", info)
}

// 注册
func (c *UserController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	realname := c.GetString("realname")
	email := c.GetString("email")
	phone_code := c.GetString("phone_code")
	phone := c.GetString("phone")
	sex, _ := c.GetInt("sex")
	avatar := c.GetString("avatar")
	entity := models.User{
		Username:  username,
		Password:  password,
		Realname:  realname,
		Email:     email,
		PhoneCode: phone_code,
		Phone:     phone,
		Avatar:    avatar,
		Sex:       int8(sex),
		RegIp:     c.getClientIp(),
	}

	//用户密码解密
	pwd, _ := base64.StdEncoding.DecodeString(entity.Password)
	entity.Password = string(pwd)

	//参数验证
	valid := validation.Validation{}
	valid.Required(entity.Username, "username")
	valid.Required(entity.Password, "password")
	valid.Required(entity.Realname, "realname")
	valid.Required(entity.Email, "email")
	valid.Required(entity.PhoneCode, "phone_code")
	valid.Required(entity.Phone, "phone")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-1, err.Key+err.Error(), nil)
		}
	}

	//保存
	service_user := new(service.UserService)
	stat, err := service_user.Save(entity)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}

// 修改用户信息
func (c *UserController) Profile() {
	auth := c.Ctx.Input.Header("Authorization")
	kv := strings.Split(auth, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		c.ErrorJson(-1, "未带token", nil)
	}
	token := kv[1]
	payload, err := models.ValidateToken(token)
	if err != nil {
		c.ErrorJson(-2, "未登录或登录超时", nil)
	}
	uid := payload.UserID //取得用户id

	realname := c.GetString("realname")
	email := c.GetString("email")
	phone_code := c.GetString("phone_code")
	phone := c.GetString("phone")
	sex, _ := c.GetInt("sex")
	avatar := c.GetString("avatar")
	entity := models.User{
		Id:        uid,
		Realname:  realname,
		Email:     email,
		PhoneCode: phone_code,
		Phone:     phone,
		Avatar:    avatar,
		Sex:       int8(sex),
	}

	//参数验证
	valid := validation.Validation{}
	valid.Required(entity.Realname, "realname")
	valid.Required(entity.Email, "email")
	valid.Required(entity.PhoneCode, "phone_code")
	valid.Required(entity.Phone, "phone")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.ErrorJson(-3, err.Key+err.Error(), nil)
		}
	}

	//保存
	service_user := new(service.UserService)
	stat, err := service_user.Save(entity)
	if stat < 0 {
		c.ErrorJson(stat, err.Error(), nil)
	}
	c.SuccessJson("success", nil)
}
