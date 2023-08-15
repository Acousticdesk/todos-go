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
	for true {
		fmt.Println("Add a new todo item:")
		scanner.Scan()
		todos = append(todos, scanner.Text())
		printHero("Your todo list:", todosToString(todos))
	}
}