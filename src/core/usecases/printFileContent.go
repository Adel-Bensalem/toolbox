package usecases

import "core/adapters"

type FileContentPrintInteractor func(file string) error

func CreateFileContentPrintInteractor(fileReader adapters.FileReader, printer adapters.Printer) FileContentPrintInteractor {
	return func(file string) error {
		if fileContent, err := fileReader.ReadFile(file); err == nil {
			printer.Print(fileContent)
		} else {
			return err
		}

		return nil
	}
}
