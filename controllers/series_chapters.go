package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"github.com/NyaaPantsu/manga/utils/zip"
	"github.com/astaxie/beego"
	"github.com/dchest/uniuri"

	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Success  bool          `json:"success"`
	Error    string        `json:"errors,omitempty`
	Response []interface{} `json:"response"`
	Count    int64         `json:"count"`
}

type UploadForm struct {
	Id                    int     `form:"-"`
	Name                  string  `form:"name, text"`
	Title                 string  `form:"title, text"`
	ChapterNumberAbsolute string  `form:"chapternum, te`
	ChapterNumberVolume   float64 `form:"chaptervol, number"`
	VolumeNumber          float64 `form:"volnum, number"`
	ChapterLanguage       string  `form:"languages, text"`
	ReleaseDelay          int     `form:"delay, number"`
	Groups                string  `form:"groups, text"`
}

// SeriesChaptersController operations for SeriesChapters
type SeriesChaptersController struct {
	beego.Controller
}

// URLMapping ...
func (c *SeriesChaptersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SeriesChapters
// @Param	body		body 	models.SeriesChapters	true		"body for SeriesChapters content"
// @Success 201 {int} models.SeriesChapters
// @Failure 403 body is empty
// @router / [post]
func (c *SeriesChaptersController) Post() {
	_, err := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	u := UploadForm{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}
	series, err := models.GetSeriesByName(u.Name)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	user, _ := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	s := strings.Split(u.Groups, ",")

	random := uniuri.New()
	chapter := models.SeriesChapters{
		Title:                 u.Title,
		SeriesId:              &series,
		ContributorId:         user,
		ChapterLanguage:       &models.Languages{Name: u.ChapterLanguage},
		Hash:                  random,
		VolumeNumber:          u.VolumeNumber,
		ChapterNumberVolume:   u.ChapterNumberVolume,
		ChapterNumberAbsolute: u.ChapterNumberAbsolute,
		TimeUploaded:          time.Now(),
	}
	id, err := models.AddSeriesChapters(&chapter)
	for _, v := range s {
		temp := models.SeriesChaptersGroups{
			ChapterId: &chapter,
			GroupName: v,
		}
		models.AddSeriesChaptersGroups(&temp)
	}

	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}

	chapters, err := models.GetSeriesChaptersById(int(id))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}

	//var img string
	files, err := c.GetFiles("files")
	for i := range files {

		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		//img = random + files[i].Filename
		//create destination file making sure the path is writeable.
		dst, err := os.Create("uploads/" + random + files[i].Filename)

		defer dst.Close()
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		err = os.Mkdir("uploads/"+random, os.ModePerm)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		fpath := "uploads/" + random + files[i].Filename
		var images []string
		if zip.IsRar(fpath) {
			images, err = zip.Unrar(fpath, "uploads/"+random)
		} else {
			images, err = zip.Unzip(fpath, "uploads/"+random)
		}
		if len(images) == 0 {
			err := errors.New("error empty archive")
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return

		}
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
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
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

	}
	c.Data["json"] = Response{
		Success: true,
	}
	c.ServeJSON()
	return

}

// GetOne ...
// @Title Get One
// @Description get SeriesChapters by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SeriesChapters
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SeriesChaptersController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	v, err := models.GetSeriesChaptersByHash(id)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	} else {
		var temp []interface{}
		temp = append(temp, v)
		c.Data["json"] = Response{
			Success:  true,
			Response: temp,
			Count:    1,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SeriesChapters
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SeriesChapters
// @Failure 403
// @router / [get]
func (c *SeriesChaptersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err := errors.New("Error: invalid query key/value pair")
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, count, err := models.GetAllSeriesChapters(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

	} else {
		c.Data["json"] = Response{
			Success:  true,
			Response: l,
			Count:    count,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SeriesChapters
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SeriesChapters	true		"body for SeriesChapters content"
// @Success 200 {object} models.SeriesChapters
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SeriesChaptersController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SeriesChapters{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSeriesChaptersById(&v); err == nil {
			c.Data["json"] = Response{
				Success: true,
			}
		} else {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SeriesChapters
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SeriesChaptersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSeriesChapters(id); err == nil {
		c.Data["json"] = Response{
			Success: true,
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}
