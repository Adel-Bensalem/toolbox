package cli

type CommandLauncher struct {
	CommandMap  *CommandMap
	Interpreter *CommandInterpreter
}

func (launcher *CommandLauncher) Launch(args []string) error {
	interpretation := launcher.Interpreter.Interpret(args)
	command, err := launcher.CommandMap.Get(interpretation.name)

	if err != nil {
		return err
	}

	command.execute(interpretation.args, interpretation.options)

	return nil
}
