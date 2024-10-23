package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mosaleh52/aw-timer/internal/api"
	"github.com/mosaleh52/aw-timer/internal/helpers"
)

func HandleOnCurrentTodo(responses []api.AwResponseEvent, req, coloringMethod string) {
	response := responses[0]
	switch req {
	case "id":
		fmt.Println(response.ID)
		os.Exit(0)
	case "label+time":
		helpers.PretyPrint(response.Data["label"].(string)+" "+helpers.TimeDifferenceFromNow(response.Timestamp), "green", coloringMethod)
		os.Exit(0)
	case "label":
		helpers.PretyPrint(response.Data["label"].(string), "green", coloringMethod)
		os.Exit(0)
	case "uuid":
		fmt.Println(response.Data["uuid"])
		os.Exit(0)
	default:
		// will leave it as an array so it would be easy for scripts not not check for multibel todos
		output, err := json.Marshal(responses)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		fmt.Println(string(output))
		os.Exit(0)
	}
}

func HandleMultibleTodos(responses []api.AwResponseEvent, req, coloringMethod string) {
	// TODO:add flag to skip massage if will not implement this
	fmt.Println("have not implemented multibel todos yet you should focus in one \nher is your todos in json")
	output, err := json.Marshal(responses)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(output))
	os.Exit(0)
}

func PrintPausedTodo(dbFilePath, apiUrl, bucketId, coloringMethod string) {
	pausedTodo, err := GetPausedTodo(dbFilePath, apiUrl, bucketId)
	timeFromStart, err := strconv.Atoi(helpers.TimeDifferenceFromNow(pausedTodo.Timestamp))
	if err != nil {
		fmt.Println("Error parsing int:", err)
		os.Exit(1)
	}
	restTime := timeFromStart - int(pausedTodo.Duration)/60
	helpers.PretyPrint(pausedTodo.Data["label"].(string)+" "+strconv.Itoa(restTime), "yellow", coloringMethod)
	if err != nil {
		fmt.Println("error printing paused todo ", err)
		os.Exit(1)
	}
}
