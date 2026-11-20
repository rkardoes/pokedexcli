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

func callEndpoint(e endpoint) *io.ReadCloser, error {
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

}

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

