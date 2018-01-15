package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SeriesChapters struct {
	Id                    int        `orm:"column(id);pk"`
	SeriesId              *Series    `orm:"column(series_id);rel(fk)"`
	Title                 string     `orm:"column(title)"`
	ChapterNumberAbsolute int        `orm:"column(chapter_number_absolute)"`
	ChapterNumberVolume   int        `orm:"column(chapter_number_volume);null"`
	VolumeNumber          int        `orm:"column(volume_number);null"`
	ChapterLanguage       *Languages `orm:"column(chapter_language);rel(fk)"`
	ContributorId         *Users     `orm:"column(contributor_id);rel(fk)"`
	TimeUploaded          time.Time  `orm:"column(time_uploaded);type(timestamp without time zone);auto_now_add"`
	Hash                  string     `orm:"column(hash)"`
}

func (t *SeriesChapters) TableName() string {
	return "series_chapters"
}

func init() {
	orm.RegisterModel(new(SeriesChapters))
}

// GetSeriesChaptersBySeriesId retrieves SeriesChapters by Id. Returns error if
// Id doesn't exist
func GetSeriesChaptersBySeriesId(id int) (v *[]SeriesChapters, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("series_chapters").Filter("series_id", id).All(&v)
	return

}

// AddSeriesChapters insert a new SeriesChapters into database and returns
// last inserted Id on success.
func AddSeriesChapters(m *SeriesChapters) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSeriesChaptersById retrieves SeriesChapters by Id. Returns error if
// Id doesn't exist
func GetSeriesChaptersById(id int) (v *SeriesChapters, err error) {
	o := orm.NewOrm()
	v = &SeriesChapters{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSeriesChapters retrieves all SeriesChapters matches certain condition. Returns empty list if
// no records exist
func GetAllSeriesChapters(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SeriesChapters))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SeriesChapters
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSeriesChapters updates SeriesChapters by Id and returns error if
// the record to be updated doesn't exist
func UpdateSeriesChaptersById(m *SeriesChapters) (err error) {
	o := orm.NewOrm()
	v := SeriesChapters{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSeriesChapters deletes SeriesChapters by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSeriesChapters(id int) (err error) {
	o := orm.NewOrm()
	v := SeriesChapters{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SeriesChapters{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
