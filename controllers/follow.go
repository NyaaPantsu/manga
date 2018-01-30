package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"github.com/astaxie/beego"

	"strconv"
)

type FollowController struct {
	beego.Controller
}

func (c *FollowController) Get() {
	var limit int64 = 25
	var offset int64

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	user, err := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	l, count, err := models.GetRecentFollowedSeriesByChapter(user.Id, offset, limit)
	c.Data["json"] = Response{
		Success:  true,
		Response: l,
		Count:    count,
	}
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
	return

}

func (c *FollowController) ToggleFollow() {
	var sid, uid int
	var err error

	user, err := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}

	uid = user.Id
	sid, err = strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}

	following := models.Following(uid, sid)
	if following {
		err = models.Unfollow(uid, sid)
		c.Data["json"] = Response{
			Success: true,
		}
	} else {
		err = models.Follow(uid, sid)
		c.Data["json"] = Response{
			Success: true,
		}
	}
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
	return

}
