package controllers

import (
	"encoding/json"
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"strconv"

	"github.com/astaxie/beego"
)

//  Series_user_ratingsController operations for Series_user_ratings
type Series_user_ratingsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Series_user_ratingsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
}

// Post ...
// @Title Post
// @Description create Series_user_ratings
// @Param	body		body 	models.Series_user_ratings	true		"body for Series_user_ratings content"
// @Success 201 {int} models.Series_user_ratings
// @Failure 403 body is empty
// @router / [post]
func (c *Series_user_ratingsController) Post() {
	user, err := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	var v models.SeriesUserRatings
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v.UserId = user
	if _, err := models.AddNewSeriesRating(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = Response{
			Success: true,
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Series_user_ratings by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Series_user_ratings
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Series_user_ratingsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAverageSeriesRating(id)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	} else {
		var temp []interface{}
		temp = append(temp, v)
		c.Data["json"] = Response{
			Success:  true,
			Response: temp,
			Count:    1,
		}
	}
	c.ServeJSON()
}
