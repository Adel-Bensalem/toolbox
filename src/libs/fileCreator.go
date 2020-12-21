package libs

import "os"

type FileCreator struct{}

func (fileCreator *FileCreator) CreateFile(file string) error {
	if openFile, err := os.OpenFile(file, os.O_CREATE, 0777); err != nil {
		return err
	} else {
		openFile.Close()
	}

	return nil
}
