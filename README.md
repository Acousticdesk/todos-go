# TODO app CLI using Go

## Pre-requisites
- Install Go
- Read the [Set up SMTP](#set-up-smtp) section

## Run the App
`go run *.go`

## About the App
The application allows you to create your todo list using CLI

### Available commands
| Command    | Description                                         |
| ---------- | --------------------------------------------------- |
| help       | output all the available commands                   |
| create     | create a new todo                                   |
| done       | mark a todo as completed                            |
| email      | send email with your todo list to a specified email |

## Set up SMPT
To send emails, you can use any SMTP server you like.
I recommend using the Gmail SMTP, but make sure to read the [Guide from Google](https://myaccount.google.com/lesssecureapps) first.

All the SMTP configuration is done via the .env file

- Rename the .env-example to .env
- Specify the credentials to your SMTP server

That's it, no more configuration is needed to send the emails

## Roadmap

- Add email notifications about the items in the todo list
- Show status of the list items using emojis
