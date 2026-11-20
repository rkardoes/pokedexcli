package main

import(
	"fmt"
	"github.com/rkardoes/pokedexcli/internal/pokeapi"
)

type mapPageItem struct {
	Name 	string 	`json:"name"`
	URL 	string 	`json:"url"`
}

type mapPageResult struct {
	Count		int 		`json:"count"`
	Results 	[]mapPageItem 	`json:"results"`
	Next		string		`json:"next"`
	Previous	string		`json:"previous"`
}

func commandMap(cfg *config) error {
	defer cfg.changeMapPage(1)
	var e pokeapi.MapAreaEndpoint
	var err error
	var pageResult mapPageResult

	offset := 20 * (cfg.mapPage)
	q := fmt.Sprintf("limit=%d&offset=%d", 20, offset)
	if e.URL, err = pokeapi.MakeUrl("location-area/", []string{q}); err != nil {
		return err
	}
	
	resp, err := pokeapi.CallEndpoint(e) 
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	pageResult, err = pokeapi.RespToStructs[mapPageResult](resp)
	if err != nil {
		return err
	}

	if len(pageResult.Results) > 20 {
		return fmt.Errorf("map returned more than 20 location areas: (%d)", len(pageResult.Results))
	}
	for _, r := range pageResult.Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
	
}

func commandMapb(cfg *config) error {
	if cfg.mapPage - 2 < 0 {
		fmt.Printf("you're on the first page\n")
		cfg.mapPage = 0
		return nil
	}
	cfg.changeMapPage(-2)
	if err := commandMap(cfg); err != nil {
		return err
	}
	return nil
}
