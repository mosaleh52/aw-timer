package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getCurrentTodoCmd() *cobra.Command {
	var req string
	var apiUrl, bucketId, dateLayout string
	cmd := &cobra.Command{
		Use:   "current",
		Short: "Print specified fields",
		Run: func(cmd *cobra.Command, args []string) {
			responses := getCurrentTodo(apiUrl, bucketId, dateLayout)
			if len(responses) == 1 {
				response := responses[0]
				switch req {
				case "id":
					fmt.Println(response.ID)
					os.Exit(1)
				case "uuid":
					fmt.Println(response.Data["label"])
					os.Exit(1)
				default:
					output, err := json.Marshal(responses)
					if err != nil {
						fmt.Println("Error marshaling JSON:", err)
						return
					}
					fmt.Println(string(output))
					os.Exit(1)
				}
			} else if len(responses) == 0 {
				fmt.Println("no running todo")
			} else {
				fmt.Println("have not implemented multibel todos yet you should focus in one \nher is your todos in json")
				output, err := json.Marshal(responses)
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)
					return
				}
				fmt.Println(string(output))
			}
		},
	}

	cmd.Flags().StringVarP(&req, "require", "r", "all", "specifie the required atter from [uuid,id,label]")
	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999Z", "specify the dateLayout used for formatting in aw server")
	return cmd
}
