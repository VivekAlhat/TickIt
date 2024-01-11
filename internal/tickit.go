package tickit

import (
	"errors"
	"time"
)

var ErrTaskNotFound error = errors.New("Task with given id does not exist")

type Task struct {
	ID       int
	Task     string
	Started  time.Time
	IsTicked bool
}

type TodoList = []Task

func NewTask(id int, task string, started time.Time, isTicked bool) Task {
	return Task{id, task, started, isTicked}
}

func TickTask(task Task) Task {
	return Task{task.ID, task.Task, task.Started, true}
}

func GetByID(id int, todos TodoList) (Task, int, error) {
	for i, t := range todos {
		if t.ID == id {
			return t, i, nil
		}
	}
	return Task{}, 0, ErrTaskNotFound
}
