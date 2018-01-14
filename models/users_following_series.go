package models

type UsersFollowingSeries struct {
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
}
