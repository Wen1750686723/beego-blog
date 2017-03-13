package main

import (
	"myproject/g"
	_ "myproject/routers"

	"github.com/astaxie/beego"
)

func main() {
	g.InitEnv()
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
