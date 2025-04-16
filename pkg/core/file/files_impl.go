package file

import (
	"encoding/json"
	"os"
)

type FilesImpl[T any] struct{}

func New[T any]() FilesApi[T] {
	return &FilesImpl[T]{}
}

func (file *FilesImpl[T]) ReadJsonFile(path string, model *T) error {
	bytesArray, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	unmarshalErr := json.Unmarshal(bytesArray, model)
	return unmarshalErr
}

func (file *FilesImpl[T]) WriteJsonFile(path string, model *T) error {
	byteArray, err := json.Marshal(&model)
	if err != nil {
		return err
	}
	writeErr := os.WriteFile(path, byteArray, 0644)
	return writeErr
}

func (file *FilesImpl[T]) ReadFolders(path string) ([]string, error) {
	dirEntry, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	var foldersArray []string
	for _, item := range dirEntry {
		foldersArray = append(foldersArray, path+"/"+item.Name())
	}
	return foldersArray, nil
}
