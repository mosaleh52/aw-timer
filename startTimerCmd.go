package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startTimerCmd() *cobra.Command {
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
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

	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999999-07:00", "specify the dateLayout used for formatting in aw server")
	return cmd
}
