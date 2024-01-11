package cmd

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

func Run(todos tickit.TodoList, FILENAME string) {
	addTodo := flag.Bool("a", false, "add a new task\nexample: tickit -a new task")
	listTodo := flag.Bool("l", false, "list all tasks")
	tickTodo := flag.Int("t", 0, "tick task with given id as complete (id must not be 0)\nexample: tickit -t 1")
	delTodo := flag.Int("d", 0, "delete task with given id (id must not be 0)\nexample: tickit -d 1")

	flag.Parse()

	switch {
	case *addTodo:
		args := flag.Args()
		if len(args) > 0 {
			task := strings.Join(args, " ")
			newTask := tickit.NewTask((len(todos) + 1), task, time.Now(), false)
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

		for _, i := range todos {
			status := color.RedString("not completed")
			if i.IsTicked {
				status = color.GreenString("completed")
			}
			t.AppendRow(table.Row{i.ID, i.Task, i.Started.Format(time.RFC822), status})
		}

		t.Render()
	case *tickTodo != 0:
		t, i, err := tickit.GetByID(*tickTodo, todos)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		task := tickit.TickTask(t)
		todos = append(todos[:i], append([]tickit.Task{task}, todos[i+1:]...)...)
		f, _ := json.MarshalIndent(&todos, "", " ")
		_ = os.WriteFile(FILENAME, f, 0644)
		fmt.Println("Task marked as complete")
	case *delTodo != 0:
		_, i, err := tickit.GetByID(*delTodo, todos)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		todos = append(todos[:i], todos[i+1:]...)
		f, _ := json.MarshalIndent(&todos, "", " ")
		_ = os.WriteFile(FILENAME, f, 0644)
		fmt.Println("Task removed from your list")
	default:
		flag.CommandLine.Usage()
	}
}
