package tickit

import (
	"testing"
	"time"
)

func TestTickit(t *testing.T) {
	t.Run("should have 5 tasks in list", func(t *testing.T) {
		var task1 Task = NewTask(1, "Go for a walk", time.Now().AddDate(0, -1, 0), true)
		var task2 Task = NewTask(2, "Play Far Cry 4", time.Now().AddDate(0, 0, -10), false)
		var task3 Task = NewTask(3, "Start a new CLI project in Go", time.Now().AddDate(0, 0, -5), true)
		var task4 Task = NewTask(4, "Write a new article", time.Now().AddDate(0, 0, -2), false)
		var task5 Task = NewTask(5, "Read 2 pages of a book", time.Now(), false)

		var todos TodoList = TodoList{task1, task2, task3, task4, task5}
		got := len(todos)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
