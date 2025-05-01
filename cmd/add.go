/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"smarttask-cli/internal/models"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Long: `Adds a new task to your task list.
Provide the task title as arguments.
For example:
smarttask-cli add "Finish the report"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a task title argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Combine args into a single title (basic implementation)
		// TODO: Implement NLP parsing here later
		title := strings.Join(args, " ")

		task := models.Task{
			Title: title,
			// Other fields will get defaults (ID, Timestamps, Status) in AddTask
		}

		err := taskRepo.AddTask(task)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}

		fmt.Printf("✓ Task added: \"%s\" (ID: %s)\n", task.Title, task.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// TODO: Add flags for specifying details like priority, due date, etc.
	// e.g., addCmd.Flags().StringP("priority", "p", string(models.PriorityMedium), "Set task priority (high, medium, low)")
	// e.g., addCmd.Flags().StringP("due", "d", "", "Set due date (e.g., 'tomorrow', 'next friday', '2025-12-31')")
}
