package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func renderDoneCommandBullet(i int) string {
	return strconv.Itoa(i + 1) + " "
}

func doneCommand(todos []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which one should be marked as done?")

	scanner.Scan()
	
	selectedOrder, err := strconv.Atoi(scanner.Text())
	if (err != nil) {
		fmt.Println("Couldn't find the todo with this number")
		return todos
	}
	
	selectedIndex := selectedOrder - 1
	
	if (selectedIndex < 0 || selectedIndex >= len(todos)) {
		fmt.Println("Couldn't find the todo with this number")
		return todos
	}

	todos = append(todos[:selectedIndex], todos[selectedIndex + 1])

	printHero("Your new list:", todosToString(todos, renderDoneCommandBullet))

	return todos
}