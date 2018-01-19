package models

type SeriesUserRatings struct {
	Id       int     `orm:"auto"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	Rating   int     `orm:"column(rating);null"`
}
