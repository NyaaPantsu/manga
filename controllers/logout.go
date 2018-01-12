package controllers

import (
	"github.com/astaxie/beego"
)

// LogoutController operations for Logout
type LogoutController struct {
	beego.Controller
}

// URLMapping ...
func (c *LogoutController) URLMapping() {
	c.Mapping("Get", c.Post)
}

// Get ...
// @Title Logout
// @Description Logout
// @Param	body		body 	models.Logout	true		"body for Logout content"
// @Success 201 {object} models.Logout
// @Failure 403 body is empty
// @router / [get]
func (c *LogoutController) Get() {
	// Check if user is logged in
	session := c.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
		session.Delete("UserName")
		session.Delete("Admin")
	}
	c.Redirect("/", 301)
}
