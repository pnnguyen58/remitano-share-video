package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginTest",
            Router: "/login/test",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetLogin",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "RegisterTest",
            Router: "/register/test",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetRegister",
            Router: "/register",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["remitano-share-video/controllers:VideoController"] = append(beego.GlobalControllerRouter["remitano-share-video/controllers:VideoController"],
        beego.ControllerComments{
            Method: "Share",
            Router: "/share",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
