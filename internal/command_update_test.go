package task

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCommandUpdate(t *testing.T) {
	t.Cleanup(deleteTempJSON)

	testState := state{Tasks: []Task{
		{ID: 1, Description: "Old Task", CreatedAt: time.Now()},
		{ID: 2, Description: "Another Task", CreatedAt: time.Now()},
	}}

	t.Run("successful update", func(t *testing.T) {
		args := []string{"1", "Updated Task"}
		if err := commandUpdate(&testState, args); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		task := testState.Tasks[0]

		if task.Description != args[1] {
			t.Errorf("expected description to be %q, got %q", args[1], task.Description)
		}

		if task.UpdatedAt.IsZero() {
			t.Error("expected UpdatedAt to be set")
		}
	})

	t.Run("error cases", func(t *testing.T) {
		tests := []struct {
			name string
			args []string
			want error
		}{
			{"missing args", []string{"1"}, ErrMissingArg},
			{"too many args", []string{"1", "desc", "extra"}, ErrTooManyArgs},
			{"empty ID", []string{"   ", "desc"}, ErrEmptyID},
			{"empty description", []string{"1", "   "}, ErrEmptyDescription},
			{"non-existent task", []string{"99", "desc"}, ErrTaskNotFound},
			{"invalid ID", []string{"abc", "desc"}, strconv.ErrSyntax}, // strconv.Atoi error
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				stateCopy := state{Tasks: append([]Task(nil), testState.Tasks...)}
				err := commandUpdate(&stateCopy, tc.args)
				if err == nil {
					t.Fatalf("expected error %v, got nil", tc.want)
				}
				if !errors.Is(err, tc.want) && err.Error() != tc.want.Error() {
					t.Fatalf("expected error %v, got %v", tc.want, err)
				}
			})
		}
	})
}

func TestCommandUpdateStdout(t *testing.T) {
	t.Cleanup(deleteTempJSON)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = old })

	testState := state{Tasks: []Task{
		{ID: 1, Description: "Old Task", CreatedAt: time.Now()},
	}}

	args := []string{"1", "Updated Task"}
	if err := commandUpdate(&testState, args); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	w.Close()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}

	output := buf.String()
	expected := "Updated Task (ID: 1) description to Updated Task\n"
	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Errorf("expected stdout %q, got %q", expected, output)
	}
}
