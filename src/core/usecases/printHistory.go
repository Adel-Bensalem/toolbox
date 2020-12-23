package usecases

import "core/adapters"

type HistoryPrintInteractor func() ([]string, error)

func CreateHistoryPrintInteractor(history adapters.History) HistoryPrintInteractor {
	return history.Get
}
