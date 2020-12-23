package libs

import (
	"core/types"
	"encoding/json"
	"io/ioutil"
	"os"
)

type CommandStack struct{}

func openHistory() (*os.File, error) {
	file, err := os.OpenFile("history.json", os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return &os.File{}, err
	}

	if infos, err := file.Stat(); err != nil {
		return &os.File{}, err
	} else if size := infos.Size(); size == 0 {
		if _, err := file.WriteString("[]"); err != nil {
			return &os.File{}, err
		}
	}

	return file, nil
}

func (stack *CommandStack) Push(command types.Command) error {
	file, err := openHistory()

	if err != nil {
		return err
	}

	defer file.Close()

	fileData, err := ioutil.ReadFile("history.json")

	if err != nil {
		return err
	}

	var fileContent []types.Command

	err = json.Unmarshal(fileData, &fileContent)

	if err != nil {
		return err
	}

	fileContent = append(fileContent, command)

	jsonData, err := json.Marshal(fileContent)

	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Write(jsonData); err != nil {
		return err
	}

	return nil
}
