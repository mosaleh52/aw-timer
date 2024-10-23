package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "todoTimer"}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(startTimerCmd(), stopTimerCmd(), getCurrentTodoCmd(), toggleCurrentTodo())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// isUpdateAvailable, _, latestVersion := CheckForUpdates()
}
