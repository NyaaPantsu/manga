package controllers

import (
	"github.com/astaxie/beego"
)

// ReaderController operations for Reader
type ReaderController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReaderController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Reader by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Reader
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ReaderController) GetOne() {

}
