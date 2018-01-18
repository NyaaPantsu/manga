package models

import (
	"github.com/astaxie/beego/orm"
)

type UsersFollowingSeries struct {
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
}

//GetRecentFollowedSeriesByChapter Gets followed Series in order of chapter updates
func GetRecentFollowedSeriesByChapter(user_id int, offset int, limit int) (b []SeriesChapters, err error) {
	o := orm.NewOrm()
	var series orm.ParamsList
	_, err = o.QueryTable("users_following_series").Filter("user_id", user_id).ValuesFlat(&series, "series_id")
	_, err = o.QueryTable("series_chapters").Filter("series_chapters__series_id__in", series...).OrderBy("time_uploaded").Limit(limit, offset).RelatedSel().All(&b)
  return
}
