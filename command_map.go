package main

import(
	"fmt"
	"https://github.com/rkardoes/pokedexcli/internal/pokeapi"
)

type mapPageItem struct {
	Name 	string 	`json: "name"`
	URL 	string 	`json: "url"`
}

type mapPageResults struct {
	Count		int
	Resutls 	[]mapPageItem
	Next		string
	Previous	string
}

func commandMap(cfg *config) error {
	defer cfg.changeMaPpage(1)
	var e pokeapi.mapAreaEndopint
	var err Error

	pageResults := make([]mapPageResults, 20)

	offset := 20 * config.page
	q := fmt.Sprintf("limit=%d&offset=%d", 20, offset)
	if e.URL, err := pokeapi.makeUrl("https://pokeapi.co/api/v2/location-area/", [q]); err != nil {
		return err
	}
	
	if resp, err := pokeapi.callEndpoint(e); err != nil {
		return err
	}

	if pageResults, err = pokeapi.respToStructs(resp, pageResults); err != nil {
		return err
	}
	
	if len(pageResults) > 1 {
		return fmt.Errorf("mapPageResults came back greater than 1: (%d)", len(pageResults))
	}
	if len(pageResults[0].Results) > 20 {
		return fmt.Errorf("map returned more than 20 location areas: (%d)", len(pageResults.Results))
	}
	for r := range pageResults[0].Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
	
}
