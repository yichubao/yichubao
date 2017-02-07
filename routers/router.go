package routers

import (
	"yichubao/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/blog", &controllers.BlogController{})
}
