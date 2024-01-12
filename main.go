package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/VivekAlhat/tickit/cmd"
	"github.com/VivekAlhat/tickit/internal"
)

const FILENAME string = "tasks.json"

func main() {
	var appDir string
	switch runtime.GOOS {
	case "darwin", "linux":
		user, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		appDir = filepath.Join(user.HomeDir, "tickit")
		if err := os.MkdirAll(appDir, 0700); err != nil {
			fmt.Println("Error occured while creating application directory", err)
			os.Exit(1)
		}
	case "windows":
		user := os.Getenv("USERPROFILE")
		if user == "" {
			fmt.Println("Error getting user profile")
			os.Exit(1)
		}
		appDir = filepath.Join(user, "tickit")
		if err := os.MkdirAll(appDir, 0700); err != nil {
			fmt.Println("Error occured while creating application directory", err)
			os.Exit(1)
		}
	default:
		fmt.Println("This OS isn't supported")
		os.Exit(1)
	}

	filePath := filepath.Join(appDir, FILENAME)
	_, err := os.Stat(filePath)

	if err != nil {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Tasks file created at location %q\n", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var todos tickit.TodoList

	json.Unmarshal(data, &todos)

	cmd.Run(todos, filePath)
}
