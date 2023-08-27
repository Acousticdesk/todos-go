package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func renderCreateCommandBullet(i int) string {
	return strconv.Itoa(i + 1) + " "
}

func internalRemoveCommand(todos []todo) {
	if (len(todos) != 0) {
		todos = todos[:len(todos) - 1]	
	}
	printHero("Your todo list:", todosToString(todos, renderCreateCommandBullet))
}

func internalCreateCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	todoTitle := scanner.Text()
	message := "Is there a due date? (YYYY-MM-DD HH:MM:SS). Type skip to continue."
	fmt.Println(message)
	var todoDueDate time.Time
	for (scanner.Text() != "skip") {
		scanner.Scan()
		scannedDueDate, err := time.Parse("2006-01-02 15:04:05", scanner.Text())
		if (err != nil) {
			fmt.Println("Oops, we did not recognize the date. Try again:")
			continue
		}
		todoDueDate = scannedDueDate
		break
	}
	todos = append(todos, todo{ title: todoTitle, completed: false, dueDate: todoDueDate })
	printHero("Your todo list:", todosToString(todos, renderCreateCommandBullet))
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
			internalRemoveCommand(todos)
			continue
		}
		internalCreateCommand()
		i++
	}
}