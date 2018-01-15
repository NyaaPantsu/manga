package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesTags struct {
	SeriesId     *Series `orm:"column(series_id);rel(fk);pk"`
	TagName      string  `orm:"column(tag_name)"`
	TagNamespace string  `orm:"column(tag_namespace)"`
}

func (t *SeriesTags) TableName() string {
	return "series_tags"
}

func init() {
	orm.RegisterModel(new(SeriesTags))
}

// AddMultiSeriesTags adds multiple files to chapter ID
func AddMultiSeriesTags(tags []*SeriesTags) (count int64, err error) {

	o := orm.NewOrm()
	count, err = o.InsertMulti(len(tags), tags)
	return
}

// GetAllTagsById gets all tag by id
func GetAllTagsById(id int) (v []*SeriesTags, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("series_tags").Filter("series_id", id).All(&v)
	return
}
