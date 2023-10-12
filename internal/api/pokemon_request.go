package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	///////////////// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok { //cache hit
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(dat, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResponse, nil
	}
	/////////////////

	// res, err := http.Get(baseURL + "/location-area/")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	// save cache
	c.cache.Add(fullURL, data)

	return pokemonResponse, nil
}
