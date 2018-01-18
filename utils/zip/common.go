package zip

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// Make File
func MakeFile(fpath string, rc io.Reader, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	f, err := os.OpenFile(
		fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, rc)
	if err != nil {
		return err
	}

	return nil
}
