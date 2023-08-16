package main

import "strings"

// import "strings"

func todosToString(todos []string) string {
	todosContent := ""
	for i := 0; i < len(todos); i ++ {
		todosContent += todos[i]
		todosContent += "\n\r"
	}
	return strings.TrimSuffix(todosContent, "\n\r")
}

const HelpCommand = "help"
const CreateCommand = "create"
