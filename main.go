package main

import (
	"encoding/json"
	"fmt"
	"github.com/VivekAlhat/tickit/cmd"
	"github.com/VivekAlhat/tickit/internal"
	"os"
)

const FILENAME string = "tasks.json"

func main() {
	_, err := os.Stat(FILENAME)

	if err != nil {
		_, err := os.Create(FILENAME)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	data, err := os.ReadFile(FILENAME)
	if err != nil {
		fmt.Println(err)
	}

	var todos tickit.TodoList

	json.Unmarshal(data, &todos)

	cmd.Run(todos, FILENAME)
}
