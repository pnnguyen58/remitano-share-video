package controllers

import (
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"remitano-share-video/models"
)

type VideoController struct {
	beego.Controller
}

// @Title Share Video
// @router /share [post]
func (v *VideoController) Share() {
	var response Response
	code := http.StatusOK
	if userId, e := checkToken(v.Ctx); e == nil && userId > 0 {
		video, err := verifyVideo(v.Ctx)
		if err == nil {
			video.UserId = userId
			if err = video.Share(); err != nil {
				code = http.StatusInternalServerError
			}
		}
		if err != nil {
			response.Error = err.Error()
		}
		response.Data = video
	} else {
		code = http.StatusUnauthorized
		response.Error = errors.New("unauthorized").Error()
	}
	v.Ctx.Output.SetStatus(code)
	v.Data["json"] = response
	defer v.Ctx.Request.Body.Close()
	_ = v.ServeJSON()
}

// @Title Share Video
// @router /list [get]
func (v *VideoController) List() {
	var response Response
	code := http.StatusOK
	if userId, e := checkToken(v.Ctx); e == nil && userId > 0 {
		var limit, offset int
		_ = v.Ctx.Input.Bind(&limit, "limit")
		_ = v.Ctx.Input.Bind(&offset, "offset")
		videos, err := models.GetVideos(userId, limit, offset)
		if err != nil {
			response.Error = err.Error()
		}
		response.Data = videos
	} else {
		code = http.StatusUnauthorized
		response.Error = errors.New("unauthorized").Error()
	}
	v.Ctx.Output.SetStatus(code)
	v.Data["json"] = response
	defer v.Ctx.Request.Body.Close()
	_ = v.ServeJSON()
}