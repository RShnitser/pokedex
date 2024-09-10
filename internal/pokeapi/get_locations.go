package pokeapi

import(
	"fmt"
	"encoding/json"
	"io"
	"net/http"
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

// func Test(){
// 	fmt.Println("test")
// }

func GetLocations(url *string) (locationData, error){
	res, err := http.Get(*url)
	if err != nil {
		return locationData{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locationData{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return locationData{}, err
	}
	
	data := locationData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationData{}, err
	}

	return data, nil
}