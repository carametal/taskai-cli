/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"taskai-cli/internal/models"
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
taskai-cli add "Finish the report"`,
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
		
		// Get flag values
		description, _ := cmd.Flags().GetString("description")
		categoryStr, _ := cmd.Flags().GetString("category")
		priorityStr, _ := cmd.Flags().GetString("priority")
		
		// Convert to proper types
		category := models.Category(categoryStr)
		priority := models.Priority(priorityStr)
		
		// Create a task object with the necessary fields
		task := models.Task{
			Title:       title,
			Description: description,
			Category:    category,
			Priority:    priority,
		}
		
		// Use the taskRepo directly to avoid the potential issue with the service
		err := taskRepo.AddTask(task)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}

		fmt.Printf("✓ Task added: \"%s\"\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Add flags for task details
	addCmd.Flags().StringP("description", "d", "", "Task description")
	addCmd.Flags().StringP("category", "c", string(models.CategoryOther), 
		fmt.Sprintf("Task category (%s, %s, %s, %s, %s)", 
			models.CategoryWork, models.CategoryPersonal, models.CategoryLearning, 
			models.CategoryHealth, models.CategoryOther))
	addCmd.Flags().StringP("priority", "p", string(models.PriorityMedium), 
		fmt.Sprintf("Task priority (%s, %s, %s)", 
			models.PriorityHigh, models.PriorityMedium, models.PriorityLow))
}
