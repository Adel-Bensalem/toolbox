package entities

import "core/types"

func IsMemoValid(memo types.Memo) bool {
	return len(memo.Title) > 0 && len(memo.Body) > 0
}
