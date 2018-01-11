package models

type SeriesTag struct {
	SeriesId *Series `orm:"column(series_id);rel(fk)"`
	TagId    *Tag    `orm:"column(tag_id);rel(fk)"`
}
