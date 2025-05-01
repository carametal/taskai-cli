/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"taskai-cli/internal/config"
	"taskai-cli/internal/repository"
	"taskai-cli/internal/service"

	"github.com/spf13/cobra"
)

var cfg *config.Config
var taskRepo repository.TaskRepository
var taskService service.TaskService

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskai-cli",
	Short: "An AI-powered task management CLI",
	Long: `TaskAI CLI is a powerful task management tool that uses
AI to help you manage your tasks effectively.

It allows you to create, manage, and organize tasks with natural language
processing capabilities to make task management more intuitive.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfigAndRepo)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskai-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle") // Example flag removed for now
}

// initConfigAndRepo loads configuration and initializes the repository.
func initConfigAndRepo() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	taskRepo, err = repository.NewJSONTaskRepository(cfg.DataFilePath)
	if err != nil {
		log.Fatalf("Error initializing repository: %v", err)
	}
	
	// Initialize task service
	taskService = service.NewTaskService(taskRepo)
	
	log.Printf("Using data file: %s", cfg.DataFilePath) // Log data file path for debugging
}
