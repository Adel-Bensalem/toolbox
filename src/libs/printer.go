package libs

import "fmt"

type Printer struct{}

func (printer *Printer) Print(str string) {
	fmt.Println(str)
}
