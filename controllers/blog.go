package controllers

import "github.com/astaxie/beego"

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {

	c.Ctx.ResponseWriter.Write([]byte("ss"))
}
