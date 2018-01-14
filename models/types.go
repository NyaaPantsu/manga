package models

type Types struct {
	Name          string `orm:"column(name)"`
	OriginDemonym string `orm:"column(origin_demonym)"`
}
