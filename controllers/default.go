package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	head := c.Ctx.Request

	fmt.Println(head)

	c.Data["Website"] = "anla.io"
	c.Data["Email"] = "anlasheng@gmail.com"
	c.TplName = "index.html"
}
