package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesUserRatings struct {
	SeriesId *Series `orm:"column(series_id);rel(fk);pk"`
	UserId   *Users  `orm:"column(user_id);rel(fk)"`
	Rating   int     `orm:"column(rating);null"`
}

func (t *SeriesUserRatings) TableName() string {
	return "series_user_ratings"
}

func init() {
	orm.RegisterModel(new(SeriesUserRatings))
}

func GetAverageSeriesRating(series_id int64) (u SeriesUserRatings, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT avg(*) FROM series_ueser_ratings WHERE series_id=?", series_id).QueryRow(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func AddNewSeriesRating(m *SeriesUserRatings) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
