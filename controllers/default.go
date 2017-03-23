package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	c.Data["json"] = struct{ Name string }{"test"}
	c.ServeJSON()
}