package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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
		case EmailCommand:
			if (len(todos) == 0) {
				fmt.Println("Looks like your todo list is empty. Use the create command to create a new item.")
				continue
			}
			emailCommand()
		default:
		}
	}
}
