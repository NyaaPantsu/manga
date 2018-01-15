package split

import (
	"github.com/NyaaPantsu/manga/models"

	"errors"
	"strings"
)

func ProcessTags(series *models.Series, tags string) (processed []*models.SeriesTags, err error) {
	var processedTags []*models.Tags
	for _, cond := range strings.Split(tags, ",") {
		kv := strings.SplitN(cond, ":", 2)
		if len(kv) != 2 {
			err := errors.New("Error: invalid query key/value pair")
			return nil, err
		}
		temp1 := models.Tags{
			Name:      kv[0],
			Namespace: kv[1],
		}
		temp := models.SeriesTags{
			SeriesId:     series,
			TagName:      kv[0],
			TagNamespace: kv[1],
		}
		processed = append(processed, &temp)
		processedTags = append(processedTags, &temp1)
	}
	models.AddMultiTags(processedTags)
	models.AddMultiSeriesTags(processed)
	return processed, nil
}

func CreateTag(name string, namespace string) (temp string) {

	temp = name + ":" + namespace
	return
}

func CreateTags(series *models.Series, name string, namespaces string) (tags []*models.SeriesTags, err error) {
	var temp []string
	for _, cond := range strings.Split(namespaces, ",") {
		temp = append(temp, CreateTag(name, cond))
	}
	tag := strings.Join(temp, ",")
	tags, err = ProcessTags(series, tag)

	return

}
