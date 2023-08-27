package main

import (
	"fmt"
	"strings"
)

var commandList = [][]string{
	{HelpCommand, "help menu"},
	{CreateCommand, "create a todo"},
	{DoneCommand, "mark a todo as done"},
	{EmailCommand, "send your todo list to your email"},
	{EmailCommand, "send your todo list to your email"},
	{ExitCommand, "quit the application"},
}

func helpCommand() {
	fmt.Println("Available Commands:")
	var sb strings.Builder
	for _, c := range commandList {
		sb.WriteString(c[0])
		sb.WriteString(" - ")
		sb.WriteString(c[1])
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}