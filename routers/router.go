package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"remitano-share-video/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
