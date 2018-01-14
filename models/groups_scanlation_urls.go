package models

import (
	"github.com/astaxie/beego/orm"
)

type GroupsScanlationUrls struct {
	GroupName string `orm:"column(group_name)"`
	Url       string `orm:"column(url)"`
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
