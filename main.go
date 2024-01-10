package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/VivekAlhat/tickit/internal"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

const FILENAME string = "tasks.json"

func handleError(err error) {
	fmt.Println(err)
}

func main() {
	_, err := os.Stat(FILENAME)

	if err != nil {
		_, err := os.Create(FILENAME)
		if err != nil {
			handleError(err)
			os.Exit(1)
		}
	}

	data, err := os.ReadFile(FILENAME)
	if err != nil {
		handleError(err)
	}

	var todos tickit.TodoList

	json.Unmarshal(data, &todos)

	addTodo := flag.Bool("a", false, "add a new task\nexample: tickit -a new task")
	listTodo := flag.Bool("l", false, "list all tasks")

	flag.Parse()

	switch {
	case *addTodo:
		args := flag.Args()
		if len(args) > 0 {
			task := strings.Join(args, " ")
			newTask := tickit.NewTask(task, time.Now(), false)
			todos = append(todos, newTask)
			f, _ := json.MarshalIndent(&todos, "", " ")
			_ = os.WriteFile(FILENAME, f, 0644)
			fmt.Println("Task added to your list")
		} else {
			fmt.Println("Add task name after the -a flag")
			os.Exit(1)
		}
	case *listTodo:
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true

		t.AppendHeader(table.Row{"#", "Task", "Started On", "Status"})

		for index, i := range todos {
			status := color.RedString("not completed")
			if i.IsTicked {
				status = color.GreenString("completed")
			}
			t.AppendRow(table.Row{(index + 1), i.Task, i.Started.Format(time.RFC822), status})
		}

		t.Render()
	default:
		flag.CommandLine.Usage()
	}
}
