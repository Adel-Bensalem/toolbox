package adapters

import "core/types"

type MemoRepository interface {
	SaveMemo(title string, body string) error
	GetMemos() ([]types.Memo, error)
}
