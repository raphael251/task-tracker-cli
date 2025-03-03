package main

import (
	"os"
	"testing"
)

func TestInvalidCommand(t *testing.T) {
	args := []string{"any-command", "any-arg"}
	expectedOutput := "Invalid command. Please run the help command"
	msg := processCommand(args)
	if msg != expectedOutput {
		t.Fatalf(`processCommand("%v") = %q, want "%v", error`, args[0], msg, expectedOutput)
	}
}

func TestAddCommand(t *testing.T) {
	args := []string{"add", "code"}
	expectedOutput := "Task added successfully (ID: 1)"

	defer os.Remove("tasks.json")
	msg := processCommand(args)
	if msg != expectedOutput {
		t.Fatalf(`processCommand("%v") = %q, want "%v", error`, args[0], msg, expectedOutput)
	}
}

func TestUpdateTaskCommand(t *testing.T) {
	idToUpdate := "1"
	newTaskDescription := "learn Cloud"

	addTask("any-task-title")

	expectedOutput := "Task updated successfully (ID: 1, new Description: \"learn Cloud\")"

	_, err := updateTask(idToUpdate, newTaskDescription)
	if err != nil {
		t.Fatalf(`update("%v", "%v") = %q, want "%v", error`, idToUpdate, newTaskDescription, err, expectedOutput)
	}
}
