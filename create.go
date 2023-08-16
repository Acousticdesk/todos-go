package main

import (
	"fmt"
	"bufio"
	"os"
)

func createCommand(todos []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type what should be done:")
	scanner.Scan()
	todos = append(todos, scanner.Text())
	printHero("Your todo list:", todosToString(todos))
	return todos
}