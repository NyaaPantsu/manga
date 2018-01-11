package models

type SeriesUserFollowing struct {
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
}
