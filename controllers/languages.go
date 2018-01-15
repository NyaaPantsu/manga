package controllers

import (
	"encoding/json"
	"github.com/NyaaPantsu/manga/models"

	"github.com/astaxie/beego"
)

// LanguagesController operations for Languages
type LanguagesController struct {
	beego.Controller
}

// URLMapping ...
func (c *LanguagesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Languages
// @Param	body		body 	models.Languages	true		"body for Languages content"
// @Success 201 {int} models.Languages
// @Failure 403 body is empty
// @router / [post]
func (c *LanguagesController) Post() {
	var v models.Languages
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLanguages(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Languages
// @Success 200 {object} []models.Languages
// @Failure 403
// @router / [get]
func (c *LanguagesController) GetAll() {

	l, err := models.GetAllLanguages()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Languages
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LanguagesController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if err := models.DeleteLanguages(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
