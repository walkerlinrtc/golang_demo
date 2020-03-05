package controllers

import (
	"StitpProjectVersion1/models"
	"StitpProjectVersion1/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, "password:", password)
	id := models.QueryUserWithParam(username,utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code":1 , "massage":"恭喜登录成功"}
	} else{
		this.Data["json"] = map[string]interface{}{"code":0 , "massage":"抱歉，登录失败"}
	}
	this.ServeJSON()

}

type UserInfo struct {
	Username string              //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
	Password string `form:"password"` //也可以指定form表单中对应的name值，如果不写将无法解析该参数的值

}

func (c *LoginController) Login() {
	//获取cookie
	username := c.Ctx.GetCookie("username")
	password := c.Ctx.GetCookie("password")

	//验证用户名和密码：
	if username != "" {
		c.Ctx.WriteString("Username:" + username + ",Password:" + password)
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:9527/login" method="post">
									用户名：<input type ="text" name="Username" />
									<br/>
									密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
									<br/>
									<input type="submit" value="提交">
								</form></html>`)
	}
}

/*func (c *LoginController) Post() {
	u := UserInfo{}
	if err := c.ParseForm(&u); err != nil {
		log.Panic(err)
	}
	//设置cookie
	c.Ctx.SetCookie("username",u.Username,100,"/")
	c.Ctx.SetCookie("password",u.Password,100,"/")

	//设置session
	c.SetSession("username",u.Username)
	c.SetSession("password",u.Password)
	c.Ctx.WriteString("Username:" + u.Username + ",Password:" + u.Password)

}*/
