package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"remitano-share-video/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/videos",
			beego.NSInclude(
				&controllers.VideoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
