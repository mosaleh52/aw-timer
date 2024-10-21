package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func getCurrentTodoCmd() *cobra.Command {
	var req string
	var coloringMethod string
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
				case "label+time":
					pretyPrint(response.Data["label"].(string)+" "+timeDifferenceFromNow(response.Timestamp), "green", coloringMethod)
					os.Exit(1)
				case "label":
					pretyPrint(response.Data["label"].(string), "green", coloringMethod)
					os.Exit(1)
				case "uuid":
					fmt.Println(response.Data["uuid"])
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
					url := "http://localhost:5600/api/0/buckets/aw-stopwatch/events/" + strconv.Itoa(retrievedNumber)
					result, err := fetchEvent(url)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					pretyPrint(result.Label+" "+timeDifferenceFromNow(result.Timestamp), "yellow", coloringMethod)
					return
				}
				pretyPrint("no running todo", "red", coloringMethod)
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
	cmd.Flags().StringVarP(&coloringMethod, "color", "c", "term", "specify the coloring method form [normal , i3 , none] default to normal ")
	cmd.Flags().StringVarP(&apiUrl, "api-url", "u", "http://127.0.0.1:5600/api/0/", "specify the api url ")
	cmd.Flags().StringVarP(&bucketId, "bucket-id", "b", "aw-stopwatch", "specify the bucket-id")
	cmd.Flags().StringVarP(&dateLayout, "date-Layout", "d", "2006-01-02T15:04:05.999Z", "specify the dateLayout used for formatting in aw server")
	return cmd
}
