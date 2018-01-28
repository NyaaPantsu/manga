package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesRatings struct {
	SeriesId *Series `orm:"column(series_id);rel(fk);pk"`
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	Rating   int     `orm:"column(rating)"`
}

func (t *SeriesRatings) TableName() string {
	return "series_ratings"
}

func init() {
	orm.RegisterModel(new(SeriesRatings))
}
