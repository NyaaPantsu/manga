package models

import (
	"github.com/astaxie/beego/orm"
)

type GroupsScanlationUrls struct {
	Id        int               `orm:"auto"`
	GroupName *GroupsScanlation `orm:"column(group_name);rel(fk)"`
	Url       string            `orm:"column(url)"`
}

func (t *GroupsScanlationUrls) TableName() string {
	return "groups_scanlation_urls"
}

func init() {
	orm.RegisterModel(new(GroupsScanlationUrls))
}

// AddMultiGroupUrl adds multiple urls to group name
func AddMultiGroupUrl(groups []GroupsScanlationUrls) (count int64, err error) {

	o := orm.NewOrm()
	count, err = o.InsertMulti(len(groups), &groups)
	return
}

func GetAllGroupUrlsByName(name string) (urls []GroupsScanlationUrls, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("groups").Filter("group_name", name).All(&urls)
	return

}
