package models

import (
	"github.com/astaxie/beego/orm"
)

type SeriesChaptersFiles struct {
	ChapterId *SeriesChapters `orm:"column(chapter_id);rel(fk); pk"`
	Name      string          `orm:"column(name)"`
}

func (t *SeriesChaptersFiles) TableName() string {
	return "series_chapters_files"
}

func init() {
	orm.RegisterModel(new(SeriesChaptersFiles))
}

// AddSeries insert a new Series into database and returns
// last inserted Id on success.
func AddSeriesChapterFiles(m *SeriesChaptersFiles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// AddMultiChapterFiles adds multiple files to chapter ID
func AddMultiChapterFiles(files []SeriesChaptersFiles) (count int64, err error) {

	o := orm.NewOrm()
	count, err = o.InsertMulti(len(files), &files)
	return
}

// GetAllChapterFilesById gets all files by id
func GetAllChapterFilesById(id int, limit int64, offset int64) (v []*SeriesChaptersFiles, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("series_chapters_files").Filter("chapter_id", id).Limit(int(limit), int(offset)).OrderBy("name").All(&v)
	return
}

// GetSeriesChaptersByHash retrieves SeriesChapters by Id. Returns error if
// Id doesn't exist
func GetSeriesChaptersFilesCount(id int) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("series_chapters_files").Filter("chapter_id", id).Count()
	return
}
