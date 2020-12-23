package types

type Command struct {
	Name    string
	Args    []string
	Options map[string]string
}
