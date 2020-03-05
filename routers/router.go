package routers

import (
	"StitpProjectVersion1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //注册
    beego.Router("/register", &controllers.RegisterController{})
    //登录
    beego.Router("/login", &controllers.LoginController{})
}
