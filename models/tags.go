package models

type Tags struct {
	Name      string `orm:"column(name)"`
	Namespace string `orm:"column(namespace)"`
}
