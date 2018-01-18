package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"

	"net/url"
	"strconv"
)

type FollowController struct {
	BaseController
}

func (c *FollowController) ToggleFollow() {
	var sid, uid int
	var err error
	r := "/"
	flash := beego.NewFlash()

	if v := c.GetString("r"); v != "" {
		r, err = url.QueryUnescape(v)
		if err != nil {
			r = "/"
		}
	}

	if !c.IsLogin {
		flash.Error("You must log in to follow your favorite series!")
		flash.Store(&c.Controller)
		c.Redirect(r, 302)
		return
	}

	uid = c.GetSession("userinfo").(int)
	sid, err = strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(r, 302)
		return
	}

	following := models.Following(uid, sid)
	if following {
		err = models.Unfollow(uid, sid)
		flash.Notice("You are no longer following this series.")
	} else {
		err = models.Follow(uid, sid)
		flash.Notice("You are now following this series!")
	}
	if err != nil {
		flash.Error(err.Error())
	}

	flash.Store(&c.Controller)
	c.Redirect(r, 302)
	return
}
