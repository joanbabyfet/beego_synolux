// 路由文件
package routers

import (
	"synolux/controllers/admin"
	controllers "synolux/controllers/api"

	beego "github.com/beego/beego/v2/server/web"
)

// 初始化
func init() {
	//用户端接口
	api := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/home_article", &controllers.ArticleController{}, "get:HomeArticle"),
			beego.NSRouter("/article", &controllers.ArticleController{}, "get:Index"),
			beego.NSRouter("/article/detail", &controllers.ArticleController{}, "get:Detail"),
			beego.NSRouter("/article/save", &controllers.ArticleController{}, "post:Save"),
			beego.NSRouter("/article/delete", &controllers.ArticleController{}, "post:Delete"),
			beego.NSRouter("/article/enable", &controllers.ArticleController{}, "post:Enable"),
			beego.NSRouter("/article/disable", &controllers.ArticleController{}, "post:Disable"),
			beego.NSRouter("/captcha", &controllers.CommonController{}, "get:Captcha"), //获取验证码
			beego.NSRouter("/ip", &controllers.CommonController{}, "get:Ip"),
			beego.NSRouter("/ping", &controllers.CommonController{}, "get:Ping"),
			beego.NSRouter("/captcha", &controllers.CommonController{}, "get:Captcha"), //获取验证码
			beego.NSRouter("/upload", &controllers.UploadController{}, "post:Upload"),
			beego.NSRouter("/download", &controllers.UploadController{}, "get:Download"),
			beego.NSRouter("/chat_gpt", &controllers.CommonController{}, "get:ChatGPT"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),            //退出
			beego.NSRouter("/set_password", &controllers.UserController{}, "post:SetPassword"), //修改密码
			beego.NSRouter("/get_userinfo", &controllers.UserController{}, "get:GetUserInfo"),  //获取用户信息
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),        //注册
			beego.NSRouter("/profile", &controllers.UserController{}, "post:Profile"),          //修改用户信息
			beego.NSRouter("/feedback", &controllers.FeedbackController{}, "post:Save"),        //提交反馈
		),
	)

	//ws服务端
	beego.Router("/ws", &controllers.WSController{}, "get:Index")

	//后端接口
	admin := beego.NewNamespace("/admin_api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/article", &admin.ArticleController{}, "get:Index"),
			beego.NSRouter("/article/detail", &admin.ArticleController{}, "get:Detail"),
			beego.NSRouter("/article/save", &admin.ArticleController{}, "post:Save"),
			beego.NSRouter("/article/delete", &admin.ArticleController{}, "post:Delete"),
			beego.NSRouter("/article/enable", &admin.ArticleController{}, "post:Enable"),
			beego.NSRouter("/article/disable", &admin.ArticleController{}, "post:Disable"),
			beego.NSRouter("/ad", &admin.AdController{}, "get:Index"),
			beego.NSRouter("/ad/detail", &admin.AdController{}, "get:Detail"),
			beego.NSRouter("/ad/save", &admin.AdController{}, "post:Save"),
			beego.NSRouter("/ad/delete", &admin.AdController{}, "post:Delete"),
			beego.NSRouter("/ad/enable", &admin.AdController{}, "post:Enable"),
			beego.NSRouter("/ad/disable", &admin.AdController{}, "post:Disable"),
			beego.NSRouter("/upload", &admin.UploadController{}, "post:Upload"),
			beego.NSRouter("/download", &admin.UploadController{}, "get:Download"),
			beego.NSRouter("/captcha", &admin.CommonController{}, "get:Captcha"), //获取验证码
			beego.NSRouter("/chat_gpt", &admin.CommonController{}, "get:ChatGPT"),
			beego.NSRouter("/ip", &admin.CommonController{}, "get:Ip"),
			beego.NSRouter("/ping", &admin.CommonController{}, "get:Ping"),
			beego.NSRouter("/test", &admin.TestController{}, "get:Test"),
		),
	)
	beego.AddNamespace(api, admin)
}
