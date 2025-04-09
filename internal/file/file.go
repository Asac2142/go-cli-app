// Package file handles file operations such as read & write.
package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const filename = "tasks.json"

// File - file generic struct. T generic.
type File[T any] struct{}

// New instantiates a File struct.
func New[T any]() *File[T] {
	return &File[T]{}
}

// Write - writes tasks into a JSON file
func (f *File[T]) Write(data *[]T) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal file: %w", err)
	}

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// Read - reads tasks from json file
func (f *File[T]) Read() ([]T, error) {
	bytes, err := os.ReadFile(filename)

	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if len(bytes) == 0 {
		return nil, nil
	}

	var data []T

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal file: %w", err)
	}

	return data, nil
}
