package usecases

import (
	"core/adapters"
	"core/entities"
	"fmt"
)

type FilesCreationInteractor func(files []string) error

func createFile(file string, finder adapters.FileFinder, creator adapters.FileCreator) error {
	if !finder.FindFile(file) && entities.IsFileValid(file) {
		return creator.CreateFile(file)
	}

	return fmt.Errorf("file %s already exists", file)
}

func forEachFile(files []string, execute func(file string) error) error {
	for _, file := range files {
		if err := execute(file); err != nil {
			return err
		}
	}

	return nil
}

func CreateFilesCreationInteractor(finder adapters.FileFinder, fileCreator adapters.FileCreator) FilesCreationInteractor {
	return func(files []string) error {
		return forEachFile(files, func(file string) error {
			return createFile(file, finder, fileCreator)
		})
	}
}
