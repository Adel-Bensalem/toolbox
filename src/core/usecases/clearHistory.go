package usecases

import "core/adapters"

type HistoryClearInteractor func() error

func CreateHistoryClearInteractor(history adapters.History) HistoryClearInteractor {
	return history.Clear
}
