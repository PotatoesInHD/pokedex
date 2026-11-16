package main
import "github.com/PotatoesInHD/pokedex/internal/pokeapi"
import "fmt"

func commandMap(cfg *config) error {
	url := ""
	if cfg.nextLocUrl != nil {
		url = *cfg.nextLocUrl
	}
	areaRes, err := pokeapi.GetAreas(url)
	if err != nil {
		fmt.Printf("Error GetAreas in commandMap: %v\n", err)
		return err
	}
	cfg.nextLocUrl = areaRes.Next
	cfg.prevLocUrl = areaRes.Prev
	for i := 0; i < len(areaRes.Results); i++{
		fmt.Println(areaRes.Results[i].Name)
	}
	return nil
}
