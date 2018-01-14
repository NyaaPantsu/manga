package models

type SeriesAliases struct {
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	Name     string  `orm:"column(name)"`
}
