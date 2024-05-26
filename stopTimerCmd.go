package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO:handel when there is no todo
func stopTimerCmd() *cobra.Command {
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a task",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				stopTimer(apiUrl, bucketId, dateLayout, getCurrentTodo(apiUrl, bucketId, dateLayout)[0].ID)
			} else {
				fmt.Println("Error: you can only stop timer for one todo")
			}
		},
	}

	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999999-07:00", "specify the dateLayout used for formatting in aw server")
	return cmd
}
