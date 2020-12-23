package libs

import (
	"encoding/json"
	"io/ioutil"
)

type History struct{}

func (history *History) GetHistory() ([]string, error) {
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
