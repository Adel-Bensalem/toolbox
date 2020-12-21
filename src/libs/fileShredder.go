package libs

import "os"

type FileShredder struct{}

func (fileShredder *FileShredder) ShredFile(file string) error {
	return os.RemoveAll(file)
}
