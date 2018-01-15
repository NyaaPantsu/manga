package models

import (
	"github.com/astaxie/beego/orm"
)

type Types struct {
	ID            int    `orm:pk`
	Name          string `orm:"column(name)"`
	OriginDemonym string `orm:"column(origin_demonym)"`
}

func (t *Types) TableName() string {
	return "types"
}

func init() {
	orm.RegisterModel(new(Types))
}

// AddTypes insert a new Types into database and returns
// last inserted Id on success.
func AddTypes(m *Types) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAllTypes gets all Types and returns an array on success
func GetAllTypes() (types []*Types, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("types").All(&types)
	return
}
