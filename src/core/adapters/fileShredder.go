package adapters

type FileShredder interface {
	ShredFile(file string) error
}
