package routers

import (
	"github.com/NyaaPantsu/manga/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
