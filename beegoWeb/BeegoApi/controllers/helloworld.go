package controllers

import (
	"context"

	beego "github.com/beego/beego/v2/server/web"
)

type HelloWorldController struct {
	beego.Controller
}

func Init(ctx *context.Context, controllerName, actionName string, app interface{}) {

}
func (c *HelloWorldController) Get() {

}
func (c *HelloWorldController) Finish() {

}
func (c *HelloWorldController) Trace() {

}
