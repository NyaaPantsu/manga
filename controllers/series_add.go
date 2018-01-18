package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/resize"
	"github.com/NyaaPantsu/manga/utils/split"
	"github.com/astaxie/beego"
	"github.com/dchest/uniuri"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"html/template"
	"io"
	"os"
	"path"
	"strings"
)

// Series_addController operations for Series_add
type Series_addController struct {
	BaseController
}

// URLMapping ...
func (c *Series_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type SeriesForm struct {
	Id          int    `form:"-"`
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	CoverImage  string `form:"cover, file"`
	TypeName    string `form:"typename, text"`
	TypeDemonym string `form:"typede, text"`
	Status      string `form:"status, text"`
	Tags        string `form:"tags, text"`
	Authors     string `form:"authors, text"`
	Artist      string `form:"artists, text"`
	Mature      int    `form:"mature, checkbox"`
}

// Post ...
// @Title Create
// @Description create Series_add
// @Param	body		body 	models.Series	true		"body for Series_add content"
// @Success 201 {object} models.Series
// @Failure 403 body is empty
// @router / [post]
func (c *Series_addController) Post() {
	flash := beego.NewFlash()
	if !c.IsLogin {
		flash.Error("Error you need to be logged in")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	u := SeriesForm{}
	if err := c.ParseForm(&u); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/comics/add", 302)
		return
	}

	var coverimg string
	exists := models.SeriesNameExists(u.Name)
	if !exists {
		files, err := c.GetFiles("cover")
		for i := range files {

			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {

				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect("/comics/add", 301)
				return
			}
			random := uniuri.New()
			coverimg = random + files[i].Filename
			//create destination file making sure the path is writeable.
			dst, err := os.Create("uploads/covers/" + random + files[i].Filename)
			defer dst.Close()
			if err != nil {

				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect("/comics/add", 301)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {

				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect("/comics/add", 301)
				return
			}
			pat := "uploads/covers/" + random + files[i].Filename
			if strings.ToLower(path.Ext(files[i].Filename)) == ".png" {
				resize.ResizePng(pat, pat+"_thumb")
			} else if strings.ToLower(path.Ext(files[i].Filename)) == ".jpg" {
				resize.ResizeJpg(pat, pat+"_thumb")

			}
		}

		unsafe := blackfriday.Run([]byte(u.Description))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		series := models.Series{
			Name:        u.Name,
			Description: string(html),
			TypeName:    u.TypeName,
			CoverImage:  coverimg,
			TypeDemonym: u.TypeDemonym,
			Status: &models.Statuses{
				Name: u.Status,
			},
		}
		id, err := models.AddSeries(&series)
		series.Id = int(id)

		tempTags, _ := split.ProcessTags(&series, u.Tags)
		authors, _ := split.CreateTags(&series, "author", u.Authors)
		artists, _ := split.CreateTags(&series, "artist", u.Artist)
		if u.Mature == 1 {
			matureTag, _ := split.CreateTags(&series, "content", "mature")
			models.AddMultiSeriesTags(matureTag)
		}
		models.AddMultiSeriesTags(tempTags)
		models.AddMultiSeriesTags(authors)
		models.AddMultiSeriesTags(artists)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/comics/add", 301)
			return
		}
		flash.Success("Successfully added series!")
		flash.Store(&c.Controller)
		c.Redirect("/comics/add", 301)
		return

	}

	flash.Error("Adding series failed")
	flash.Store(&c.Controller)
	c.Redirect("/comics/add", 302)
	return
}

// Get ...
// @Title Get
// @Description get Series
// @Success 200 {object} models.Series
// @Failure 403
// @router / [get]
func (c *Series_addController) Get() {
	flash := beego.NewFlash()
	if !c.IsLogin {
		flash.Error("Error you need to be logged in")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "series_add.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}

// Put ...
// @Title Put
// @Description update the Series_add
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Series	true		"body for Series_add content"
// @Success 200 {object} models.Series
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Series_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Series
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Series_addController) Delete() {

}
