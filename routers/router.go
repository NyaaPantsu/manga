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
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/series",
			beego.NSInclude(
				&controllers.SeriesController{},
			),
		),

		beego.NSNamespace("/groups",
			beego.NSInclude(
				&controllers.GroupsController{},
			),
		),

		beego.NSNamespace("/chapter_language",
			beego.NSInclude(
				&controllers.ChapterLanguageController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/series_type",
			beego.NSInclude(
				&controllers.SeriesTypeController{},
			),
		),

		beego.NSNamespace("/chapter_user_contributions",
			beego.NSInclude(
				&controllers.ChapterUserContributionsController{},
			),
		),

		beego.NSNamespace("/chapter_group",
			beego.NSInclude(
				&controllers.ChapterGroupController{},
			),
		),

		beego.NSNamespace("/chapter",
			beego.NSInclude(
				&controllers.ChapterController{},
			),
		),

		beego.NSNamespace("/series_chapter",
			beego.NSInclude(
				&controllers.SeriesChapterController{},
			),
		),

		beego.NSNamespace("/series_status",
			beego.NSInclude(
				&controllers.SeriesStatusController{},
			),
		),

		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
