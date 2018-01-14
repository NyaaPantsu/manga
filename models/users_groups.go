package models

type UsersGroups struct {
	UserId    *Users `orm:"column(user_id);rel(fk)"`
	GroupName string `orm:"column(group_name)"`
}
