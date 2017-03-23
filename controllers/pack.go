package controllers

import "github.com/astaxie/beego"

type PackController struct {
	beego.Controller
}

func (c *PackController) Post() {
	c.Data["json"] = struct{ Name string }{"test"}
	c.ServeJSON()
}
