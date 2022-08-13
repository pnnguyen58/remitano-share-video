package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Prepare() {
}

// @Title Register
// @router /login [get]
func (u *UserController) GetLogin() {
	u.TplName = "login.html"
}

// @Title Register
// @router /register [get]
func (u *UserController) GetRegister() {
	u.TplName = "register.html"
}

// @Title Register
// @router /register [post]
func (u *UserController) Register() {
	user, err := getUser(u)
	if err == nil {
		if err = user.Register(); err == nil {
			u.TplName = "login.html"
		} else {
			u.TplName = "register.html"
		}
	} else {
		u.TplName = "register.html"
	}
}

// @Title Login
// @router /login [post]
func (u *UserController) Login() {
	user, err := getUser(u)
	if err == nil {
		if err = user.Login(); err == nil {
			_ = u.SetSession("token", user.Token)
			_ = u.SetSession("userId", user.ID)
		}
	}
	if err != nil {
		u.TplName = "login.html"
	} else {
		u.Redirect("/", 302)
	}
}


// @Title Register
// @router /register/test [post]
func (u *UserController) RegisterTest() {
	var response Response
	code := http.StatusOK
	user, err := verifyUser(u.Ctx)
	if err == nil {
		if err = user.Register(); err == nil {
			response.Data = user
		}
	}
	if err != nil {
		response.Error = err.Error()
		code = http.StatusInternalServerError
	}
	u.Ctx.Output.SetStatus(code)
	u.Data["json"] = response
	defer u.Ctx.Request.Body.Close()
	_ = u.ServeJSON()
}

// @Title Login
// @router /login/test [post]
func (u *UserController) LoginTest() {
	var response Response
	code := http.StatusOK
	user, err := verifyUser(u.Ctx)
	if err == nil {
		if err = user.Login(); err == nil {
			response.Data = user
		}
	}
	if err != nil {
		response.Error = err.Error()
		code = http.StatusInternalServerError
	}
	u.Ctx.Output.SetStatus(code)
	u.Data["json"] = response
	defer u.Ctx.Request.Body.Close()
	_ = u.ServeJSON()
}