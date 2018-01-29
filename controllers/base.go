package controllers

import (

	"github.com/astaxie/beego"
	"reflect"
)
// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

// JSONSuccess 成功时直接返回json数据
func (b *BaseController) JSONSuccess(i interface{}) {
	t := reflect.ValueOf(i)
	if t.Kind() == reflect.Slice && t.IsNil() {
		i = []string{}
	} else if t.Kind() == reflect.Map && t.IsNil() {
		i = map[string]string{}
	}
	b.Ctx.Output.SetStatus(200)
	b.Data["json"]=i
	b.ServeJSON()
}