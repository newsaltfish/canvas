package routers

import (
	"canvas/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pack/*", &controllers.PackController{})
	beego.Router("/app/login", &controllers.PackController{})

}
