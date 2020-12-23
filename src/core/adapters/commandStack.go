package adapters

import "core/types"

type CommandStack interface {
	Push(command types.Command) error
}
