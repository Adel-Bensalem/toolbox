package libs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MemoRepository struct{}

type Memo struct {
	Title string
	Body  string
}

func openMemo() (*os.File, error) {
	file, err := os.OpenFile("memo.json", os.O_CREATE|os.O_RDWR, 0644)

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

func (repository *MemoRepository) SaveMemo(title string, body string) error {
	file, err := openMemo()

	if err != nil {
		return err
	}

	defer file.Close()

	fileData, err := ioutil.ReadFile("memo.json")

	if err != nil {
		return err
	}

	var fileContent []Memo

	err = json.Unmarshal(fileData, &fileContent)

	if err != nil {
		return err
	}

	fileContent = append(fileContent, Memo{
		Title: title,
		Body:  body,
	})

	jsonData, err := json.Marshal(fileContent)

	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.WriteAt(jsonData, 0); err != nil {
		return err
	}

	return nil
}
