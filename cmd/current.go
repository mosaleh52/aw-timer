package cmd

import (
	"fmt"
	"os"

	"github.com/mosaleh52/aw-timer/internal/helpers"
	"github.com/mosaleh52/aw-timer/internal/timer"
	"github.com/mosaleh52/aw-timer/internal/utils"
	"github.com/spf13/cobra"
)

func getCurrentTodoCmd() *cobra.Command {
	var requiredAttr string
	var dbFilePath string
	var coloringMethod string
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
		Use:   "current",
		Short: "Print specified fields",
		Run: func(cmd *cobra.Command, args []string) {
			responses := timer.GetCurrentTodos(apiUrl, bucketId, dateLayout)
			if len(responses) == 1 {
				utils.HandleOnCurrentTodo(responses, requiredAttr, coloringMethod)
			}
			if len(responses) > 1 {
				utils.HandleMultibleTodos(responses, requiredAttr, coloringMethod)
			}
			if utils.CheckIfPaused(dbFilePath) {
				utils.PrintPausedTodo(dbFilePath, apiUrl, bucketId, coloringMethod)
			} else if len(responses) == 0 {
				helpers.PretyPrint("no running todo", "red", coloringMethod)
				os.Exit(0)
			}
			fmt.Println("this should not happen")
			os.Exit(-1)
		},
	}

	cmd.Flags().StringVarP(&requiredAttr, "require", "r", "all", "specifie the required atter from [uuid,id,label]")
	cmd.Flags().StringVarP(&dbFilePath, "file-path", "p", "./timer.txt", "specify path for file storing paused todo")
	cmd.Flags().StringVarP(&coloringMethod, "color", "c", "term", "specify the coloring method form [normal , i3 , none] default to normal ")
	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999Z", "specify the dateLayout used for formatting in aw server")
	return cmd
}
