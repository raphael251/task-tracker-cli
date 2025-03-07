package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/raphael251/tasktrackercli"
)

var _ = Describe("Tasktrackercli", func() {
	var args []string

	BeforeEach(func() {
		args = []string{}
	})

	Describe("Executing any command", func() {
		Context("with no args", func() {
			It("should return a text error", func() {
				args = []string{}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("You need to pass at least one command"))
			})
		})

		Context("that does not exists", func() {
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
	})

	Describe("Executing the list command", func() {
		Context("with no tasks to show", func() {
			It("should return a text explaining that there is no tasks to show", func() {
				args = []string{"list"}
				response := main.ProcessCommand(args)
				Expect(response).To(Equal("There is no tasks to show"))
			})
		})
		// Context("with one or more tasks to show", func() {
		// 	It("should return a text with the tasks, one per line, with the id and the title separated by a comma")
		// })
	})
})
