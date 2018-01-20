package controllers

import (
	"github.com/NyaaPantsu/manga/models"

	"strconv"
)

type FollowController struct {
	BaseController
}

func (c *FollowController) ToggleFollow() {
	var sid, uid int
	var err error

	user, _ := models.GetUserByUsername(c.Claims())

	uid = user.Id
	sid, err = strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = "{'error':'" + err.Error() + "'}"
		c.ServeJSON()
		return
	}

	following := models.Following(uid, sid)
	if following {
		err = models.Unfollow(uid, sid)
		c.Data["json"] = "{'success': 'You are no longer following this series.'}"
	} else {
		err = models.Follow(uid, sid)
		c.Data["json"] = "{'success':'You are now following this series!'}"
	}
	if err != nil {
		c.Data["json"] = "{'error':'" + err.Error() + "'}"
	}
	c.ServeJSON()
	return

}
