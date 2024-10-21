package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func stopTimer(apiUrl, bucketId, dateLayout string, taskId int, state bool) {
	url := fmt.Sprintf("%sbuckets/%s/events/%d", apiUrl, bucketId, taskId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status: %d\n", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(jsonData["timestamp"])
	startTime, err := time.Parse(dateLayout, jsonData["timestamp"].(string))
	currentTime := time.Now()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing timestamp", err)
		os.Exit(-1)
	}
	jsonData["duration"] = currentTime.Sub(startTime).Seconds()
	// Check if 'data' is present and is a map
	if data, ok := jsonData["data"].(map[string]interface{}); ok {
		// Update "running" and "duration"
		data["running"] = state

		// Check if 'nested' is present and is a map
		if nested, ok := data["additionalTags"].(map[string]interface{}); ok {
			// Add a new property to the nested map
			nested["newProperty"] = "New Value"
		} else {
			fmt.Println("'nested' is not a map[string]interface{}")
		}
	} else {
		fmt.Println("'data' is not a map[string]interface{}")
	}

	// Marshal the modified map back to JSON
	modifiedJSON, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	err = createAwEvent(apiUrl, bucketId, modifiedJSON)
	if err != nil {
		fmt.Println(err)
	}
}
