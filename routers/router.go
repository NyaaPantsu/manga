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
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Router("/rss/latest", &controllers.RssController{}, "get:GetNew")
	beego.Router("/rss/followed", &controllers.RssController{}, "get:GetFollowed")

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
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.SeriesChaptersController{},
			),
		),

		beego.NSNamespace("/series",
			beego.NSInclude(
				&controllers.SeriesController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSRouter("/logout",
				&controllers.LogoutController{},
			),
			beego.NSRouter("/login",
				&controllers.LoginController{},
				"post:Post",
			),
			beego.NSRouter("/register",
				&controllers.SignupController{},
			),
		),
		beego.NSRouter("/reader/:hash", &controllers.ReaderController{}, "get:GetOne"),
		beego.NSRouter("/comics/add", &controllers.Series_addController{}),
		beego.NSRouter("/groups/add", &controllers.Groups_addController{}, "get:Get;post:Post"),
		beego.NSRouter("/follow/:id", &controllers.FollowController{}, "get,post:ToggleFollow"),
		beego.NSRouter("/follows", &controllers.ImportController{}),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)
	beego.SetStaticPath("/uploads", "uploads")
}
