package core

import (
	"core/adapters"
	"core/usecases"
)

type Core struct {
	DeleteFile       usecases.FilesDeletionInteractor
	PrintFileContent usecases.FileContentPrintInteractor
	SendRequest      usecases.RequestSendInteractor
}

func CreateCore(
	fileFinder adapters.FileFinder,
	fileShredder adapters.FileShredder,
	fileReader adapters.FileReader,
	printer adapters.Printer,
	requestClient adapters.RequestClient,
) Core {
	return Core{
		DeleteFile:       usecases.CreateFilesDeletionInteractor(fileFinder, fileShredder),
		PrintFileContent: usecases.CreateFileContentPrintInteractor(fileReader, printer),
		SendRequest:      usecases.CreateRequestSendInteractor(requestClient),
	}
}
