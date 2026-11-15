package todo

import (
	"time"
)

type Task struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`

	CreatedAt  time.Time  `json:"created_at"`
	CompleteAt *time.Time `json:"complete_at,omitempty"`
}

func NewTask(title string, text string) Task {
	return Task{
		Title:     title,
		Text:      text,
		Completed: false,

		CreatedAt:  time.Now(),
		CompleteAt: nil,
	}
}

func (t *Task) Complete() {
	t.Completed = true

	timeNow := time.Now()
	t.CompleteAt = &timeNow
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompleteAt = nil
}
