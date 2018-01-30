package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/resize"
	"github.com/NyaaPantsu/manga/utils/split"
	"github.com/astaxie/beego"
	"github.com/dchest/uniuri"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// SeriesController operations for Series
type SeriesController struct {
	beego.Controller
}

// URLMapping ...
func (c *SeriesController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

type SeriesForm struct {
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	CoverImage  string `form:"cover, file"`
	TypeName    string `form:"typename, text"`
	TypeDemonym string `form:"typedemonym, text"`
	Status      string `form:"status, text"`
	Tags        string `form:"tags, text"`
	Authors     string `form:"authors, text"`
	Artist      string `form:"artists, text"`
	Mature      bool   `form:"mature, checkbox"`
}

// Post ...
// @Title Create
// @Description create Series_add
// @Param	body		body 	models.Series	true		"body for Series_add content"
// @Success 201 {object} models.Series
// @Failure 403 body is empty
// @router / [post]
func (c *SeriesController) Post() {

	u := SeriesForm{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
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
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}
			random := uniuri.New()
			coverimg = random + files[i].Filename
			//create destination file making sure the path is writeable.
			dst, err := os.Create("uploads/covers/" + random + files[i].Filename)
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
			pat := "uploads/covers/" + random + files[i].Filename
			err = resize.ResizeImage(pat, pat+"_thumb")
			if err != nil {
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}

		}

		unsafe := blackfriday.Run([]byte(u.Description))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		status := models.Statuses{Name: u.Status}
		series := models.Series{
			Name:        u.Name,
			Description: string(html),
			TypeName:    u.TypeName,
			CoverImage:  coverimg,
			TypeDemonym: u.TypeDemonym,
			Status:      &status,
		}
		id, err := models.AddSeries(&series)
		series.Id = int(id)

		split.CreateTags(&series, u.Tags, "genre")
		split.CreateTags(&series, u.Authors, "author")
		split.CreateTags(&series, u.Artist, "artist")
		if u.Mature {
			split.CreateTags(&series, "mature", "content")
		}
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return
		}
		c.Data["json"] = Response{
			Success: true,
		}
		c.ServeJSON()
		return

	}
	err := errors.New("Adding series failed")
	c.Data["json"] = Response{
		Success: false,
		Error:   err.Error(),
	}
	c.ServeJSON()
	return

}

// GetOne ...
// @Title Get One
// @Description get Series by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Series
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SeriesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSeriesById(id)
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
// @Description get Series
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Series
// @Failure 403
// @router / [get]
func (c *SeriesController) GetAll() {
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

	l, count, err := models.GetAllSeries(query, fields, sortby, order, offset, limit)

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
// @Description update the Series
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Series	true		"body for Series content"
// @Success 200 {object} models.Series
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SeriesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Series{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSeriesById(&v); err == nil {
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
// @Description delete the Series
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SeriesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSeries(id); err == nil {
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
