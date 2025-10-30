package http
package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendPostRequest(url string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("received non-2xx response: %d", resp.StatusCode)
	}

	fmt.Printf("Successfully sent request to %s\n", url)
	return nil
}
