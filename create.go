package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func renderCreateCommandBullet(i int) string {
	return strconv.Itoa(i + 1) + " "
}

func createCommand(todos []todo) []todo {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for {
		message := "Type what should be done:"
		if (i > 0) {
			message = "What else should be done? (type exit to get back to the main menu or remove to remove the last entry)"
		} 
		fmt.Println(message)
		scanner.Scan()
		if (scanner.Text() == "exit") {
			return todos
		}
		if (scanner.Text() == "remove") {
			if (len(todos) != 0) {
				todos = todos[:len(todos) - 1]	
			}
			printHero("Your todo list:", todosToString(todos, renderCreateCommandBullet))
			continue
		}
		todos = append(todos, todo{ title: scanner.Text(), completed: false})
		printHero("Your todo list:", todosToString(todos, renderCreateCommandBullet))
		i++
	}
}