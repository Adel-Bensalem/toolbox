package cli

type Command struct {
	name    string
	execute func(args []string, options map[string]string)
}
