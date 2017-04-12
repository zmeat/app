package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "zmeat.site"
	c.Data["Email"] = "zmeatzhou@126.com"
	c.TplName = "index.tpl"
}
