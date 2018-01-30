package models

import (
	"github.com/astaxie/beego/orm"
)

type UsersFollowingSeries struct {
	UserId   *Users  `orm:"column(user_id);rel(fk);pk"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
}

func (t *UsersFollowingSeries) TableName() string {
	return "users_following_series"
}

func init() {
	orm.RegisterModel(new(UsersFollowingSeries))
}

//GetRecentFollowedSeriesByChapter Gets followed Series in order of chapter updates
func GetRecentFollowedSeriesByChapter(user_id int, offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	var series orm.ParamsList
	var b []SeriesChapters

	_, err = o.QueryTable("users_following_series").Filter("user_id", user_id).ValuesFlat(&series, "series_id")
	_, err = o.QueryTable("series_chapters").Filter("SeriesId__in", series...).OrderBy("time_uploaded").Limit(limit, offset).RelatedSel().All(&b)
	count, err = o.QueryTable("series_chapters").Filter("SeriesId__in", series...).OrderBy("time_uploaded").Count()
	for _, v := range b {
		ml = append(ml, v)
	}
	return
}

func Following(user_id, series_id int) (b bool) {
	var u UsersFollowingSeries
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM users_following_series WHERE user_id=? AND series_id=?", user_id, series_id).QueryRow(&u)
	if err != nil {
		b = false
	} else {
		b = true
	}
	return
}

func Follow(user_id, series_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO users_following_series VALUES (?, ?)", user_id, series_id).Exec()
	return
}

func Unfollow(user_id, series_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("DELETE FROM users_following_series WHERE user_id=? AND series_id=?", user_id, series_id).Exec()
	return
}
