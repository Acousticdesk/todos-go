package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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
				var sb strings.Builder
				for _, c := range commandList {
					sb.WriteString(c[0])
					sb.WriteString(" - ")
					sb.WriteString(c[1])
					sb.WriteString("\n")
				}
				fmt.Println(sb.String())
			case CreateCommand:
				fmt.Println("Type what should be done:")
				scanner.Scan()
				todos = append(todos, scanner.Text())
				printHero("Your todo list:", todosToString(todos))
			case DoneCommand:
				fmt.Println("Type what should be done:")
				scanner.Scan()
				todos = append(todos, scanner.Text())
				printHero("Your todo list:", todosToString(todos))
			default:
		}
	}
}