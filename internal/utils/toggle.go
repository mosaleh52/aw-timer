package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mosaleh52/aw-timer/internal/api"
)

func CheckIfPaused(dbFilePath string) bool {
	fileInfo, err := os.Stat(dbFilePath)
	if err == nil && fileInfo.Size() > 0 {
		return true
	}
	return false
}

func GetPausedTodo(dbFilePath, apiUrl, bucketId string) (res api.AwResponseEvent, err error) {
	fileData, err := os.ReadFile(dbFilePath)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return api.AwResponseEvent{}, err
	}
	// need to check for  file content
	// _ , err := strconv.Atoi(string(fileData))
	// if err != nil {
	// 	fmt.Println("Error converting string to number:", err)
	// 	return
	// }
	eventId := strings.TrimSuffix(string(fileData), "\n")
	result, err := api.GetAwEvent(apiUrl, bucketId, eventId)
	if err != nil {
		fmt.Println("Error:", err)
		return api.AwResponseEvent{}, err
	}
	return result, nil
}

func CloneTodoDataAsByte(todo api.AwResponseEvent, layout string) ([]byte, error) {
	todo.Data["running"] = true
	modifiedTodo := api.AwResponseEvent{
		Timestamp: time.Now().UTC().Format(layout),
		Data:      todo.Data,
	}
	payload, err := json.Marshal(modifiedTodo)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
