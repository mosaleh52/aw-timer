package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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
	dateLayout := "2006-01-02T15:04:05.999999-07:00"

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
		Running:        "true",
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

func pretyPrint(str, color, method string) {
	Reset := "\033[0m"

	colorMap := map[string]map[string]string{
		// colors made with chatGPT take this to considreation
		"red": {
			"i3":   "#bf616a",
			"term": "\033[31m",
		},
		"green": {
			"i3":   "#a3be8c",
			"term": "\033[32m",
		},
		"yellow": {
			"i3":   "#ebcb8b",
			"term": "\033[33m",
		},
		"blue": {
			"i3":   "#81a1c1",
			"term": "\033[34m",
		},
		"magenta": {
			"i3":   "#b48ead",
			"term": "\033[35m",
		},
		"cyan": {
			"i3":   "#88c0d0",
			"term": "\033[36m",
		},
		"white": {
			"i3":   "#e5e9f0",
			"term": "\033[37m",
		},
		"black": {
			"i3":   "#2e3440",
			"term": "\033[30m",
		},
	}

	selectedColor, exists := colorMap[color]
	if !exists {
		fmt.Println("Invalid color specified")
		os.Exit(1)
	}

	var chosenColor string
	switch method {
	case "i3":
		chosenColor = selectedColor["i3"]
		// i3 format is to print output then empty line then line contain color
		fmt.Println(str + "\n\n" + chosenColor)
		os.Exit(0)
	case "term":
		chosenColor = selectedColor["term"]
		fmt.Println(chosenColor + str + Reset)
		os.Exit(0)
	case "none":
		fmt.Println(str)
		os.Exit(0)
	default:
		fmt.Println("Invalid method specified")
		os.Exit(1)
	}
}

func timeDifferenceFromNow(timestamp string) string {
	// Define the layout that matches the input timestamp format
	layout := "2006-01-02T15:04:05.999999-07:00"

	// Parse the input timestamp
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return fmt.Sprintf("Error parsing timestamp: %v", err)
	}

	// Get the current time
	now := time.Now()

	// Calculate the difference
	duration := t.Sub(now)

	// Format the duration as a string
	return formatDuration(duration)
}

func formatDuration(d time.Duration) string {
	// days := int(d.Hours() / 24)
	// hours := int(d.Hours()) % 24
	minutes := int(d.Minutes())
	// seconds := int(d.Seconds()) % 60

	// return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
	return strconv.Itoa(minutes * -1)
}

type Result struct {
	Timestamp string `json:"timestamp"`
	Label     string `json:"label"`
}

func fetchEvent(url string) (*Result, error) {
	// Construct the curl command with jq
	cmd := exec.Command("bash", "-c", fmt.Sprintf(`curl '%s' | jq '{timestamp, label: .data.label}'`, url))

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running curl command: %v", err)
	}

	// Unmarshal the JSON response into a Result struct
	var result Result
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &result, nil
}
