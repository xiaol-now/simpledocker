package container

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
)

func ImportImage(fpath string, dstPath string) error {
	if !ExistFile(fpath) {
		return errors.New("file not exist: " + fpath)
	}
	src, err := os.Open(fpath)
	if err != nil {
		return err
	}
	imageFile := path.Join(dstPath, filepath.Base(fpath))
	dst, err := os.OpenFile(imageFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer src.Close()
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}
