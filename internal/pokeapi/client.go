package pokeapi

// Intenal package to handle the API calls
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetLocations() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}
	defer res.Body.Close()

	var locations []Location

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locations); err != nil {
		fmt.Println("Error decoding response body.")
		return err
	}
	fmt.Println(locations)
	return err
}
