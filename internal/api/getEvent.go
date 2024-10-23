package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetAwEvent(apiUrl, bucketId, eventId string) (AwResponseEvent, error) {
	url := fmt.Sprintf("%s/buckets/%s/events/%s", apiUrl, bucketId, eventId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AwResponseEvent{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AwResponseEvent{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "Unexpected status code:", resp.StatusCode)
		body, _ := io.ReadAll(io.Reader(resp.Body))
		fmt.Fprintln(os.Stderr, "Response Body:", string(body))
		os.Exit(-1)
	}
	var response AwResponseEvent

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		body, _ := io.ReadAll(io.Reader(resp.Body))
		fmt.Fprintln(os.Stderr, " 37 error Decoding =>", string(body), "<==", err)
		os.Exit(-1)
	}
	return response, nil
}
