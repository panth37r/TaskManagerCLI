package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCommand(s *bufio.Scanner) []string {
	s.Scan()
	lineFields := strings.Fields(s.Text())
	return lineFields
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	app := NewApp()
	fmt.Print("Please enter your name: ")

	s.Scan()
	err := app.InitApp(strings.TrimSpace(s.Text()))
	if err != nil {
		panic(err)
	}

	ClearScreen()
	fmt.Printf("Welcome %s \n", app.CurrentUser)

	for {
		fmt.Print("Enter command: ")
		args := GetCommand(s)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "exit":
			err = app.ExitApp()
			if err != nil {
				fmt.Print(err)
				fmt.Println("something went wrong. please fix it by yourself, im too lazy")

			} else {
				fmt.Println("bye")
			}
			return
		case "add":
			if len(args) < 2 {
				ClearScreen()
				fmt.Println("Enter description please: add \"description\"")
				continue
			}
			description := strings.Join(args[1:], " ")
			err = app.CreateTask(description)
			if err != nil {
				fmt.Println(err)
			}
		case "update":
			if len(args) < 3 {
				ClearScreen()
				fmt.Println("Enter description please: update id \"new_description\"")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			description := strings.Join(args[2:], " ")
			err = app.UpdateDescriptionTask(description, id)
			if err != nil {
				fmt.Println(err)
			}
		case "status":
			if len(args) < 3 {
				ClearScreen()
				fmt.Println("Enter id and status please: status id \"status\"[done,undone,in-progress]")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			status, err := app.StatusToCode(args[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = app.UpdateTaskStatus(id, status)
			if err != nil {
				fmt.Println(err)
			}
		case "all":
			ClearScreen()
			json, err := json.MarshalIndent(app.GetAllTasks(), "", "  ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(json))

		case "bystatus":
			ClearScreen()
			if len(args) < 2 {
				fmt.Println("enter a status code: bystatus STATUS [done, undone, in-progress]")
				continue
			}
			status, err := app.StatusToCode(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			json, err := json.MarshalIndent(app.GetTaskByStatus(status), "", "  ")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(json))

		case "delete":
			if len(args) < 2 {
				fmt.Println("enter id: delete id")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = app.DeleteTask(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "help":
			ClearScreen()
			fmt.Println("to delete: delete id")
			fmt.Println("to get all tasks: all")
			fmt.Println("to getbystatus: bystatus [done/undone/in-progress]")
			fmt.Println("to add: add description")
			fmt.Println("to updateDescription: update id newDescription")
			fmt.Println("to change status: status id [done/undone/in-progress]")
			fmt.Println("to leave: exit")
		default:
			ClearScreen()
			fmt.Println("Unknown command try enter: help")
		}

	}
}
