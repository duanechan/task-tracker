package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// State object of the Task Tracker CLI.
type state struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

var (
	cwd, _   = os.Getwd()
	filename = "/.tasktracker.json"
)

// Reads the saved state from JSON file.
func readState() (*state, error) {
	path := filepath.Join(cwd, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := saveState(&state{NextID: 0, Tasks: []Task{}}); err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 06444)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state state
	if err = json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}

	return &state, nil
}

// Writes the given application state to the JSON file.
func saveState(state *state) error {
	path := filepath.Join(cwd, filename)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}

func (s state) displayTasks(pred func(t Task) bool) {
	sb := strings.Builder{}
	
	for i, t := range s.Tasks {
		if pred(t) {
			sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, t))
		}
	}
	
	if sb.String() == "" {
		sb.WriteString("No tasks to display.")
	}

	fmt.Println(sb.String())
}
