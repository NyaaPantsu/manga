package zip

import (
	"github.com/nwaples/rardecode"

	"io"
	"os"
	"path/filepath"
)

// Returns 1 if the file named src has RAR magic bytes 
func IsRar(src string) (bool) {
	f, err := os.Open(src)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	_, err = f.Read(buf)
	if err != nil {
		return false
	}

	return string(buf) == "Rar!"
}

// Unrar will un-compress a rar archive,
// moving all files and folders to an output directory
func Unrar(src, dest string) ([]string, error) {
	var filenames []string

	rf, err := os.Open(src)
	if err != nil {
		return filenames, err
	}
	defer rf.Close()

	r, err := rardecode.NewReader(rf, "")
	if err != nil {
		return filenames, err
	}

	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return filenames, err
		}

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, header.Name)
		filenames = append(filenames, fpath)

		if header.IsDir {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			err = MakeFile(fpath, r, header.Mode())
			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}
