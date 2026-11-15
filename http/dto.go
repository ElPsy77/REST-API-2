package http

import (
	"encoding/json"
	"errors"
	"time"
)

type CompleteTaskDTO struct {
	Complete bool `json:"complete"`
}

// DTO (Data Transfer Object)
type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("Title is empty")
	}

	if t.Description == "" {
		return errors.New("Description is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
