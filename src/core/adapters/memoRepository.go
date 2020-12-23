package adapters

import "core/types"

type MemoRepository interface {
	SaveMemo(memo types.Memo) error
	GetMemos() ([]types.Memo, error)
}
