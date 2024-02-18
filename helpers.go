package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/1set/todotxt"
)

func createAwEvent(apiUrl, bucketId string, payload []byte) (err error) {
	url := fmt.Sprintf("%s/buckets/%s/events", apiUrl, bucketId)
	fmt.Println(string(payload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(responseBody))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status code: %d", resp)
	}

	return nil
}

type TodoEvent struct {
	*todotxt.Task
}

func (task TodoEvent) MarshalJSON() ([]byte, error) {
	dateLayout := "2006-01-02T15:04:05.999Z"

	type InternalTodoEvent struct {
		Label          string            `json:"label"`
		Priority       string            `json:"priority,omitempty"`
		Projects       []string          `json:"projects,omitempty"`
		Contexts       []string          `json:"contexts,omitempty"`
		AdditionalTags map[string]string `json:"additionalTags,omitempty"`
		CreatedDate    time.Time         `json:"createdDate,omitempty"`
		DueDate        time.Time         `json:"dueDate,omitempty"`
		CompletedDate  time.Time         `json:"completedDate,omitempty"`
		Completed      bool              `json:"completed,omitempty"`
		Uuid           string            `json:"uuid"`
		Running        string            `json:"running"`
	}

	internalEvent := InternalTodoEvent{
		Label:          task.Task.Todo,
		Uuid:           task.Task.AdditionalTags["uuid"],
		Running:        task.Task.AdditionalTags["running"],
		Priority:       task.Task.Priority,
		Projects:       task.Task.Projects,
		Contexts:       task.Task.Contexts,
		AdditionalTags: task.Task.AdditionalTags,
		CreatedDate:    task.Task.CreatedDate,
		DueDate:        task.Task.DueDate,
		CompletedDate:  task.Task.CompletedDate,
		Completed:      task.Task.Completed,
	}

	wrappedEvent := struct {
		Timestamp string            `json:"timestamp"`
		Data      InternalTodoEvent `json:"data"`
	}{
		Timestamp: time.Now().UTC().Format(dateLayout),
		Data:      internalEvent,
	}

	return json.Marshal(wrappedEvent)
}
