package tickit

import "time"

type Task struct {
	Task     string
	Started  time.Time
	IsTicked bool
}

func NewTask(task string, started time.Time, isTicked bool) Task {
	return Task{task, started, isTicked}
}
