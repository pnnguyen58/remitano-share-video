package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "remitano-share-video/routers"
)

func main() {
	beego.Run()
}

