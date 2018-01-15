package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dchest/uniuri"

	"html/template"
	"io"
	"os"
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
	Title                 string `form:"title, text"`
	ChapterNumberAbsolute int    `form:"chapternum, number`
	ChapterNumberVolume   int    `form:"chaptervol, number"`
	VolumeNumber          int    `form:"volnum, number"`
	ChapterLanguage       string `form:"languages, text"`
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
	if !c.IsLogin {
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
	series, err := models.GetSeriesByName(u.Title)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/upload", 302)
		return
	}
	//var img string
	files, err := c.GetFiles("files")
	for i := range files {

		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}

		random := uniuri.New()
		//img = random + files[i].Filename
		//create destination file making sure the path is writeable.
		dst, err := os.Create("upload/" + random + files[i].Filename)
		defer dst.Close()
		if err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}
	}
	chapters := models.SeriesChapters{
		Title:           u.Title,
		SeriesId:        &series,
		ChapterLanguage: &models.Languages{Name: u.ChapterLanguage},
		ContributorId:   c.Userinfo,
		Hash:            uniuri.New(),
	}
	id, err := models.AddSeriesChapters(&chapters)
	log := logs.GetLogger()
	log.Println(id)
	if err != nil {

		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/upload", 301)
		return
	}
	flash.Success("Successfully added series!")
	flash.Store(&c.Controller)
	c.Redirect("/upload", 301)
	return

}

// Get ...
// @Title Get
// @Description get Upload by id
// @Success 200 {object} models.Upload
// @router / [get]
func (c *UploadController) Get() {
	flash := beego.NewFlash()
	if !c.IsLogin {
		flash.Error("Error you must be logged in to upload")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}
	c.TplName = "upload.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	l, _ := models.GetAllLanguages()
	series, _ := models.GetAllSeriesArray()
	c.Data["languages"] = l
	c.Data["series"] = series

	c.Render()
}
