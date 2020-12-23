package adapters

type History interface {
	GetHistory() ([]string, error)
}
