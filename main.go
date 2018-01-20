package main

import (
	_ "github.com/NyaaPantsu/manga/routers"
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func main() {

	orm.Debug = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.Run()
}
