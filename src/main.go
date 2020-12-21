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
	fileReader := libs.FileReader{}
	printer := libs.Printer{}
	requestClient := libs.RequestClient{}
	c := core.CreateCore(&fileFinder, &fileShredder, &fileReader, &printer, &requestClient)
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

	commandMap.Add("cat", func(args []string, options map[string]string) {
		if err := c.PrintFileContent(args[0]); err != nil {
			fmt.Printf("command \"cat\" failed: %s", err)
		}
	})

	commandMap.Add("send", func(args []string, options map[string]string) {
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
