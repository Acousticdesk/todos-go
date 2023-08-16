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

func createCommand(todos []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type what should be done:")
	scanner.Scan()
	todos = append(todos, scanner.Text())
	printHero("Your todo list:", todosToString(todos, renderCreateCommandBullet))
	return todos
}