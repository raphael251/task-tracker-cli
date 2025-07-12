package main_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/raphael251/tasktrackercli"
)

var _ = Describe("Tasktrackercli", func() {
	var args []string

	BeforeEach(func() {
		args = []string{}
		os.Remove("tasks.json")
	})

	Describe("Executing the app", func() {
		Context("with no command", func() {
			It("should return a text error", func() {
				args = []string{}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You need to pass at least one command"))
			})
		})

		Context("with a command that does not exists", func() {
			It("should return a text error", func() {
				args = []string{"anycommand"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("Invalid command. Please run the help command"))
			})
		})
	})

	Describe("Executing the add command", func() {
		Context("with less than one arg (title)", func() {
			It("should return a text error", func() {
				args = []string{"add"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly one argument to the add command, which is the description of your task"))
			})
		})

		Context("with more than one arg (title)", func() {
			It("should return a text error", func() {
				args = []string{"add", "one", "test"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly one argument to the add command, which is the description of your task"))
			})
		})
	})

	Describe("Executing the update command", func() {
		Context("with less than two args", func() {
			It("should return a text error", func() {
				args = []string{"update", "1"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly two params: one for the ID you want to update and one for the new description you want"))
			})
		})

		Context("with more than two args", func() {
			It("should return a text error", func() {
				args = []string{"update", "1", "new todo", "another todo"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly two params: one for the ID you want to update and one for the new description you want"))
			})
		})
	})

	Describe("Executing the delete command", func() {
		Context("with less than one arg (id)", func() {
			It("should return a text error", func() {
				args = []string{"delete"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly one argument to the delete command, which is the id of the task to be deleted"))
			})
		})

		Context("with more than one args (id)", func() {
			It("should return a text error", func() {
				args = []string{"delete", "1", "deleting"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You must pass exactly one argument to the delete command, which is the id of the task to be deleted"))
			})
		})

		Context("with the right quantity of args (one)", func() {
			BeforeEach(func() {
				main.ProcessCommand([]string{"add", "unit test all the code"})
			})

			It("should return the text informing that the task was deleted successfully", func() {
				args = []string{"delete", "1"}
				response := main.ProcessCommand((args))
				Expect(response).To(Equal("Task deleted successfully (ID: 1)"))
			})
		})
	})

	Describe("Executing the list command", func() {
		Context("with no tasks to show", func() {
			It("should return a text explaining that there is no tasks to show", func() {
				args = []string{"list"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("There is no tasks to show"))
			})
		})

		Context("with one or more tasks to show", func() {
			It("should return a text with the tasks, one per line, with the id and the title separated by a comma", func() {
				main.ProcessCommand([]string{"add", "learn go"})

				args = []string{"list"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("1 (todo), learn go\n"))
			})
		})

		Context("with the additional arg for filter", func() {
			Context("with more than one argument for the filter", func() {
				It("should use just the first argument and ignore the extra ones", func() {
					main.ProcessCommand([]string{"add", "learn go"})
					main.ProcessCommand([]string{"add", "learn java"})
					main.ProcessCommand([]string{"add", "learn rust"})
					main.ProcessCommand([]string{"mark-in-progress", "1"})

					args = []string{"list", "in-progress", "extra"}
					response := main.ProcessCommand(args)
					Expect(response).To(Equal("1 (in-progress), learn go\n"))
				})
			})

			Context("with only one arg for filter (the correct way)", func() {
				It("should use just the first argument and ignore the extra ones", func() {
					main.ProcessCommand([]string{"add", "learn go"})
					main.ProcessCommand([]string{"add", "learn java"})
					main.ProcessCommand([]string{"add", "learn rust"})
					main.ProcessCommand([]string{"mark-in-progress", "1"})

					args = []string{"list", "in-progress"}
					response := main.ProcessCommand(args)
					Expect(response).To(Equal("1 (in-progress), learn go\n"))
				})
			})
		})
	})

	Describe("Executing the mark-in-progress command", func() {
		Context("with no args being passed", func() {
			It("should return a text error", func() {
				response := main.ProcessCommand([]string{"mark-in-progress"})

				Expect(response).To(Equal("You must pass exactly one argument to the mark-in-progress command, which is the id of the task to be marked"))
			})
		})

		Context("with a task with a todo status", func() {
			It("should return a successful text message, containing the id of the marked task", func() {
				main.ProcessCommand([]string{"add", "learn golang"})

				response := main.ProcessCommand([]string{"mark-in-progress", "1"})

				Expect(response).To(Equal("Task marked as in-progress successfully (ID: 1)"))
			})
		})

		Context("with a task already with a in-progress status", func() {
			It("should return a successful text message, containing the id of the task", func() {
				main.ProcessCommand([]string{"add", "learn golang"})
				main.ProcessCommand([]string{"mark-in-progress", "1"})

				response := main.ProcessCommand([]string{"mark-in-progress", "1"})

				Expect(response).To(Equal("Task marked as in-progress successfully (ID: 1)"))
			})
		})

		Context("with a non-existing task id", func() {
			It("should return a successful text message, containing the id of the task", func() {
				response := main.ProcessCommand([]string{"mark-in-progress", "10"})

				Expect(response).To(Equal("Error marking task as in-progress"))
			})
		})
	})

	Describe("Executing the mark-done command", func() {
		Context("with no args being passed", func() {
			It("should return a text error", func() {
				response := main.ProcessCommand([]string{"mark-done"})

				Expect(response).To(Equal("You must pass exactly one argument to the mark-done command, which is the id of the task to be marked"))
			})
		})

		Context("with a task with a todo status", func() {
			It("should return a successful text message, containing the id of the marked task", func() {
				main.ProcessCommand([]string{"add", "learn golang"})

				response := main.ProcessCommand([]string{"mark-done", "1"})

				Expect(response).To(Equal("Task marked as done successfully (ID: 1)"))
			})
		})

		Context("with a task with a in-progress status", func() {
			It("should return a successful text message, containing the id of the marked task", func() {
				main.ProcessCommand([]string{"add", "learn golang"})
				main.ProcessCommand([]string{"mark-in-progress", "1"})

				response := main.ProcessCommand([]string{"mark-done", "1"})

				Expect(response).To(Equal("Task marked as done successfully (ID: 1)"))
			})
		})

		Context("with a task already with a done status", func() {
			It("should return a successful text message, containing the id of the task", func() {
				main.ProcessCommand([]string{"add", "learn golang"})
				main.ProcessCommand([]string{"mark-done", "1"})

				response := main.ProcessCommand([]string{"mark-done", "1"})

				Expect(response).To(Equal("Task marked as done successfully (ID: 1)"))
			})
		})

		Context("with a non-existing task id", func() {
			It("should return a successful text message, containing the id of the task", func() {
				response := main.ProcessCommand([]string{"mark-done", "10"})

				Expect(response).To(Equal("Error marking task as done"))
			})
		})
	})
})
