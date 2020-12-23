package usecases

import (
	"core/adapters"
	"core/entities"
	"errors"
)

type MemoSaveInteractor func(title string, body string) error

func CreateMemoSaveInteractor(repository adapters.MemoRepository) MemoSaveInteractor {
	return func(title string, body string) error {
		if !entities.IsMemoValid(title, body) {
			return errors.New("invalid memo provided")
		}

		return repository.SaveMemo(title, body)
	}
}
