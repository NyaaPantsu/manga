package models

import (
	"github.com/astaxie/beego/orm"
)

type Tags struct {
	Id        int    `orm:auto`
	Name      string `orm:"column(name)"`
	Namespace string `orm:"column(namespace)"`
}

func (t *Tags) TableName() string {
	return "tags"
}

func init() {
	orm.RegisterModel(new(Tags))
}

// AddMultiSeriesTags adds multiple files to chapter ID
func AddMultiTags(tags []*Tags) (count int64, err error) {

	o := orm.NewOrm()
	count, err = o.InsertMulti(len(tags), tags)
	return
}
