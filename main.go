package main

import (
	"html/template"
	"net/http"
	_ "yichubao/routers"

	"github.com/astaxie/beego"
)

func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles("./views/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	beego.ErrorHandler("404", pageNotFound)

	beego.Run()
}
