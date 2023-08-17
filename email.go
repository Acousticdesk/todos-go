package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func sendMail(message string, userEmail string) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_FROM"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	to := []string{userEmail}
	currentTime := time.Now()
	date := strconv.Itoa(currentTime.Year()) + "-" + strconv.Itoa(int(currentTime.Month())) + "-" + strconv.Itoa(currentTime.Day())
	body := []byte("Subject: Your Todo List for " + date + "\n\r" + message)
  
	err := smtp.SendMail(os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_FROM"), to, body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func emailCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	if userEmail == "" {
		fmt.Println("Looks like you haven't shared your email yet. Please type it in below:")
		scanner.Scan()
		userEmail = scanner.Text()
	}
	sendMail(todosToString(todos, (func(i int) string {
		return strconv.Itoa(i + 1) + ") "
	})), userEmail)
}


