package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/dchest/uniuri"

	"github.com/NyaaPantsu/manga/utils/zip"
	"html/template"
	"io"
	"os"
	"time"
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
	Id                    int     `form:"-"`
	Title                 string  `form:"title, text"`
	ChapterNumberAbsolute float64 `form:"chapternum, number`
	ChapterNumberVolume   float64 `form:"chaptervol, number"`
	VolumeNumber          float64 `form:"volnum, number"`
	ChapterLanguage       string  `form:"languages, text"`
	ReleaseDelay          int     `form:"delay, number"`
	Groups1               string  `form:"group1, text"`
	Groups2               string  `form:"group2, text"`
	Groups3               string  `form:"group3, text"`
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

	random := uniuri.New()
	chapter := models.SeriesChapters{
		Title:                 u.Title,
		SeriesId:              &series,
		ChapterLanguage:       &models.Languages{Name: u.ChapterLanguage},
		ContributorId:         c.Userinfo,
		Hash:                  random,
		VolumeNumber:          u.VolumeNumber,
		ChapterNumberVolume:   u.ChapterNumberVolume,
		ChapterNumberAbsolute: u.ChapterNumberAbsolute,
		TimeUploaded:          time.Now(),
	}
	id, err := models.AddSeriesChapters(&chapter)

	if err != nil {

		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/upload", 301)
		return
	}

	chapters, err := models.GetSeriesChaptersById(int(id))
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/upload", 301)
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

		//img = random + files[i].Filename
		//create destination file making sure the path is writeable.
		dst, err := os.Create("uploads/" + random + files[i].Filename)

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

		err = os.Mkdir("uploads/"+random, os.ModePerm)
		if err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}

		images, err := zip.Unzip("uploads/"+random+files[i].Filename, "uploads/"+random)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}
		var temp []models.SeriesChaptersFiles
		for _, cond := range images {
			temp2 := models.SeriesChaptersFiles{
				ChapterId: chapters,
				Name:      cond,
			}
			temp = append(temp, temp2)

		}
		_, err = models.AddMultiChapterFiles(temp)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/upload", 301)
			return
		}

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
