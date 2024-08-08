package main

import (
	"BeegoAPIDemo/crontab"
	_ "BeegoAPIDemo/routers"

	"github.com/astaxie/beego/toolbox"

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
