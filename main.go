package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/VivekAlhat/tickit/internal"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

type TodoList = []tickit.Task

func main() {
	var task1 tickit.Task = tickit.NewTask("Go for a walk", time.Now().AddDate(0, -1, 0), true)
	var task2 tickit.Task = tickit.NewTask("Play Far Cry 4", time.Now().AddDate(0, 0, -10), false)
	var task3 tickit.Task = tickit.NewTask("Start a new CLI project in Go", time.Now().AddDate(0, 0, -5), true)
	var task4 tickit.Task = tickit.NewTask("Write a new article", time.Now().AddDate(0, 0, -2), false)
	var task5 tickit.Task = tickit.NewTask("Read 2 pages of a book", time.Now(), false)
	var todos TodoList = TodoList{task1, task2, task3, task4, task5}

	addTodo := flag.Bool("a", false, "add a new task")
	listTodo := flag.Bool("l", false, "list all tasks")
	hideCompleted := flag.Bool("x", false, "hide all completed tasks")

	flag.Parse()

	switch {
	case *addTodo:
		_ = append(todos, todos...)
		fmt.Println("Task added to your list")
	case *listTodo:
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true

		t.AppendHeader(table.Row{"#", "Task", "Started On", "Is Ticked"})

		for index, i := range todos {
			status := color.RedString("false")
			if i.IsTicked {
				status = color.GreenString("true")
			}
			t.AppendRow(table.Row{(index + 1), i.Task, i.Started.Format(time.RFC822), status})
		}

		t.Render()
	case *hideCompleted:
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleRounded)
		t.Style().Options.SeparateRows = true

		t.AppendHeader(table.Row{"#", "Task", "Started On", "Is Ticked"})

		for index, i := range todos {
			status := color.RedString("false")
			if !i.IsTicked {
				t.AppendRow(table.Row{(index + 1), i.Task, i.Started.Format(time.RFC822), status})
			}
		}

		t.Render()
	default:
		flag.CommandLine.Usage()
	}
}
