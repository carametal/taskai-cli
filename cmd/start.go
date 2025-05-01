/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"taskai-cli/internal/models"
	"taskai-cli/internal/repository" // Import repository for error checking

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [task ID]",
	Short: "Mark a task as in-progress",
	Long: `Marks the task with the specified ID as in-progress.
You can provide the full ID or the first 8 characters.
For example:
taskai-cli start 12345678`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one argument: the task ID")
		}
		// Basic ID format check (can be improved)
		if len(args[0]) < 8 {
			return errors.New("task ID must be at least 8 characters long")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]

		// Attempt to get the task first to ensure it exists
		task, err := taskService.GetTaskByID(taskID)
		if err != nil {
			if errors.Is(err, repository.ErrTaskNotFound) {
				fmt.Printf("Error: Task with ID '%s' not found.\n", taskID)
			} else {
				fmt.Printf("Error retrieving task: %v\n", err)
			}
			return
		}

		// Check if already in progress
		if task.Status == models.StatusInProgress {
			fmt.Printf("Task '%s' (%s) is already marked as in-progress.\n", task.Title, task.ID[:8])
			return
		}

		// Check if it's already completed
		if task.Status == models.StatusDone {
			fmt.Printf("Task '%s' (%s) is already marked as done. Cannot set to in-progress.\n", task.Title, task.ID[:8])
			return
		}

		// Use the service to start the task
		err = taskService.StartTask(taskID)
		if err != nil {
			fmt.Printf("Error starting task: %v\n", err)
			return
		}

		fmt.Printf("▶ Task marked as in-progress: \"%s\" (ID: %s)\n", task.Title, task.ID[:8])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// No flags needed for basic start for now
}