package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
	}
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {

	}
	password := []byte(u.Password)
	// Comparing the password with the hash

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {

	}
	c.Redirect("/", 301)

}

// Get ...
// @Title GetAll
// @Description get Login
// @Success 200 {object} models.Login
// @Failure 403
// @router / [get]
func (c *LoginController) Get() {
	c.TplName = "login.tpl"
	c.Layout = "index.tpl"
	c.Data["Form"] = &User{}
	c.Render()
}
