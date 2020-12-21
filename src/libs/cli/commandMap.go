package cli

import "errors"

type CommandMap struct {
	commandList []Command
}

func (commandMap *CommandMap) Add(name string, execute func(args []string, options map[string]string)) {
	commandMap.commandList = append(commandMap.commandList, Command{name: name, execute: execute})
}

func (commandMap *CommandMap) Get(name string) (*Command, error) {
	for _, command := range commandMap.commandList {
		if command.name == name {
			return &command, nil
		}
	}

	return nil, errors.New("command \"" + name + "\" not found")
}
