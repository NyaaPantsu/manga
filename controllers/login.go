package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
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
			c.Data["json"] = "{'error': 'Error invalid username'}"
			c.ServeJSON()
			return
		}
		password := []byte(v.Password)

		user, err := models.GetUserByUsername(v.Username)
		if err != nil {
			c.Data["json"] = "{'error': '" + err.Error() + "'}"
			c.ServeJSON()
			return

		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), password)
		if err != nil {
			c.Data["json"] = "{'error': '" + err.Error() + "'}"
			c.ServeJSON()
			return
		}
		et := jwtbeego.EasyToken{
			Username: user.Username,
			Expires:  time.Now().Unix() + 3600,
		}
		tokenString, _ := et.GetToken()
		c.Data["json"] = "{'tokenString': '" + tokenString + "', 'username':'" + user.Username + "'}"
		c.ServeJSON()
		return

	} else {
		c.Data["json"] = "{'error': '" + err.Error() + "'}"
	}
	c.ServeJSON()
	return
}
