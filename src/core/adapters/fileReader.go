package adapters

type FileReader interface {
	ReadFile(file string) (string, error)
}
