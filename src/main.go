package main

import (
	"core"
	"fmt"
	"libs"
	"libs/cli"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("you need to provide commands")

		return
	}

	fileFinder := libs.FileFinder{}
	fileShredder := libs.FileShredder{}
	fileReader := libs.FileReader{}
	printer := libs.Printer{}
	requestClient := libs.RequestClient{}
	fileCreator := libs.FileCreator{}
	commandStack := libs.CommandStack{}
	history := libs.History{}
	c := core.CreateCore(
		&fileFinder,
		&fileShredder,
		&fileReader,
		&printer,
		&requestClient,
		&fileCreator,
		&commandStack,
		&history,
	)
	commandMap := cli.CommandMap{}
	commandInterpreter := cli.CommandInterpreter{}
	commandLauncher := cli.CommandLauncher{
		CommandMap:  &commandMap,
		Interpreter: &commandInterpreter,
	}
	registerCommand := func(commandName string, handleCommand func(args []string, options map[string]string)) {
		commandMap.Add(commandName, func(args []string, options map[string]string) {
			c.PushHistory(commandName, args, options)

			handleCommand(args, options)
		})
	}

	registerCommand("rm", func(args []string, options map[string]string) {
		if err := c.DeleteFile(args); err != nil {
			fmt.Printf("command \"rm\" failed: %s", err)
		}
	})

	registerCommand("cat", func(args []string, options map[string]string) {
		if err := c.PrintFileContent(args[0]); err != nil {
			fmt.Printf("command \"cat\" failed: %s", err)
		}
	})

	registerCommand("history", func(args []string, options map[string]string) {
		if len(args) == 0 {
			fmt.Printf("command \"history\" failed: history command requires an intent parameter")
			return
		}

		switch args[0] {
		case "ls":
			if commands, err := c.PrintHistory(); err != nil {
				fmt.Printf("command \"history ls\" failed: %s", err)
			} else {
				fmt.Println(strings.Join(commands, "\n"))
			}
		case "clear":
			if err := c.ClearHistory(); err != nil {
				fmt.Printf("command \"history clear\" failed: %s", err)
			}
		default:
			fmt.Printf("command \"history\" failed: history command requires an intent parameter")
		}
	})

	registerCommand("make", func(args []string, options map[string]string) {
		if err := c.CreateFiles(args); err != nil {
			fmt.Printf("command \"make\" failed: %s", err)
		}
	})

	registerCommand("send", func(args []string, options map[string]string) {
		if res, err := c.SendRequest(struct {
			Data   string
			Url    string
			Method string
		}{Data: options["data"], Url: options["url"], Method: options["method"]}); err != nil {
			fmt.Printf("command \"send\" failed: %s", err)
		} else {
			fmt.Println("status: ", res.Code)
			fmt.Println("body:\n", res.Body)
		}
	})

	if err := commandLauncher.Launch(os.Args[1:]); err != nil {
		fmt.Printf("an unexpected error occured: %s", err)
	}
}
