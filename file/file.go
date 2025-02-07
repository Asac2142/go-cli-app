// Package file - Handle file operations
package file

import (
	"encoding/json"
	"errors"
	"os"
)

const filename = "tasks.json"

// File - file generic struct. T generic.
type File[T any] struct{}

// New - instantiates a File struct.
func New[T any]() *File[T] {
	return &File[T]{}
}

// Write - writes tasks into a JSON file
func (f *File[T]) Write(data *[]T) error {
	asJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, asJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read - reads tasks from json file
func (f *File[T]) Read() ([]T, error) {
	jsonFile, err := os.Open(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			jsonFile, err = os.Create(filename)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	defer jsonFile.Close()

	bytes, err := os.ReadFile(jsonFile.Name())
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return make([]T, 0), nil
	}

	var data []T

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
