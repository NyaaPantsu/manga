package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/microcosm-cc/bluemonday"

	"errors"
	"gopkg.in/russross/blackfriday.v2"
	"strings"
)

// Groups_addController operations for Groups_add
type Groups_addController struct {
	BaseController
}

// URLMapping ...
func (c *Groups_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)

	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type GroupsForm struct {
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	ReleaseDlay int    `form:"release_delay, int"`
	Urls        string `form:"urls, text"`
}

// Post ...
// @Title Create
// @Description create Groups
// @Param	body		body 	models.GroupsScanlation	true		"body for Groups_add content"
// @Success 201 {object} models.GroupsScanlation
// @Failure 403 body is empty
// @router / [post]
func (c *Groups_addController) Post() {
	_, err := models.GetUserByUsername(c.Claims())
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	u := GroupsForm{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}

	exists := models.GroupsScanlationNameExists(u.Name)
	if !exists {
		unsafe := blackfriday.Run([]byte(u.Description))
		html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		groups := models.GroupsScanlation{
			Name:        u.Name,
			Description: string(html),
		}
		// todo prevent xss
		urls := []models.GroupsScanlationUrls{}
		for _, cond := range strings.Split(u.Urls, ",") {
			temp := models.GroupsScanlationUrls{
				GroupName: &groups,
				Url:       cond,
			}
			urls = append(urls, temp)
		}

		if err := models.AddGroupsScanlation(&groups); err == nil {
			usergroups := models.UsersGroups{
				GroupName: u.Name,
			}
			if err = models.AddUserGroups(&usergroups); err == nil {

				var k []interface{}
				k = append(k, groups)
				k = append(k, usergroups)
				c.Data["json"] = Response{
					Success:  true,
					Response: k,
					Count:    1,
				}
				c.ServeJSON()
				return
			}

			if err != nil {
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}
		}

		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
			c.ServeJSON()
			return
		}

		return

	}
	err = errors.New("error unable to add new group")
	c.Data["json"] = Response{
		Success: false,
		Error:   err.Error(),
	}
	c.ServeJSON()
	return
}
