package pokeapi

// Intenal package to handle the API calls
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous interface{}        `json:"previous"`
	Results  []LocationAreaItem `json:"results"`
}

type LocationAreaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return fmt.Errorf("Request failed: %v", err)
	}
	defer res.Body.Close()

	var locationResponse LocationAreaResponse

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locationResponse); err != nil {
		fmt.Println("Error decoding response body.")
		return err
	}
	locations := locationResponse.Results
	for i := range locations {
		fmt.Println(locations[i].Name)
	}
	return nil
}
