package main

import (
	_ "github.com/NyaaPantsu/manga/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/NyaaPantsu/manga/utils/templates"
	"github.com/astaxie/beego/session"
)

func main() {

	orm.Debug = true
	beego.AddFuncMap("until", templates.Until)
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	sessionconf := &session.ManagerConfig{
		CookieName: "manga",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}
