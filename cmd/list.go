/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Displays a list of all tasks currently stored.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use the service to get all tasks
		tasks, err := taskService.GetAllTasks()
		if err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Initialize tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0) // minwidth, tabwidth, padding, padchar, flags
		fmt.Fprintln(w, "ID\tSTATUS\tPRIORITY\tCATEGORY\tTITLE\tDUE DATE")
		fmt.Fprintln(w, "--\t------\t--------\t--------\t-----\t--------")

		for _, task := range tasks {
			// Format fields for display
			idShort := task.ID
			if len(idShort) > 8 {
				idShort = idShort[:8] // Show only first 8 chars of UUID
			}
			status := string(task.Status)
			priority := string(task.Priority)
			if priority == "" {
				priority = "-" // Show dash if not set
			}
			category := string(task.Category)
			if category == "" {
				category = "-" // Show dash if not set
			}
			dueDate := "-"
			if !task.DueDate.IsZero() {
				dueDate = task.DueDate.Format(time.DateOnly) // Format as YYYY-MM-DD
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
				idShort,
				status,
				priority,
				category,
				task.Title,
				dueDate,
			)
		}

		w.Flush() // Flush buffered output
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// TODO: Add flags for filtering (category, priority, due date)
	// e.g., listCmd.Flags().StringP("category", "c", "", "Filter by category")
	// e.g., listCmd.Flags().StringP("priority", "p", "", "Filter by priority")
	// e.g., listCmd.Flags().String("due", "", "Filter by due date (e.g., 'today', 'week')")
}
