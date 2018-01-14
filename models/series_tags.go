package models

type SeriesTags struct {
	SeriesId     *Series `orm:"column(series_id);rel(fk)"`
	TagName      string  `orm:"column(tag_name)"`
	TagNamespace string  `orm:"column(tag_namespace)"`
}
