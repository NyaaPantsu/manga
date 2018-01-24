package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"errors"
)

// RegisterController operations for Signup
type RegisterController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegisterController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Signup	true		"body for Users content"
// @Success 201 {string} models.Users
// @Failure 403 {string}
// @router / [post]
func (c *RegisterController) Post() {
	var v models.Signup
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		username := models.UsernameExists(v.Username)
		email := models.EmailExists(v.Email)
		if email || username {
			err := errors.New("Username or email in use")
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return

		}

		password := []byte(v.Password)

		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
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
			users.PasswordHash = ""
			var temp []interface{}
			temp = append(temp, users)
			c.Data["json"] = Response{
				Success:  true,
				Response: temp,
				Count:    1,
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
