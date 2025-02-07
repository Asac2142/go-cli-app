// Package cmd - package
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github/Asac2142/go-cli-app/file"
	"github/Asac2142/go-cli-app/task"
)

const (
	add            string = "add"
	update         string = "update"
	deleteTask     string = "delete"
	markInProgress string = "mark-in-progress"
	markDone       string = "mark-done"
	list           string = "list"
)

// HandleTrackerCLI - Handles user's input
func HandleTrackerCLI(f *file.File[task.TContent]) {
	task := task.NewTask(f)
	args := os.Args[1:]

	if len(args) == 0 {
		err := errors.New("Invalid input")
		printError(err)
	}

	switch args[0] {
	case add:
		err := addTask(&args, task)
		printError(err)
	case update:
		err := updateTask(&args, task)
		printError(err)
	case deleteTask:
		err := removeTask(&args, task)
		printError(err)
	case markInProgress:
		err := updateStatus(&args, markInProgress, task)
		printError(err)
	case markDone:
		err := updateStatus(&args, markDone, task)
		printError(err)
	case list:
		err := listTasks(&args, task)
		printError(err)
	default:
		err := errors.New("Invalid task command")
		printError(err)
	}
}

func addTask(args *[]string, t *task.Task) error {
	if len(*args) == 1 {
		return errors.New("add task requires a description")
	}

	desc := (*args)[1]
	added, err := t.Add(desc)

	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", added.ID)

	return nil
}

func updateTask(args *[]string, t *task.Task) error {
	if len(*args) != 3 {
		return errors.New("update task requires an id & description")
	}

	strID := (*args)[1]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return errors.New("invalid id provided")
	}

	dsc := (*args)[2]
	return t.Update(id, dsc)
}

func removeTask(args *[]string, t *task.Task) error {
	if len(*args) != 2 {
		return errors.New("delete task requires an id")
	}

	asString := (*args)[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid id provided")
	}

	return t.Delete(id)
}

func updateStatus(args *[]string, status string, t *task.Task) error {
	if len(*args) != 2 {
		return errors.New("mark-in-progress task requires an id")
	}

	asString := (*args)[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid task id provided")
	}

	var s task.Status

	if status == markInProgress {
		s = task.Inprogress
	}

	if status == markDone {
		s = task.Done
	}

	return t.UpdateStatus(id, s)
}

func listTasks(args *[]string, t *task.Task) error {
	if len(*args) == 1 {
		handlePrintTasks(nil, t)
		return nil
	}

	if len(*args) != 2 {
		return errors.New("list command requires a task status")
	}

	status := task.Status((*args)[1])

	if status == task.Done {
		return handlePrintTasks(&status, t)
	}

	if status == task.Inprogress {
		return handlePrintTasks(&status, t)
	}

	if status == task.Todo {
		return handlePrintTasks(&status, t)
	}

	return errors.New("invalid task status")
}

func handlePrintTasks(s *task.Status, t *task.Task) error {
	tasks, err := t.GetByStatus(s)
	if err != nil {
		return err
	}

	if len(*tasks) == 0 {
		return errors.New("no tasks registered yet")
	}

	for _, v := range *tasks {
		fmt.Printf("ID: %-12d DESCRIPTION: %-50s STATUS: %-12s CREATED AT: %-50v UPDATED AT: %-50v\n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
	}

	return nil
}

func printError(err error) {
	if err != nil {
		fmt.Printf("Issue found: %v\n", err)
	}
}
