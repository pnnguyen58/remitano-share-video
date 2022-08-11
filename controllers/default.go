package controllers

import (
	"encoding/json"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"remitano-share-video/models"
)

type MainController struct {
	beego.Controller
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"err"`
}

type Header struct {
	Authorization []string `json:"Authorization"`
}

func (c *MainController) Get() {
	c.Data["Username"] = "beego.me"
	c.Data["Password"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func checkToken(ctx *context.Context) (userId int64, err error) {
	var header Header
	data, _ := json.Marshal(ctx.Request.Header)
	if err = json.Unmarshal(data, &header); err == nil {
		if len(header.Authorization) > 0 {
			return models.CheckToken(header.Authorization[0])
		}
	}
	return
}

func verifyUser(ctx *context.Context) (user models.User, err error) {
	if err = json.Unmarshal(ctx.Input.CopyBody(1024), &user); err != nil {
		return
	} else if user.Username == "" || user.Password == "" {
		return user, errors.New("missing required input")
	}
	return
}

func verifyVideo(ctx *context.Context) (video models.Video, err error) {
	if err = json.Unmarshal(ctx.Input.CopyBody(1024), &video); err != nil {
		return
	} else if video.Title == "" || video.Link == ""{
		return video, errors.New("missing required input")
	}
	return
}