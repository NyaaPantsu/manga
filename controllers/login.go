package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"errors"
	"time"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title login
// @Description create Login
// @Param   body        body    models.Login   true        "The object content"
// @Success 201 {string}
// @Failure 403 {string}
// @router / [post]
func (c *LoginController) Post() {
	var v models.Login

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		username := models.UsernameExists(v.Username)
		if !username {
			err := errors.New("Error invalid username")
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return
		}
		password := []byte(v.Password)

		user, err := models.GetUserByUsername(v.Username)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), password)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return
		}
		et := jwtbeego.EasyToken{
			Username: v.Username,
			Expires:  time.Now().Unix() + 3600,
		}
		tokenString, _ := et.GetToken()
		type Auth struct {
			Token    string `json:"token"`
			Username string `json:"username"`
		}
		auth := Auth{
			Token:    tokenString,
			Username: v.Username,
		}
		var temp []interface{}
		temp = append(temp, auth)
		c.Data["json"] = Response{
			Success:  true,
			Response: temp,
			Count:    1,
		}

		c.ServeJSON()
		return

	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
	return
}
