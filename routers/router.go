// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/NyaaPantsu/manga/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.SearchController{}, "get:GetAll")
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/comics", &controllers.ComicsController{}, "get:GetAll")
	beego.Router("/comics/:name", &controllers.ComicsController{}, "get:GetOne")
	beego.Router("/comics/add", &controllers.Series_addController{})
	beego.Router("/groups", &controllers.Groups_addController{})
	beego.Router("/groups/add", &controllers.Groups_addController{})

	beego.Router("/reader", &controllers.ReaderController{})

	ns1 := beego.NewNamespace("/auth",
		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/signup",
			beego.NSInclude(
				&controllers.SignupController{},
			),
		),
	)
	ns2 := beego.NewNamespace("/mod",
		beego.NSNamespace("/reports",
			beego.NSInclude(
				&controllers.SignupController{},
			),
		),
	)

	ns := beego.NewNamespace("/api/v1",

		beego.NSNamespace("/statuses",
			beego.NSInclude(
				&controllers.StatusesController{},
			),
		),

		beego.NSNamespace("/languages",
			beego.NSInclude(
				&controllers.LanguagesController{},
			),
		),

		beego.NSNamespace("/groups_scanlation",
			beego.NSInclude(
				&controllers.GroupsScanlationController{},
			),
		),

		beego.NSNamespace("/series_chapters",
			beego.NSInclude(
				&controllers.SeriesChaptersController{},
			),
		),

		beego.NSNamespace("/series",
			beego.NSInclude(
				&controllers.SeriesController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns1)
	beego.AddNamespace(ns2)
}
