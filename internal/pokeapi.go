package pokeapi

import(
	"io"
	"fmt"
	"net/http"
	"encoding/json/v2"
	"time"
)

type endpoint interface {
	makeEndpoint() *http.Request, error;
}

type mapAreaEndpoint struct {
	URL	string;
}
func (e mapAreaEndpoint) makeEndpoint() *http.Request, error{
	if req, err := http.NewRequest("GET", e.url, nil); if err != nil {
		return nil, err
	}
	return req, nil
	
}

func callEndpoint(e endpoint) (*http.Response, error) {
	c := &http.Client{
		Timeout := 15 * time.Second
	}

	req, err := e.makeEndpoint()
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		reutrn nil, fmt.Errorf("HTTP status code error, status code: %d", resp.StatusCode)
	}
	
	return &resp, nil
}

func respToStructs[T any](resp *http.Respone, t []T) (t []T, error) {
	decoder := json/v2.NewDecoder(resp.Body)
	if err := decoder.Decode(&t); err != nil {
		return nil, err
	}
	return t, nil

func makeUrl(path string, queries []string) string, error {
	const prefix = "https://pokeapi.co/api/v2/"
	query := ""
	if path == nil | path == ""{
		err := fmt.Errorf("No path provided for request URL")
		return "", err
	}
	if len(queries) > 0 {
		query = "?"
		for i,  q := range queries {
			query += q
			if i != len(queries)-1 {
				query += "&"
			}
		}
	}
	return prefix + path + query, nil
}

