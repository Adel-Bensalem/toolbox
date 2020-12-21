package cli

import "strings"

type CommandInterpreter struct{}

type CommandInterpretation struct {
	name    string
	options map[string]string
	args    []string
}

func splitOption(option string) (string, string) {
	optionWithoutDash := option[1:]
	keyValuePair := strings.Split(optionWithoutDash, "=")

	return keyValuePair[0], keyValuePair[1]
}

func extractOptions(args []string) (map[string]string, []string) {
	options := make(map[string]string)
	argsRest := make([]string, len(args))

	for index, arg := range args {
		if strings.Index(arg, "-") != -1 {
			name, value := splitOption(args[index])
			options[name] = value
		} else {
			argsRest = append(argsRest, arg)
		}
	}

	return options, argsRest
}

func (interpreter *CommandInterpreter) Interpret(args []string) CommandInterpretation {
	name := args[0]
	argsWithoutName := args[1:]
	options, argsWithoutOptions := extractOptions(argsWithoutName)

	return CommandInterpretation{name: name, options: options, args: argsWithoutOptions}
}
