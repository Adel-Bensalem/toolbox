package libs

import (
	"io/ioutil"
	"path/filepath"
)

type FileReader struct{}

func (fileReader *FileReader) ReadFile(file string) (string, error) {
	absPath, _ := filepath.Abs(file)

	if content, err := ioutil.ReadFile(filepath.ToSlash(absPath)); err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
