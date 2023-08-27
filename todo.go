package main

import (
	"strings"
	"time"
)

type todo struct {
	title string
	completed bool
	dueDate time.Time
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
		todosContent += "⏰ "
		todosContent += todos[i].dueDate.Format("2006-01-02 15:04:05")
		todosContent += "\n\r"
	}
	return strings.TrimSuffix(todosContent, "\n\r")
}

const HelpCommand = "help"
const CreateCommand = "create"
const DoneCommand = "done"
const EmailCommand = "email"
const ExitCommand = "exit"
