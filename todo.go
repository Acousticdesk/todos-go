package main

import "strings"

var todos = []string{}

func todosToString(todos []string, renderBullet func(int) string) string {
	todosContent := ""
	for i := 0; i < len(todos); i++ {
		todosContent += renderBullet(i)
		todosContent += todos[i]
		todosContent += "\n\r"
	}
	return strings.TrimSuffix(todosContent, "\n\r")
}

const HelpCommand = "help"
const CreateCommand = "create"
const DoneCommand = "done"
