package resize

import (
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"os"
)

func ResizePng(open string, filepath string) {
	log := logs.GetLogger()
	file, err := os.Open(open)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// resize to width 60 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(200, 0, img, resize.Lanczos3)

	out, err := os.Create(filepath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

func ResizeJpg(open string, filepath string) {
	log := logs.GetLogger()
	file, err := os.Open(open)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println(err.Error())
		return
	}
	file.Close()

	// and preserve aspect ratio
	m := resize.Thumbnail(200, 0, img, resize.Lanczos3)

	out, err := os.Create(filepath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
