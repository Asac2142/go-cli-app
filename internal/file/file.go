// Package file handles file operations such as read & write.
package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		return fmt.Errorf("marshalling file: %w", err)
	}

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

// Read - reads tasks from json file
func (f *File[T]) Read() ([]T, error) {
	file, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		if file, err = os.Create(filename); err != nil {
			return nil, fmt.Errorf("creating file: %w", err)
		}
	} else {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	defer closer(file, &err, "close file")

	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	if len(bytes) == 0 {
		//return make([]T, 0), nil
		return nil, nil
	}

	var data []T

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling file: %w", err)
	}

	return data, nil
}

func closer(c io.Closer, errp *error, msg string) {
	err := c.Close()
	if *errp == nil && err != nil {
		*errp = fmt.Errorf("%v: %w", msg, err)
	}
}
