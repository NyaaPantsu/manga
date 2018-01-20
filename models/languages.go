package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Languages struct {
	Name string `orm:"column(name);pk"`
	Code string `orm:"column(code)"`
}

func (t *Languages) TableName() string {
	return "languages"
}

func init() {
	orm.RegisterModel(new(Languages))
}

// AddLanguages insert a new Languages into database and returns
// last inserted Id on success.
func AddLanguages(m *Languages) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAllLanguages gets all languages and returns an array on success
func GetAllLanguages() (languages []*Languages, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("languages").OrderBy("name").All(&languages)
	return
}

// DeleteLanguages deletes Languages by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLanguages(name string) (err error) {
	o := orm.NewOrm()
	v := Languages{Name: name}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Languages{Name: name}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
