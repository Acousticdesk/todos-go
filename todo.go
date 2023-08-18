package main

import "strings"

type todo struct {
	title string
	completed bool
}

func todosToString(todos []todo, renderBullet func(int) string) string {
	todosContent := ""
	for i := 0; i < len(todos); i++ {
		todosContent += renderBullet(i)
		if (todos[i].completed) {
			todosContent += "✅"
		} else {
			todosContent += "❌"
		}
		todosContent += " "
		todosContent += todos[i].title
		todosContent += "\n\r"
	}
	return strings.TrimSuffix(todosContent, "\n\r")
}

const HelpCommand = "help"
const CreateCommand = "create"
const DoneCommand = "done"
const EmailCommand = "email"
