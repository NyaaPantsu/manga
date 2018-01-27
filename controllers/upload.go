package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"github.com/NyaaPantsu/manga/utils/zip"
	"github.com/dchest/uniuri"

	"errors"
	"io"
	"os"
	"time"
)

// UploadController operations for Upload
type UploadController struct {
	BaseController
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

type UploadForm struct {
	Id                    int     `form:"-"`
	Title                 string  `form:"title, text"`
	ChapterNumberAbsolute string  `form:"chapternum, te`
	ChapterNumberVolume   float64 `form:"chaptervol, number"`
	VolumeNumber          float64 `form:"volnum, number"`
	ChapterLanguage       string  `form:"languages, text"`
	ReleaseDelay          int     `form:"delay, number"`
	Groups1               string  `form:"group1, text"`
	Groups2               string  `form:"group2, text"`
	Groups3               string  `form:"group3, text"`
}

// Post ...
// @Title Create
// @Description create Upload
// @Param	body		body 	models.Upload	true		"body for Upload content"
// @Success 201 {object} models.Upload
// @Failure 403 body is empty
// @router / [post]
func (c *UploadController) Post() {

	u := UploadForm{}
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}
	series, err := models.GetSeriesByName(u.Title)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	user, _ := models.GetUserByUsername(auth.GetUsername(c.Ctx))

	random := uniuri.New()
	chapter := models.SeriesChapters{
		Title:                 u.Title,
		SeriesId:              &series,
		ContributorId:         user,
		ChapterLanguage:       &models.Languages{Name: u.ChapterLanguage},
		Hash:                  random,
		VolumeNumber:          u.VolumeNumber,
		ChapterNumberVolume:   u.ChapterNumberVolume,
		ChapterNumberAbsolute: u.ChapterNumberAbsolute,
		TimeUploaded:          time.Now(),
	}
	id, err := models.AddSeriesChapters(&chapter)

	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}

	chapters, err := models.GetSeriesChaptersById(int(id))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

		c.ServeJSON()
		return
	}

	//var img string
	files, err := c.GetFiles("files")
	for i := range files {

		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		//img = random + files[i].Filename
		//create destination file making sure the path is writeable.
		dst, err := os.Create("uploads/" + random + files[i].Filename)

		defer dst.Close()
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		err = os.Mkdir("uploads/"+random, os.ModePerm)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

		fpath := "uploads/" + random + files[i].Filename
		var images []string
		if zip.IsRar(fpath) {
			images, err = zip.Unrar(fpath, "uploads/"+random)
		} else {
			images, err = zip.Unzip(fpath, "uploads/"+random)
		}
		if len(images) == 0 {
			err := errors.New("error empty archive")
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return

		}
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}
		var temp []models.SeriesChaptersFiles
		for _, cond := range images {
			temp2 := models.SeriesChaptersFiles{
				ChapterId: chapters,
				Name:      cond,
			}
			temp = append(temp, temp2)

		}
		_, err = models.AddMultiChapterFiles(temp)
		if err != nil {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}

			c.ServeJSON()
			return
		}

	}
	c.ServeJSON()
	return

}
