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
		fmt.Println("-----------------")
		fmt.Println("Your todo list:")
		for i := 0; i < len(todos); i ++ {
			fmt.Println(todos[i])
		}
		fmt.Println("-----------------")
	}
}