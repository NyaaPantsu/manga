package models

import (
	"github.com/astaxie/beego/orm"
)

type UsersGroups struct {
	Id        int    `orm:"auto"`
	UserId    *Users `orm:"column(user_id);rel(fk)"`
	GroupName string `orm:"column(group_name)"`
}

func (t *UsersGroups) TableName() string {
	return "users_groups"
}
func init() {
	orm.RegisterModel(new(UsersGroups))
}

// AddUserGroups insert a new UserGroups into database and returns
// last inserted Id on success.
func AddUserGroups(m *UsersGroups) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(m)
	return
}
