package usecases

import (
	"core/adapters"
	"core/types"
)

type MemosReadInteractor func() ([]types.Memo, error)

func CreateMemosReadInteractor(repository adapters.MemoRepository) MemosReadInteractor {
	return repository.GetMemos
}
