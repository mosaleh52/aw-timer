package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ResponseEvent []struct {
	Data      map[string]interface{} `json:"data"`
	Duration  float64                `json:"duration"`
	ID        int                    `json:"id"`
	Timestamp string                 `json:"timestamp"`
}

func getCurrentTodo(apiUrl, bucketId, dateLayout string) ResponseEvent {
	url := fmt.Sprintf("%squery/", apiUrl)

	queryData := map[string]interface{}{
		"query": []string{
			fmt.Sprintf("stop = query_bucket(find_bucket(\"%s\"));", bucketId),
			"run = filter_keyvals(stop, \"running\", [true , \"true\"]);",
			"RETURN = sort_by_duration(run);",
			";",
		},
		"timeperiods": []string{
			fmt.Sprintf("%s/%s", (time.Now().Add(-24 * time.Hour)).Format(dateLayout), (time.Now().Add(24 * time.Hour)).Format(dateLayout)),
		},
	}

	requestBody, err := json.Marshal(queryData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing query", err)
		os.Exit(-1)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error making request", err)
		os.Exit(-1)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting  response", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "Unexpected status code:", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintln(os.Stderr, "Response Body:", string(body))
		os.Exit(-1)
	}
	var response []ResponseEvent

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintln(os.Stderr, "error Decoding", string(body), err)
		os.Exit(-1)
	}
	return response[0]
}
