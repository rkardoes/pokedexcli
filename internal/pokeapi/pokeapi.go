package pokeapi

import(
	"fmt"
	"net/http"
	jsonv2 "encoding/json/v2"
	"time"
)

type endpoint interface {
	makeEndpoint() (*http.Request, error)
}

type MapAreaEndpoint struct {
	URL	string
}

func (e MapAreaEndpoint) makeEndpoint() (*http.Request, error){
	req, err := http.NewRequest("GET", e.URL, nil)
	if err != nil {
    		return nil, err
	}
	return req, nil	
}

func CallEndpoint(e endpoint) (*http.Response, error) {
	c := &http.Client{
		Timeout:  15 * time.Second,
	}

	req, err := e.makeEndpoint()
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code error, status code: %d", resp.StatusCode)
	}
	
	return resp, nil
}

func RespToStructs[T any](resp *http.Response) (T, error) {
	var result T
	if err := jsonv2.UnmarshalRead(resp.Body, &result); err != nil {
		return result, err
	}
	return result, nil
}

func MakeUrl(path string, queries []string) (string, error) {
	const prefix = "https://pokeapi.co/api/v2/"
	query := ""
	if path == ""{
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

