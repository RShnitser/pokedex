package pokeapi

import(
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

type locationData struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url *string, cache *pokecache.Cache) (locationData, error){
	
	bytes, ok := cache.Get(*url)
	if !ok{
		
		res, err := http.Get(*url)
		if err != nil {
			return locationData{}, err
		}
		bytes, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return locationData{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, bytes)
		}
		if err != nil {
			return locationData{}, err
		}
		cache.Add(*url, bytes)
	}

	data := locationData{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return locationData{}, err
	}

	return data, nil
}