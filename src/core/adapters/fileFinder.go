package adapters

type FileFinder interface {
	FindFile(file string) bool
}
