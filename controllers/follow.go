package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"

	"strconv"
)

type FollowController struct {
	BaseController
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
