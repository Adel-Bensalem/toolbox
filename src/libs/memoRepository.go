package libs

import (
	"core/types"
	"encoding/json"
	"io/ioutil"
	"os"
)

type MemoRepository struct{}

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

func (repository *MemoRepository) SaveMemo(memo types.Memo) error {
	file, err := openMemo()

	if err != nil {
		return err
	}

	defer file.Close()

	fileData, err := ioutil.ReadFile("memo.json")

	if err != nil {
		return err
	}

	var fileContent []types.Memo

	err = json.Unmarshal(fileData, &fileContent)

	if err != nil {
		return err
	}

	fileContent = append(fileContent, memo)

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

func (repository *MemoRepository) GetMemos() ([]types.Memo, error) {
	file, err := openMemo()

	if err != nil {
		return make([]types.Memo, 0), err
	}

	defer file.Close()

	fileData, err := ioutil.ReadFile("memo.json")

	if err != nil {
		return make([]types.Memo, 0), err
	}

	var fileContent []types.Memo

	err = json.Unmarshal(fileData, &fileContent)

	if err != nil {
		return make([]types.Memo, 0), err
	}

	return fileContent, nil
}
