package controllers

import (
	"github.com/astaxie/beego"
)

// Group_addController operations for Group_add
type Group_addController struct {
	beego.Controller
}

// URLMapping ...
func (c *Group_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Group_add
// @Param	body		body 	models.Group_add	true		"body for Group_add content"
// @Success 201 {object} models.Group_add
// @Failure 403 body is empty
// @router / [post]
func (c *Group_addController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Group_add by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Group_add
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Group_addController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Group_add
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Group_add
// @Failure 403
// @router / [get]
func (c *Group_addController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Group_add
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Group_add	true		"body for Group_add content"
// @Success 200 {object} models.Group_add
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Group_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Group_add
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Group_addController) Delete() {

}
