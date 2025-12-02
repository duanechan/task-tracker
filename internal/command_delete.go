package task

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func commandDelete(c *CLI, args []string) error {
	if len(args) < 1 {
		return ErrMissingArg
	}

	if len(args) > 1 {
		return ErrTooManyArgs
	}

	idString := strings.TrimSpace(args[0])
	if idString == "" {
		return ErrEmptyArgs
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted := false
	c.state.Tasks = slices.DeleteFunc(c.state.Tasks, func(t Task) bool {
		if t.ID == id {
			fmt.Printf("Deleted Task: %s\n", t)
			deleted = true
			return true
		}
		return false
	})

	if !deleted {
		return ErrTaskNotFound
	}

	return saveState(c.state)
}
