package routers

import (
	"github.com/astaxie/beego"
	"github.com/zmeat/app/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
