package models

import (
	"github.com/astaxie/beego/orm"
)

type Statuses struct {
	Name string `orm:"column(name);pk"`
}

func (t *Statuses) TableName() string {
	return "statuses"
}

func init() {
	orm.RegisterModel(new(Statuses))
}

// AddStatuses insert a new Statuses into database and returns
// last inserted Id on success.
func AddStatuses(m *Statuses) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAllStatuses gets all Types and returns an array on success
func GetAllStatuses() (ml []interface{}, err error) {
	var l []Statuses
	o := orm.NewOrm()
	_, err = o.QueryTable("statuses").All(&l)
	for _, v := range l {
		ml = append(ml, v)
	}
	return
}
