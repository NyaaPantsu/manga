package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
)

// SignupController operations for Signup
type SignupController struct {
	beego.Controller
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Signup	true		"body for Users content"
// @Success 201 {string} models.Users
// @Failure 403 {string}
// @router / [post]
func (c *UsersController) Post() {
	var v models.Signup
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		username := models.UsernameExists(v.Username)
		email := models.EmailExists(v.Email)
		if email || username {
			c.Data["json"] = "{'error':'Username or email in use'}"
			c.ServeJSON()
			return
		}

		password := []byte(v.Password)

		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			c.Data["json"] = "{'error': '" + err.Error() + "'}"
			c.ServeJSON()
			return

		}
		users := models.Users{
			Username:     v.Username,
			Email:        v.Email,
			PasswordHash: string(hashedPassword),
		}

		if _, err := models.AddUsers(&users); err == nil {
			c.Ctx.Output.SetStatus(201)
			v.Password = ""
			c.Data["json"] = v
		} else {
			c.Data["json"] = "{'error': '" + err.Error() + "'}"
		}
	} else {
		c.Data["json"] = "{'error': '" + err.Error() + "'}"
	}
	c.ServeJSON()
}
