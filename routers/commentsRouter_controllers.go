package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterGroupController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterLanguageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ChapterUserContributionsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:ComicController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:GroupsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesChapterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SeriesTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SignupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SignupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SignupController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:SignupController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/NyaaPantsu/manga/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
