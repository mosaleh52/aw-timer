package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendAwQuery(apiUrl string, queryData QueryData) []AwResponseEvent {
	requestBody, err := json.Marshal(queryData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing query", err)
		os.Exit(-1)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
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
		body, _ := io.ReadAll(io.Reader(resp.Body))
		fmt.Fprintln(os.Stderr, "Response Body:", string(body))
		os.Exit(-1)
	}
	var response [][]AwResponseEvent

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		body, _ := io.ReadAll(io.Reader(resp.Body))
		fmt.Fprintln(os.Stderr, " 45 error Decoding =>", string(body), "<==", err)
		os.Exit(-1)
	}
	// the api return response as a nested array
	return response[0]
}
