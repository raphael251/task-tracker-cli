package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/raphael251/tasktrackercli"
)

var _ = Describe("Tasktrackercli", func() {
	Describe("Executing the add command", func() {
		Context("with less args than expected", func() {
			It("should return an error text", func() {

				Expect(main.ProcessCommand([]string{"add"})).To(Equal("You must pass exactly one argument to the add command, which is the description of your task"))
			})
		})
	})
})
