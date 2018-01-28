package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesChaptersGroups struct {
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk);pk"`
	GroupName string          `orm:"column(group_name)"`
}

func (t *SeriesChaptersGroups) TableName() string {
	return "series_chapters_groups"
}

func init() {
	orm.RegisterModel(new(SeriesChaptersGroups))
}

// AddSeriesChaptersGroups insert a new SeriesChaptersGroups into database and returns
// last inserted Id on success.
func AddSeriesChaptersGroups(m *SeriesChaptersGroups) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
