package task

import (
	"fmt"
	"slices"
)

// Command definition of the Task Tracker CLI.
type Command struct {
	name        string
	description string
	usage       string
	callback    func(*State, []string) error
}

// State object of the Task Tracker CLI.
type State struct {
	currentID int
	commands  map[string]Command
	Tasks     []Task
}

func LoadState() (*State, error) {
	return &State{
		commands: map[string]Command{
			"add": {
				name:        "add",
				description: "Adds a task to the list.",
				usage:       "add <description>",
				callback:    commandAdd,
			},
			"update": {
				name:        "update",
				description: "Updates the task of a given ID with an updated description.",
				usage:       "update <id> <updated_description>",
				callback:    commandUpdate,
			},
			"delete": {
				name:        "delete",
				description: "Deletes a task of a given ID.",
				usage:       "delete <id>",
				callback:    commandDelete,
			},
			"list": {
				name:        "list",
				description: "Lists tasks by status or all.",
				usage:       "list [done|todo|in-progress]",
				callback:    commandList,
			},
		},
	}, nil
}

// Run the CLI state with the given arguments.
func (s *State) Run(args []string) error {
	name := args[0]
	commandArgs := args[1:]

	cmd, exists := s.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	return cmd.callback(s, commandArgs)
}

// Display Task Tracker CLI commands.
func (s *State) DisplayCommands() {
	keys := []string{}
	for k := range s.commands {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println("Task Tracker CLI v1.0")
	fmt.Println("Commands:")
	fmt.Println()

	for _, k := range keys {
		cmd := s.commands[k]
		fmt.Printf("* %s - %s\n  Usage: %s\n\n", cmd.name, cmd.description, cmd.usage)
	}
}
