package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"

	"encoding/json"
	"strconv"
)

// UsersController operations for Users
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	username := auth.GetUsername(c.Ctx)
	user, _ := models.GetUserByUsername(username)

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if user.Id != id {
		c.Data["json"] = "{'error':'cannot edit another user'}"
		c.ServeJSON()
		return
	}

	v := models.Users{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUsersById(&v); err == nil {
			c.Data["json"] = "{'success':'OK'}"
		} else {
			c.Data["json"] = "{'error':'" + err.Error() + "'}"
		}
	} else {

		c.Data["json"] = "{'error':'" + err.Error() + "'}"
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	username := auth.GetUsername(c.Ctx)
	user, _ := models.GetUserByUsername(username)

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if user.Id != id {
		c.Data["json"] = "{'error':'cannot delete another user'}"
		c.ServeJSON()
		return
	}
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "{'success':'OK'}"
	} else {
		c.Data["json"] = "{'error':'" + err.Error() + "'}"
	}
	c.ServeJSON()
}
