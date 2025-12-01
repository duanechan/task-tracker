package task

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func commandMarkDone(state *state, args []string) error {
	if len(args) < 1 {
		return ErrMissingArg
	}

	if len(args) > 1 {
		return ErrTooManyArgs
	}

	idString := strings.TrimSpace(args[0])
	if idString == "" {
		return ErrEmptyID
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	marked := false
	for i, t := range state.Tasks {
		if t.ID == id {
			fmt.Printf("Task %s status updated to: %s\n", t, "Done")
			state.Tasks[i].Status = Done
			state.Tasks[i].UpdatedAt = time.Now()
			marked = true
			break
		} 
	}

	if !marked {
		return ErrTaskNotFound
	}
	
	return saveState(state)
}