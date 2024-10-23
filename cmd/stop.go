package cmd

import (
	"fmt"
	"strconv"

	"github.com/mosaleh52/aw-timer/internal/helpers"
	"github.com/mosaleh52/aw-timer/internal/timer"
	"github.com/spf13/cobra"
)

// TODO:handel when there is no todo
func stopTimerCmd() *cobra.Command {
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a task",
		Run: func(cmd *cobra.Command, args []string) {
			currentTodos := timer.GetCurrentTodos(apiUrl, bucketId, dateLayout)
			if len(currentTodos) == 1 {
				timer.StopTimer(apiUrl, bucketId, dateLayout, strconv.Itoa(currentTodos[0].ID))
			} else if len(currentTodos) == 0 {
				helpers.PretyPrint("no running todo", "red", "term")
			} else {
				fmt.Println("havenot implement stoping multible todos yet")
			}
		},
	}

	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999999-07:00", "specify the dateLayout used for formatting in aw server")
	return cmd
}
