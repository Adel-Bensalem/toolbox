package adapters

type FileCreator interface {
	CreateFile(filename string) error
}
