package timer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mosaleh52/aw-timer/internal/api"
)

func StopTimer(apiUrl, bucketId, dateLayout, TodoId string) {
	todoResponse, err := api.GetAwEvent(apiUrl, bucketId, TodoId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting todo", err)
		os.Exit(-1)

	}
	startTime, err := time.Parse(dateLayout, todoResponse.Timestamp)
	currentTime := time.Now()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing timestamp", err)
		os.Exit(-1)
	}
	todoResponse.Duration = currentTime.Sub(startTime).Seconds()
	todoResponse.Data["running"] = false
	modifiedTodo, err := json.Marshal(todoResponse)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	err = api.CreateAwEvent(apiUrl, bucketId, modifiedTodo)
	if err != nil {
		fmt.Println(err)
	}
}
