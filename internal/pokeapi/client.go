package pokeapi

// Intenal package to handle the API calls
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaItem `json:"results"`
}

type LocationAreaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(URL string) (next string, previous string, err error) {
	res, err := http.Get(URL)
	if err != nil {
		return "", "", fmt.Errorf("Request failed: %w", err)
	}
	defer res.Body.Close()

	var locationResponse LocationAreaResponse

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	if err := json.Unmarshal(data, &locationResponse); err != nil {
		fmt.Println("Error decoding response body.")
		return "", "", err
	}
	if locationResponse.Next != nil {
		next = *locationResponse.Next
	} else {
		next = ""
	}
	if locationResponse.Previous != nil {
		previous = *locationResponse.Previous
	} else {
		previous = ""
	}
	locations := locationResponse.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return next, previous, nil
}
