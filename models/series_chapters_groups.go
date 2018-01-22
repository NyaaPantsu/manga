package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesChaptersGroups struct {
	Id        int             `orm:"auto"`
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk)"`
	GroupName string          `orm:"column(group_name)"`
}

func (t *SeriesChaptersGroups) TableName() string {
	return "series_chapters_groups"
}

func init() {
	orm.RegisterModel(new(SeriesChaptersGroups))
}
