package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// TODO:handel when there is no todo
func toggleCurrentTodo() *cobra.Command {
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
		Use:   "toggle Current Todo",
		Short: "toggle current  task for rest",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO:refactor to a diffrent file
			// handle paused first
			fileInfo, err := os.Stat("./db")
			if err == nil && fileInfo.Size() > 0 {
				fileData, err := os.ReadFile("./db")
				if err != nil {
					fmt.Println("Error reading from file:", err)
					return
				}
				retrievedNumber, err := strconv.Atoi(string(fileData))
				if err != nil {
					fmt.Println("Error converting string to number:", err)
					return
				}
				stopTimer(apiUrl, bucketId, dateLayout, retrievedNumber, true)
				err = os.Remove("./db")
				if err != nil {
					fmt.Println("Error removing file:", err)
					return
				}
				return
			}
			// handle stop and record
			res := getCurrentTodo(apiUrl, bucketId, dateLayout)
			if len(res) == 1 {
				err := os.WriteFile("./db", []byte(strconv.Itoa(res[0].ID)), 0644)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
				stopTimer(apiUrl, bucketId, dateLayout, res[0].ID, false)
				fmt.Println(res[0])
			}
		},
	}

	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999999-07:00", "specify the dateLayout used for formatting in aw server")
	return cmd
}
