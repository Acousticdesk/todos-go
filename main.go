package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
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