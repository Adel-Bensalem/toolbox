package adapters

type CommandStack interface {
	Push(name string, args []string, options map[string]string) error
}
