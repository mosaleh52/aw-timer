package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "todoTimer"}

	rootCmd.AddCommand(startTimerCmd(), stopTimerCmd(), getCurrentTodoCmd())
	rootCmd.Execute()
}
