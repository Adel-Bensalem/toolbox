package adapters

type History interface {
	Get() ([]string, error)
	Clear() error
}
