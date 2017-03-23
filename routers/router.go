package routers

import (
	"github.com/astaxie/beego"
	"gitlab.cern.ch/flutter/flutter-rest/controllers"
)

func init() {
	beego.Debug("Router")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
}
