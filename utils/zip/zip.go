package zip

import (
	"archive/zip"
	"os"
	"path/filepath"
)

// Unzip will un-compress a zip archive,
// moving all files and folders to an output directory
func Unzip(src, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {
			err = MakeFile(fpath, rc, f.Mode())
			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}
