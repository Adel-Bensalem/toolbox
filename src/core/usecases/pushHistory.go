package usecases

import (
	"core/adapters"
	"core/types"
)

type HistoryPushInteractor func(command types.Command) error

func CreateHistoryPushInteractor(stack adapters.CommandStack) HistoryPushInteractor {
	return stack.Push
}
