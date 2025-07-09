# Task Tracker CLI

This project is a Task Tracker where you can create your tasks and keep track of their status (todo, in-progress and done).

The database is a JSON file that you can take it wherever you want, so your data won't be in anyone's database.

Project based on the following requirements: https://roadmap.sh/projects/task-tracker

## Running the unit tests

You can run the unit tests by simply running `go test -v ./*.go`. The `-v` option is for the verbose output.

## Building the project

To build the project, you will need to have golang installed. Then execute the command `go build -o tt` (tt being the name of the executable file and tt ).

## Features

Add a new task:

```sh
./tt add "download the Task Tracker CLI"
# output: Task added successfully (ID: 1)
```

List all tasks:

```sh
./tt list
# output: 1 (todo), download the Task Tracker CLI
```

Update a task:

```sh
./tt update 1
# output: Task updated successfully (ID: 1)
```

Delete a task:

```sh
./tt delete 1
# output: Task deleted successfully (ID: 1)
```
