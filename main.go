package main

import (
	_ "github.com/NyaaPantsu/manga/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName: "manga",
		Gclifetime: 3600,
	}
	globalSessions, err := session.NewManager("memory", sessionConfig)
	if err != nil {
		log := logs.NewLogger()
		log.SetLogger(logs.AdapterConsole)
		log.Error(err.Error())
	}
	go globalSessions.GC()
}
func main() {
	if beego.BConfig.RunMode == "dev" {

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}
	beego.Run()
}
