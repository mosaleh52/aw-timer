package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "todoTimer"}

	cmdStop := &cobra.Command{
		Use:   "stop",
		Short: "stop current running todo",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// get the current todo and stop it by adding duriont to aw-entry  returning the stoped todo
			} else {
				fmt.Println("Error: the stop command does not accept any commands")
			}
		},
	}

	rootCmd.AddCommand(startTimerCmd(), cmdStop, getCurrentTodoCmd())
	rootCmd.Execute()
}
