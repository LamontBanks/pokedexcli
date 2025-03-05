package httpclient

import (
	"encoding/json"
	"fmt"
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

	// Unmarshal the JSON response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &responseStruct); err != nil {
		return err
	}
	fmt.Println(responseStruct)

	return nil
}
