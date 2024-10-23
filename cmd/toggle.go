package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mosaleh52/aw-timer/internal/api"
	"github.com/mosaleh52/aw-timer/internal/helpers"
	"github.com/mosaleh52/aw-timer/internal/timer"
	"github.com/mosaleh52/aw-timer/internal/utils"
	"github.com/spf13/cobra"
)

// TODO:handel when there is no todo
func toggleCurrentTodo() *cobra.Command {
	var apiUrl, bucketId, dateLayout, dbFilePath, coloringMethod string
	cmd := &cobra.Command{
		Use:   "toggle Current Todo",
		Short: "toggle current  task for rest",
		Run: func(cmd *cobra.Command, args []string) {
			currentTodos := timer.GetCurrentTodos(apiUrl, bucketId, dateLayout)
			if len(currentTodos) == 1 {
				err := os.WriteFile(dbFilePath, []byte(strconv.Itoa(currentTodos[0].ID)), 0644)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
				timer.StopTimer(apiUrl, bucketId, dateLayout, strconv.Itoa(currentTodos[0].ID))
				os.Exit(0)
			} else if utils.CheckIfPaused(dbFilePath) {
				pausedTodo, err := utils.GetPausedTodo(dbFilePath, apiUrl, bucketId)
				if err != nil {
					fmt.Println("error getting paused todo ", err)
					os.Exit(1)
				}
				modifiedTodo, err := utils.CloneTodoDataAsByte(pausedTodo, dateLayout)
				api.CreateAwEvent(apiUrl, bucketId, modifiedTodo)
				err = os.Remove(dbFilePath)
				if err != nil {
					fmt.Println("Error removing file:", err)
					os.Exit(1)
				}
				os.Exit(0)

			}
			helpers.PretyPrint("nothing to toggle ", "red", coloringMethod)
			os.Exit(0)
		},
	}

	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999999-07:00", "specify the dateLayout used for formatting in aw server")
	cmd.Flags().StringVarP(&dbFilePath, "file-path", "p", "./timer.txt", "specify path for file storing paused todo")
	cmd.Flags().StringVarP(&coloringMethod, "color", "c", "term", "specify the coloring method form [normal , i3 , none] default to normal ")
	return cmd
}
