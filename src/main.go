package main

import (
	"core"
	"fmt"
	"libs"
	"libs/cli"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("you need to provide commands")

		return
	}

	fileFinder := libs.FileFinder{}
	fileShredder := libs.FileShredder{}
	c := core.CreateCore(&fileFinder, &fileShredder)
	commandMap := cli.CommandMap{}
	commandInterpreter := cli.CommandInterpreter{}
	commandLauncher := cli.CommandLauncher{
		CommandMap:  &commandMap,
		Interpreter: &commandInterpreter,
	}

	commandMap.Add("rm", func(args []string, options map[string]string) {
		if err := c.DeleteFile(args); err != nil {
			fmt.Printf("command \"rm\" failed: %s", err)
		}
	})

	if err := commandLauncher.Launch(os.Args[1:]); err != nil {
		fmt.Printf("an unexpected error occured: %s", err)
	}
}
