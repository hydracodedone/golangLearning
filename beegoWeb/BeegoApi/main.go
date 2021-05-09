package main

import (
	_ "BeegoApi/routers"

	"BeegoApi/crontab"

	"github.com/beego/beego/v2/adapter/toolbox"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	crontab.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
