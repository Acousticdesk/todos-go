package main

import (
	"fmt"
	"bufio"
	"os"
)

// todo akicha: create a reminder command
// todo akicha: add possibility to edit a todo
// todo akicha: show the current status of the todos using emojis
func main() {
	// todo akicha: the todos should not be here
	// todo akicha: add status done
	var todos []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("TODO App")
	for {
		fmt.Println("What should we do next? (type help to see the options)")
		scanner.Scan()
		switch command := scanner.Text(); command {
			case HelpCommand:
				helpCommand()
			case CreateCommand:
				todos = createCommand(todos)
			case DoneCommand:
				todos = doneCommand(todos)
			default:
		}
	}
}