package usecases

import (
	"core/adapters"
	"fmt"
)

type FilesDeletionInteractor func(files []string) error

func forEach(files []string, execute func(file string) error) error {
	for _, file := range files {
		if err := execute(file); err != nil {
			return err
		}
	}

	return nil
}

func createFileNotFoundError(file string) error {
	return fmt.Errorf("file %s was not found", file)
}

func findFiles(files []string, fileFinder adapters.FileFinder) error {
	for _, file := range files {
		if !fileFinder.FindFile(file) {
			return createFileNotFoundError(file)
		}
	}

	return nil
}

func CreateFilesDeletionInteractor(fileFinder adapters.FileFinder, fileShredder adapters.FileShredder) FilesDeletionInteractor {
	return func(files []string) error {
		err := findFiles(files, fileFinder)

		if err != nil {
			return err
		}

		return forEach(files, fileShredder.ShredFile)
	}
}
