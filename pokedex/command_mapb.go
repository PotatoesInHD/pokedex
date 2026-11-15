package main
import "github.com/PotatoesInHD/pokedex/internal/pokeapi"
import "fmt"


func commandMapb(cfg *config) error {
	if cfg.prevLocUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areaRes, err := pokeapi.GetAreas(*cfg.prevLocUrl)
	cfg.nextLocUrl = areaRes.Next
	cfg.prevLocUrl = areaRes.Prev
	if err != nil {
		fmt.Printf("Error after get areas: %v \n",err)
	}
	for i := 0; i < len(areaRes.Results); i++{
		fmt.Println(areaRes.Results[i].Name)
	}
	return nil
}
