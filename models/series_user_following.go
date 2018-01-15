package models

type SeriesUserFollowing struct {
	Id       int     `orm:"auto"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
}
