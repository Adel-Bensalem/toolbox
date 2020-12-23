package usecases

import "core/adapters"

type HistoryPushInteractor func(name string, args []string, options map[string]string) error

func CreateHistoryPushInteractor(stack adapters.CommandStack) HistoryPushInteractor {
	return stack.Push
}
