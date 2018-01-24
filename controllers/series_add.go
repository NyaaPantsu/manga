package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/resize"
	"github.com/NyaaPantsu/manga/utils/split"
	"github.com/dchest/uniuri"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"io"
	"os"
)

// Series_addController operations for Series_add
type Series_addController struct {
	BaseController
}

// URLMapping ...
func (c *Series_addController) URLMapping() {
	c.Mapping("Post", c.Post)
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

	u := SeriesForm{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = "{'error': '" + err.Error() + "'}"
		c.ServeJSON()
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
				c.Data["json"] = "{'error': '" + err.Error() + "'}"
				c.ServeJSON()
				return
			}
			random := uniuri.New()
			coverimg = random + files[i].Filename
			//create destination file making sure the path is writeable.
			dst, err := os.Create("uploads/covers/" + random + files[i].Filename)
			defer dst.Close()
			if err != nil {
				c.Data["json"] = "{'error': '" + err.Error() + "'}"
				c.ServeJSON()
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				c.Data["json"] = "{'error': '" + err.Error() + "'}"
				c.ServeJSON()
				return
			}
			pat := "uploads/covers/" + random + files[i].Filename
			err = resize.ResizeImage(pat, pat+"_thumb")
			if err != nil {
				c.Data["json"] = "{'error': '" + err.Error() + "'}"
				c.ServeJSON()
				return
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
			c.Data["json"] = "{'error': '" + err.Error() + "'}"
			c.ServeJSON()
			return
		}
		c.Data["json"] = "{'success': 'Added series!'}"

	}

	c.Data["json"] = "{'error': 'Adding series failed'}"
	c.ServeJSON()
	return

}
