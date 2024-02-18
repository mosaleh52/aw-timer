package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var apiUrl, bucketId, dateLayout string
	rootCmd := &cobra.Command{Use: "todoTimer"}
	rootCmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	rootCmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	rootCmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999Z", "specify the dateLayout used for formatting in aw server")

	cmdStart := &cobra.Command{
		Use:   "start",
		Short: "Start a task",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				startTimer(apiUrl, bucketId, dateLayout, args[0])
			} else {
				fmt.Println("Error: you can only start timer for one todo")
			}
		},
	}

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

	cmdCurrent := &cobra.Command{
		Use:   "current",
		Short: "get the current running todo",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// check for the current todo and return to the stdout would be usefull for bars and scripts
			} else {
				fmt.Println("Error: the current command does not accept any commands")
			}
		},
	}

	rootCmd.AddCommand(cmdStart, cmdStop, cmdCurrent)
	rootCmd.Execute()
}
