package libs

import (
	"os"
	"path/filepath"
)

type FileFinder struct{}

func (fileFinder *FileFinder) FindFile(file string) bool {
	absPath, _ := filepath.Abs(file)
	openedFile, err := os.Open(filepath.ToSlash(absPath))

	if err != nil {
		return false
	}

	_ = openedFile.Close()

	return true
}
