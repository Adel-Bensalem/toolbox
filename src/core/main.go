package core

import (
	"core/adapters"
	"core/usecases"
)

type Core struct {
	DeleteFile usecases.FilesDeletionInteractor
}

func CreateCore(fileFinder adapters.FileFinder, fileShredder adapters.FileShredder) Core {
	return Core{
		DeleteFile: usecases.CreateFilesDeletionInteractor(fileFinder, fileShredder),
	}
}
