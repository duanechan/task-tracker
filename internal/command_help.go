package task

import (
	"fmt"
	"slices"
	"strings"
)

func commandHelp(c *CLI, args []string) error {
	if len(args) > 1 {
		return ErrTooManyArgs
	}

	if len(args) == 0 {
		c.DisplayCommands()
		return nil
	}

	name := strings.TrimSpace(args[0])
	command, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", name)
	}

	fmt.Println(Bold+Blue+"Task Tracker"+Reset, Bold+c.version+Reset)
	fmt.Println()
	fmt.Printf("%s: %s\n\n", Bold+Blue+name+Reset, command.description)
	fmt.Printf("%sUsage%s: task-cli %s\n\n", Bold, Reset, command.usage)

	if len(command.params) == 0 {
		fmt.Println(Bold+"Params:"+Reset, "none")
		return nil
	}

	fmt.Println(Bold + "Params:" + Reset)

	keys := make([]string, 0, len(command.params))
	for k := range command.params {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	maxLen := 0
	for _, k := range keys {
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}

	for _, param := range keys {
		desc := command.params[param]
		coloredParam := Bold + Blue + param + Reset
		padding := maxLen - len(param)
		fmt.Printf("   %s%s - %s\n", coloredParam, strings.Repeat(" ", padding), desc)
	}

	return nil
}
