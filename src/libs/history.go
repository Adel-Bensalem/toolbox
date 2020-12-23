package libs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type History struct{}

func (history *History) Get() ([]string, error) {
	var list []string
	fileData, err := ioutil.ReadFile("history.json")

	if err != nil {
		return make([]string, 0), err
	}

	var commands []Command

	err = json.Unmarshal(fileData, &commands)

	for _, command := range commands {
		list = append(list, command.Name)
	}

	return list, nil
}

func (history *History) Clear() error {
	file, err := os.OpenFile("history.json", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.WriteString("[]"); err != nil {
		return err
	}
	return nil
}
