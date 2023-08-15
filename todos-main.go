package main

import "strings"

func todosToString(todos []string) string {
	todosContent := ""
	for i := 0; i < len(todos); i ++ {
		todosContent += todos[i]
		todosContent += "\n\r"
	}
	return strings.Replace(todosContent, "\n\r", "", -1)
}