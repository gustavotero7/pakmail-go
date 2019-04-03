package pakmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gustavotero7/pakmail-go/models"
)

// APIURL Base API url
const APIURL = "https://api224.pakmail.com.mx/api/v1"

type client struct {
	apiURL string
	apiKey string
}

var (
	// APIKey is the Pakmail api key
	APIKey = ""
	c      *client
	once   sync.Once
)

func getClient() *client {
	if c == nil {
		once.Do(func() {
			c = &client{
				apiURL: APIURL,
				apiKey: APIKey,
			}
		})
	}
	return c
}

// Get do get request
func Get(path string) ([]byte, error) {
	return getClient().send(http.MethodGet, path, nil)
}

// Post do post request
func Post(path string, payload interface{}) ([]byte, error) {
	return getClient().send(http.MethodPost, path, payload)
}

// Put do put request
func Put(path string, payload interface{}) ([]byte, error) {
	return getClient().send(http.MethodPut, path, payload)
}

// Delete do delete request
func Delete(path string) ([]byte, error) {
	return getClient().send(http.MethodDelete, path, nil)
}

func (c *client) send(method, path string, b interface{}) ([]byte, error) {
	// Encode request body
	payload, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	// Create http request
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.apiURL, path), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	// Set api key in request
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(c.apiKey, "")
	// NOTE: Seems like basic auth isn't working, so we also set the api key as a query param
	req.URL.Query().Add("api_key", c.apiKey)

	// Do http request
	log.Printf("Calling %s: %s\n", method, req.URL.String())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// If status isn't OK (2xx) then try to decode it into a pakmail error
	// If decode fails then return decode (Unmarshal) error and nil response
	if res.StatusCode < 200 || res.StatusCode > 299 {
		pakmailErr := &models.Error{}
		if err := json.Unmarshal(body, pakmailErr); err != nil {
			return nil, err
		}
		return nil, pakmailErr
	}

	return body, nil
}

// SetClientURL re-sets api URL, by default is "https://api224.pakmail.com.mx/api/v1"
// For now this is only used for testing purposes
func SetClientURL(apiURL string) {
	getClient().apiURL = apiURL
}
