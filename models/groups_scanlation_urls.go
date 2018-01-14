package models

type GroupsScanlationUrls struct {
	GroupName string `orm:"column(group_name)"`
	Url       string `orm:"column(url)"`
}
