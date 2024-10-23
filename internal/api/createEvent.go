package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func CreateAwEvent(apiUrl, bucketId string, payload []byte) (err error) {
	fmt.Println(apiUrl, bucketId, string(payload))
	url := fmt.Sprintf("%s/buckets/%s/events", apiUrl, bucketId)
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status code: %d", resp)
	}
	body, _ := io.ReadAll(io.Reader(resp.Body))
	fmt.Println(string(body))
	return nil
}
