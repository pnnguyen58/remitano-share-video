package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"remitano-share-video/models"
)

type VideoController struct {
	beego.Controller
}

func (v *VideoController) Prepare() {
	token := v.GetSession("token")
	if token == nil {
		v.Redirect("/v1/users/login", 302)
	}
	if userId, err := checkToken(token.(string)); err != nil {
		v.Redirect("/v1/users/login", 302)
	} else {
		_ = models.SetToken(userId, token.(string))
	}

}

// @Title Share Video
// @router /share [post]
func (v *VideoController) Share() {
	userId := v.GetSession("userId")
	video, err := getVideo(v)
	if err == nil {
		video.UserId = userId.(int64)
		err = video.Share()
	}
	v.Redirect("/", 302)
}