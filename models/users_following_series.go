package models

type UsersFollowingSeries struct {
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
}

//GetRecentFollowedSeriesByChapter Gets followed Series in order of chapter updates
func GetRecentFollowedSeriesByChapter(userId UserId, offset int, limit int) (b []SeriesChapter, err error) {
  o := orm.NewOrm()
  var series orm.ParamsList
  _, err = o.QueryTable("users_following_series").Filter("user_id", userId).ValuesFlat(&series, "series_id")
  _, err = o.QueryTable("series_chapters").Filter("series_chapters__series_id__in", series...).OrderBy("time_uploaded").Limit(limit, offset).RelatedSel().All(&b)
}
