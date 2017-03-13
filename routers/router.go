package routers

import (
	"flutter-rest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
}
