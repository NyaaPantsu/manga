package models

type SeriesChaptersFiles struct {
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk)"`
	Name      string          `orm:"column(name)"`
}
