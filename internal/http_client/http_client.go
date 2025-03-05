package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
)

// Make a GET request and stores the JSON body response into the responseStruct
func Get(url string, responseStruct any) error {
	// Make the request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Read the JSON []byte
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the JSON []byte into the provided response struct
	if err := json.Unmarshal(data, &responseStruct); err != nil {
		return err
	}

	return nil
}
