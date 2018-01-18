package resize

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"gopkg.in/h2non/filetype.v1"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func ResizeImage(open string, filepath string) (err error) {
	buf, _ := ioutil.ReadFile(open)
	log := logs.GetLogger()
	file, err := os.Open(open)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer file.Close()

	if filetype.IsMIME(buf, "image/png") {

		// decode png into image.Image
		img, err := png.Decode(file)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		// resize to width 60 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Thumbnail(500, 500, img, resize.Lanczos3)

		out, err := os.Create(filepath)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)
		return nil
	}

	if filetype.IsMIME(buf, "image/jpeg") {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		// and preserve aspect ratio
		m := resize.Thumbnail(500, 500, img, resize.Lanczos3)

		out, err := os.Create(filepath)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		return nil
	}
	return errors.New("error must be png or jpg")
}
