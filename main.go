package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var todos []string
	fmt.Println("TODO App")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("What should we do next? (type help to see the options)")
		scanner.Scan()
		switch command := scanner.Text(); command {
			case HelpCommand:
				fmt.Println("Available Commands:")
				// todo: iterate with for loop
				fmt.Println("help - help menu")
				fmt.Println("create - create a todo")
			case CreateCommand:
				fmt.Println("Type what should be done:")
				scanner.Scan()
				todos = append(todos, scanner.Text())
				printHero("Your todo list:", todosToString(todos))
			default:
		}
	}
}