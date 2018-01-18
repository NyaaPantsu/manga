package controllers

import (
	"time"
	"fmt"

	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/feeds"
)

// RssController operations for Rss
type RssController struct {
	BaseController
}

// URLMapping ...
func (c *RssController) URLMapping() {
	c.Mapping("GetNew", c.GetNew)
	c.Mapping("GetFollowed", c.GetFollowed)
}

// TODO: fix function documentation

// GetNew ...
// @Title GetNew
// @Description get recently updated series
// @Param id		path	string	true		"The key for staticblock"
// @Success 200 {object} models.Rss
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RssController) GetNew() {
	flash := beego.NewFlash()

	recents, err := models.GetRecentSeriesByChapter(0, 50)
	if err != nil {
		flash.Error("Recent comics not found")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	now := time.Now()
	feed := &feeds.Feed {
		Title: "New chapters",
		Link: &feeds.Link{Href: "https://manga.sh"},
		Description: "Recently updated manga",
		Created: now,
	}
	feed.Items = make([]*feeds.Item, 50)
	for index, element := range recents {
		feed.Items[index] = &feeds.Item{
			Title: fmt.Sprintf("%s - Chapter %d", element.SeriesId.Name, element.ChapterNumberAbsolute),
			Link: &feeds.Link{Href: fmt.Sprint("https://manga.sh/reader/", element.Hash)},
		}
	}
	switch c.GetString("format") {
	case "atom":
		atom, err := feed.ToAtom()
		if err != nil {
			flash.Error("Atom generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/atom+xml")
		c.Ctx.Output.Body([]byte(atom))
	case "json":
		json, err := feed.ToJSON()
		if err != nil {
			flash.Error("JSON generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/json")
		c.Ctx.Output.Body([]byte(json))
	default:
		rss, err := feed.ToRss()
		if err != nil {
			flash.Error("RSS generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/rss+xml")
		c.Ctx.Output.Body([]byte(rss))
	}
	return
}

// GetFollowed ...
// @Title GetFollowed
// @Description get recently updated series from followed series
// @Param query query string	false "Filter. e.g. col1:v1,col2:v2 ..."
// @Param fields	query string	false "Fields returned. e.g. col1,col2 ..."
// @Param sortby	query string	false "Sorted-by fields. e.g. col1,col2 ..."
// @Param order query string	false "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param limit query string	false "Limit the size of result set. Must be an integer"
// @Param offset	query string	false "Start position of result set. Must be an integer"
// @Success 200 {object} models.Rss
// @Failure 403
// @router / [get]
func (c *RssController) GetFollowed() {
	flash := beego.NewFlash()

	if !c.IsLogin {
		flash.Error("You must be logged in to view followed series")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}

	recents, err := models.GetRecentFollowedSeriesByChapter(c.GetSession("userinfo").(int), 0, 50)
	if err != nil {
		flash.Error("Followed comics not found")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	now := time.Now()
	feed := &feeds.Feed {
		Title: "New chapters",
		Link: &feeds.Link{Href: "https://manga.sh"},
		Description: "Recently updated followed manga",
		Created: now,
	}
	feed.Items = make([]*feeds.Item, 50)
	for index, element := range recents {
		feed.Items[index] = &feeds.Item{
			Title: fmt.Sprintf("%s - Chapter %d", element.SeriesId.Name, element.ChapterNumberAbsolute),
			Link: &feeds.Link{Href: fmt.Sprint("https://manga.sh/reader/", element.Hash)},
		}
	}
	switch c.GetString("format") {
	case "atom":
		atom, err := feed.ToAtom()
		if err != nil {
			flash.Error("Atom generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/atom+xml")
		c.Ctx.Output.Body([]byte(atom))
	case "json":
		json, err := feed.ToJSON()
		if err != nil {
			flash.Error("JSON generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/json")
		c.Ctx.Output.Body([]byte(json))
	default:
		rss, err := feed.ToRss()
		if err != nil {
			flash.Error("RSS generation failed")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		c.Ctx.Output.Header("Content-Type", "application/rss+xml")
		c.Ctx.Output.Body([]byte(rss))
	}
	return
}
