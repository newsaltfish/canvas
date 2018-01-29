package controllers

import "github.com/astaxie/beego"

// PackController 传输
type PackController struct {
	beego.Controller
}

// GetLoc 银行卡认证(模拟成绑卡)
// @Title 银行卡认证
// @Description 银行卡认证
// @Param	body		body 	models.BaofooVerifycardReq	true		"The object content"
// @Success 200 发送成功
// @Failure 202 发送失败
// @router /getloc [post]
func (c *PackController) GetLoc() {
	c.Data["json"] = struct{ Name string }{"test"}
	c.ServeJSON()
}

