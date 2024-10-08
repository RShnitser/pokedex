package pokeapi

import(
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

type allLocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocation(name string, cache *pokecache.Cache) (allLocationData, error){
	
	url := "https://pokeapi.co/api/v2/location-area/" + name
	bytes, ok := cache.Get(url)
	if !ok{
		
		res, err := http.Get(url)
		if err != nil {
			return allLocationData{}, err
		}
		bytes, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return allLocationData{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, bytes)
		}
		if err != nil {
			return allLocationData{}, err
		}
		cache.Add(url, bytes)
	}

	data := allLocationData{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return allLocationData{}, err
	}

	return data, nil
}