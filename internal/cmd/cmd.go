// Package cmd handles user interaction with the CLI.
package cmd

import (
	"errors"
	"fmt"
	"github/Asac2142/go-cli-app/internal/file"
	"github/Asac2142/go-cli-app/internal/task"
	"os"
	"strconv"
)

const (
	add            = "add"
	update         = "update"
	deleteTask     = "delete"
	markInProgress = "mark-in-progress"
	markDone       = "mark-done"
	list           = "list"
)

// HandleTrackerCLI handles the user's input regarding a task operation.
func HandleTrackerCLI(f *file.File[task.TContent]) error {
	var err error
	task := task.NewTask(f)
	args := os.Args[1:]

	if len(args) == 0 {
		err = errors.New("no arguments were provided")
		return err
	}

	switch args[0] {
	case add:
		err = addTask(args, task)
		return err
	case update:
		err = updateTask(args, task)
		return err
	case deleteTask:
		err = removeTask(args, task)
		return err
	case markInProgress:
		err = updateStatus(args, markInProgress, task)
		return err
	case markDone:
		err = updateStatus(args, markDone, task)
		return err
	case list:
		err = listTasks(args, task)
		return err
	default:
		err = errors.New("unrecognized command")
		return err
	}
}

func addTask(args []string, t *task.Task) error {
	if len(args) == 1 {
		return errors.New("new task requires a description")
	}

	desc := args[1]
	added, err := t.Add(desc)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", added.ID)

	return nil
}

func updateTask(args []string, t *task.Task) error {
	if len(args) != 3 {
		return errors.New("update task requires an id & description")
	}

	strID := args[1]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return errors.New("invalid id provided")
	}

	dsc := args[2]
	return t.Update(id, dsc)
}

func removeTask(args []string, t *task.Task) error {
	if len(args) != 2 {
		return errors.New("delete task requires an id")
	}

	asString := args[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid id provided")
	}

	return t.Delete(id)
}

func updateStatus(args []string, status string, t *task.Task) error {
	if len(args) != 2 {
		return errors.New("mark-in-progress task requires an id")
	}

	asString := args[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid task id provided")
	}

	var s task.Status

	if status == markInProgress {
		s = task.InProgress
	}

	if status == markDone {
		s = task.Done
	}

	return t.UpdateStatus(id, s)
}

func listTasks(args []string, t *task.Task) error {
	if len(args) == 1 {
		handlePrintTasks(nil, t)
		return nil
	}

	if len(args) > 2 {
		return errors.New("listing tasks only requires 2 arguments")
	}

	status := task.Status(args[1])

	switch status {
	case task.Done:
		return handlePrintTasks(&status, t)
	case task.InProgress:
		return handlePrintTasks(&status, t)
	case task.Todo:
		return handlePrintTasks(&status, t)
	default:
		return errors.New("invalid provided task status")
	}
}

func handlePrintTasks(s *task.Status, t *task.Task) error {
	tasks, err := t.GetByStatus(s)
	if err != nil {
		return fmt.Errorf("an error occured while printing tasks %w", err)
	}

	if len(tasks) == 0 {
		return errors.New("no tasks registered yet")
	}

	for _, v := range tasks {
		fmt.Printf("ID: %-12d DESCRIPTION: %-50s STATUS: %-12s CREATED AT: %-50v UPDATED AT: %-50v\n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
	}

	return nil
}
