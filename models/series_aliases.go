package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesAliases struct {
	Id       int     `orm:"auto"`
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	Name     string  `orm:"column(name)"`
}

func (t *SeriesAliases) TableName() string {
	return "series_aliases"
}

func init() {
	orm.RegisterModel(new(SeriesAliases))
}

// AddSeriesAlias insert a new Series into database and returns
// last inserted Id on success.
func AddSeriesAlias(m *SeriesAliases) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//  GetSeriesAliases checks to see if username exists
// returns bool
func GetAllSeriesAliases(id int) (aliases []*SeriesAliases, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("series_aliases").Filter("series_id", id).All(&aliases)
	return
}
