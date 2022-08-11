package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["Error"] = "beego.me"
	u.Data["Password"] = "astaxie@gmail.com"
	u.TplName = "home.html"
}


// @Title Register
// @router /register [post]
func (u *UserController) Register() {
	var response Response
	code := http.StatusOK
	user, err := verifyUser(u.Ctx)
	if err == nil {
		if err = user.Register(); err != nil {
			code = http.StatusInternalServerError
		}
	}
	if err != nil {
		response.Error = err.Error()
	}
	response.Data = user
	u.Ctx.Output.SetStatus(code)
	u.Data["json"] = response
	defer u.Ctx.Request.Body.Close()
	_ = u.ServeJSON()
}

// @Title Login
// @router /login [post]
func (u *UserController) Login() {
	var response Response
	code := http.StatusOK
	user, err := verifyUser(u.Ctx)
	if err == nil {
		if err = user.Login(); err != nil {
			code = http.StatusInternalServerError
		}
	}
	if err != nil {
		response.Error = err.Error()
	}
	response.Data = user
	u.Ctx.Output.SetStatus(code)
	u.Data["json"] = response
	defer u.Ctx.Request.Body.Close()
	_ = u.ServeJSON()
}