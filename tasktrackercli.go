package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
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

		taskAdded, err := addTask(commandArgs[0])

		if err != nil {
			fmt.Println(err)
			return "Error adding new task"
		}

		return fmt.Sprintf("Task added successfully (ID: %v)", taskAdded.Id)
	case "update":
		if len(commandArgs) != 2 {
			return "You must pass exactly two params: one for the ID you want to update and one for the new description you want"
		}

		taskUpdated, err := updateTask(commandArgs[0], commandArgs[1])

		if err != nil {
			fmt.Println(err)
			return "Error updating a task"
		}

		return fmt.Sprintf("Task updated successfully (ID: %v)", taskUpdated.Id)
	case "delete":
		if len(commandArgs) != 1 {
			return "You must pass exactly one argument to the delete command, which is the id of the task to be deleted"
		}

		err := deleteTask(commandArgs[0])

		if err != nil {
			fmt.Println(err)
			return "Error deleting a task"
		}

		return fmt.Sprintf("Task deleted successfully (ID: %v)", commandArgs[0])
	default:
		return "Invalid command. Please run the help command"
	}
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type DB struct {
	LastUsedId int `json:"id"`
	Tasks      []Task
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "inprogress"
	StatusDone       = "done"
)

func getFilePath() string {
	return "tasks.json"
}

func addTask(title string) (*Task, error) {
	db, err := getDB()

	// If it's the first task added (the file tasks.json doesn't exists)
	if err != nil {
		firstId := 1

		task := Task{
			fmt.Sprint(firstId),
			title,
			StatusTodo,
		}

		db := &DB{
			LastUsedId: firstId,
			Tasks: []Task{
				task,
			},
		}

		saveDB(db)

		return &task, nil
	}

	incrementedId := db.LastUsedId + 1

	db.LastUsedId = incrementedId

	task := Task{
		fmt.Sprint(incrementedId),
		title,
		StatusTodo,
	}

	db.Tasks = append(db.Tasks, task)

	saveDB(db)

	return &task, nil
}

func updateTask(id string, newTitle string) (*Task, error) {
	db, err := getDB()

	if err != nil {
		return nil, fmt.Errorf("there is no tasks to update")
	}

	var updatedTask *Task

	for i := 0; i < len(db.Tasks); i += 1 {
		if db.Tasks[i].Id == id {
			db.Tasks[i].Title = newTitle
			updatedTask = &db.Tasks[i]

			saveDB(db)

			return updatedTask, nil
		}
	}

	return nil, fmt.Errorf("the task with the specified ID was not found")
}

func deleteTask(id string) error {
	db, err := getDB()

	if err != nil {
		return fmt.Errorf("could not delete the task")
	}

	for i := 0; i < len(db.Tasks); i += 1 {
		if db.Tasks[i].Id == id {
			db.Tasks = slices.Delete(db.Tasks, i, i+1)

			saveDB(db)

			return nil
		}
	}

	return fmt.Errorf("the task with the specified ID was not found")
}

func getDB() (*DB, error) {
	jsonFile, err := os.Open(getFilePath())

	if err != nil {
		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var db DB

	json.Unmarshal(byteValue, &db)

	return &db, nil
}

func saveDB(db *DB) error {
	marshalledDB, err := json.Marshal(db)

	if err != nil {
		return err
	}

	createdFile, err := os.Create(getFilePath())

	if err != nil {
		return err
	}

	createdFile.Write(marshalledDB)

	return nil
}
