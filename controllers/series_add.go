package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/dchest/uniuri"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"html/template"
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
	Username    string `form:"username,text"`
	Id          int    `form:"-"`
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	CoverImage  string `form:"cover, file"`
	TypeName    string `form:"typename, text"`
	TypeDemonym string `form:"typede, text"`
	Status      string `form:"status, text"`
	Tags        string `form:"tags, text"`
	Authors     string `form:"author, text"`
	Artist      string `form:"artist, text"`
}

// Post ...
// @Title Create
// @Description create Series_add
// @Param	body		body 	models.Series_add	true		"body for Series_add content"
// @Success 201 {object} models.Series_add
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
		flash.Error("Series invalid")
		flash.Store(&c.Controller)
		c.Redirect("/comics/add", 302)
		return
	}

	var coverimg string
	exists := models.SeriesNameExists(u.Name)
	if !exists {

		file, header, err := c.GetFile("cover") // where <<this>> is the controller and <<file>> the id of your form field
		if file != nil {
			// get the filename
			filename := header.Filename
			random := uniuri.New()
			// save to server
			coverimg = random + filename
			err := c.SaveToFile("file", "/disk1/covers"+random+filename)
			if err != nil {

				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect("/comics/add", 301)
				return
			}
		}
		if err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/comics/add", 301)
			return
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
		_, err = models.AddSeries(&series)
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
// @Description get Series_add
// @Success 200 {object} models.Series_add
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
// @Param	body		body 	models.Series_add	true		"body for Series_add content"
// @Success 200 {object} models.Series_add
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Series_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Series_add
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Series_addController) Delete() {

}
