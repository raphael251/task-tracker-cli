package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/raphael251/tasktrackercli/taskmanager"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to pass at least one command")
		return
	}

	result := ProcessCommand(os.Args[1:])
	fmt.Println(result)
}

func ProcessCommand(args []string) string {
	if len(args) == 0 {
		return "You need to pass at least one command"
	}

	command := args[0]
	commandArgs := args[1:]

	switch command {
	case "add":
		if len(commandArgs) != 1 {
			return "You must pass exactly one argument to the add command, which is the description of your task"
		}

		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("could not add task")
			return "Error adding new task"
		}

		taskAdded := tasksManager.AddTask(commandArgs[0])

		tasksManager.Save()

		return fmt.Sprintf("Task added successfully (ID: %v)", taskAdded.Id)
	case "update":
		if len(commandArgs) != 2 {
			return "You must pass exactly two params: one for the ID you want to update and one for the new description you want"
		}

		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("there is no tasks to update")
			return "Error updating a task"
		}

		taskUpdated := tasksManager.UpdateTask(commandArgs[0], commandArgs[1])

		if taskUpdated.Id == "" {
			return "the task with the specified ID was not found"
		}

		tasksManager.Save()

		return fmt.Sprintf("Task updated successfully (ID: %v)", taskUpdated.Id)
	case "delete":
		if len(commandArgs) != 1 {
			return "You must pass exactly one argument to the delete command, which is the id of the task to be deleted"
		}

		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("could not delete the task")
			return "Error deleting a task"
		}

		deletedTask := tasksManager.DeleteTask(commandArgs[0])

		if deletedTask.Id == "" {
			return "the task with the specified ID was not found"
		}

		tasksManager.Save()

		return fmt.Sprintf("Task deleted successfully (ID: %v)", commandArgs[0])
	case "list":
		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("could not list tasks")
			return "Error listing tasks"
		}

		tasks := tasksManager.ListTasks()

		if tasks == nil || len(tasks) == 0 {
			return "There is no tasks to show"
		}

		var sb strings.Builder

		for _, task := range tasks {
			sb.WriteString(fmt.Sprintf("%v (%v), %v\n", task.Id, task.Status, task.Description))
		}

		return sb.String()
	case "mark-in-progress":
		if len(commandArgs) != 1 {
			return "You must pass exactly one argument to the mark-in-progress command, which is the id of the task to be marked"
		}

		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("could not mark task as in-progress")
			return "Error marking task as in-progress"
		}

		task := tasksManager.MarkTaskInProgress(commandArgs[0])

		if task.Id == "" {
			fmt.Println("task not found")
			return "Error marking task as in-progress"
		}

		tasksManager.Save()

		return fmt.Sprintf("Task marked as in-progress successfully (ID: %v)", task.Id)
	case "mark-done":
		if len(commandArgs) != 1 {
			return "You must pass exactly one argument to the mark-done command, which is the id of the task to be marked"
		}

		tasksManager, err := taskmanager.GetTaskManager()

		if err != nil {
			fmt.Println("could not mark task as done")
			return "Error marking task as done"
		}

		task := tasksManager.MarkTaskDone(commandArgs[0])

		if task.Id == "" {
			fmt.Println("task not found")
			return "Error marking task as done"
		}

		tasksManager.Save()

		return fmt.Sprintf("Task marked as done successfully (ID: %v)", task.Id)
	default:
		return "Invalid command. Please run the help command"
	}
}
