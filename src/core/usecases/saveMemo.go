package usecases

import (
	"core/adapters"
	"core/entities"
	"core/types"
	"errors"
)

type MemoSaveInteractor func(memo types.Memo) error

func CreateMemoSaveInteractor(repository adapters.MemoRepository) MemoSaveInteractor {
	return func(memo types.Memo) error {
		if !entities.IsMemoValid(memo) {
			return errors.New("invalid memo provided")
		}

		return repository.SaveMemo(memo)
	}
}
