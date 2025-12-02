package task

import (
	"fmt"
	"time"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func (t Task) String() string {
	color := Reset
	status := ""
	switch t.Status {
	case Todo:
		color = Yellow
		status = "[TODO]"
	case InProgress:
		color = Blue
		status = "[IN-PROGRESS]"
	case Done:
		color = Green
		status = "[DONE]"
	default:
		color = Red
		status = "[UNKNOWN]"
	}

	return fmt.Sprintf("%s%s%s%s ID:%d - %s", Bold, color, status, Reset, t.ID, t.Description)
}
