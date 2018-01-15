package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"

	"html/template"
)

// UploadController operations for Upload
type UploadController struct {
	BaseController
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

type UploadForm struct {
	Id                    int    `form:"-"`
	Title                 string `form:"series, text"`
	ChapterNumberAbsolute int    `form:"chapternum, number`
	ChapterNumberVolume   int    `form:"chaptervol, number"`
	VolumeNumber          int    `form:"volnum, number"`
	ChapterLanguage       string `form:"languages, number"`
	ReleaseDelay          int    `form:"delay, number"`
	Groups1               string `form:"group1, text"`
	Groups2               string `form:"group2, text"`
	Groups3               string `form:"group3, text"`
}

// Post ...
// @Title Create
// @Description create Upload
// @Param	body		body 	models.Upload	true		"body for Upload content"
// @Success 201 {object} models.Upload
// @Failure 403 body is empty
// @router / [post]
func (c *UploadController) Post() {
	flash := beego.NewFlash()
	if c.IsLogin {
		flash.Error("Error you must be logged in to upload")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}

	u := UploadForm{}
	if err := c.ParseForm(&u); err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/upload", 302)
		return
	}
	file, header, err := c.GetFile("file") // where <<this>> is the controller and <<file>> the id of your form field
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}
	if file != nil {
		// get the filename
		fileName := header.Filename
		// save to server
		err := c.SaveToFile("file", "/disk1/archives"+fileName)
		if err != nil {
			flash.Warning(err.Error())
			flash.Store(&c.Controller)
			return
		}
		if err != nil {
			flash.Warning(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/uploads", 302)
			return

		}

	}

}

// Get ...
// @Title Get
// @Description get Upload by id
// @Success 200 {object} models.Upload
// @router / [get]
func (c *UploadController) Get() {
	c.TplName = "upload.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	l, _ := models.GetAllLanguages()
	series, _ := models.GetAllSeriesArray()
	c.Data["languages"] = l
	c.Data["series"] = series

	c.Render()
}
