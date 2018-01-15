package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesChaptersFiles struct {
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk)"`
	Name      string          `orm:"column(name)"`
}

// AddMultiChapterFiles adds multiple files to chapter ID
func AddMultiChapterFiles(files []SeriesChaptersFiles) (count int64, err error) {

	o := orm.NewOrm()
	count, err = o.InsertMulti(len(files), &files)
	return
}

// GetAllChapterFilesById gets all files by id
func GetAllChapterFilesById(id int) (v []*SeriesChaptersFiles, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("series_chapters_files").Filter("chapter_id", id).All(&v)
	return
}
