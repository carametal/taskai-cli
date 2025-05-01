/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"taskai-cli/internal/repository"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long: `Deletes the task with the specified ID.
You can provide the full ID or the first 8 characters.
For example:
taskai-cli delete 12345678`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one argument: the task ID")
		}
		// Basic ID format check
		if len(args[0]) < 8 {
			return errors.New("task ID must be at least 8 characters long")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]

		// First get the task to show details when deleting
		task, err := taskRepo.GetTaskByID(taskID)
		if err != nil {
			if errors.Is(err, repository.ErrTaskNotFound) {
				fmt.Printf("Error: Task with ID '%s' not found.\n", taskID)
			} else {
				fmt.Printf("Error retrieving task: %v\n", err)
			}
			return
		}

		// Get confirmation if force flag is not set
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			fmt.Printf("Are you sure you want to delete task \"%s\" (ID: %s)? [y/N]: ", 
				task.Title, task.ID[:8])
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Deletion canceled.")
				return
			}
		}

		// Delete the task
		err = taskRepo.DeleteTask(task.ID)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}

		fmt.Printf("✓ Task deleted: \"%s\" (ID: %s)\n", task.Title, task.ID[:8])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Add force flag to skip confirmation
	deleteCmd.Flags().BoolP("force", "f", false, "Delete without confirmation")
}
