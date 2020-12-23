package core

import (
	"core/adapters"
	"core/usecases"
)

type Core struct {
	DeleteFile       usecases.FilesDeletionInteractor
	PrintFileContent usecases.FileContentPrintInteractor
	SendRequest      usecases.RequestSendInteractor
	CreateFiles      usecases.FilesCreationInteractor
	PushHistory      usecases.HistoryPushInteractor
}

func CreateCore(
	fileFinder adapters.FileFinder,
	fileShredder adapters.FileShredder,
	fileReader adapters.FileReader,
	printer adapters.Printer,
	requestClient adapters.RequestClient,
	fileCreator adapters.FileCreator,
	commandStack adapters.CommandStack,
) Core {
	return Core{
		DeleteFile:       usecases.CreateFilesDeletionInteractor(fileFinder, fileShredder),
		PrintFileContent: usecases.CreateFileContentPrintInteractor(fileReader, printer),
		SendRequest:      usecases.CreateRequestSendInteractor(requestClient),
		CreateFiles:      usecases.CreateFilesCreationInteractor(fileFinder, fileCreator),
		PushHistory:      usecases.CreateHistoryPushInteractor(commandStack),
	}
}
