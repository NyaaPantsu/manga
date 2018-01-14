package models

type SeriesChaptersGroups struct {
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk)"`
	GroupName string          `orm:"column(group_name)"`
}
