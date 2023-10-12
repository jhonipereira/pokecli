package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	///////////////// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok { //cache hit
		locationResponse := LocationAreas{}
		err := json.Unmarshal(dat, &locationResponse)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationResponse, nil
	}
	/////////////////

	// res, err := http.Get(baseURL + "/location-area/")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationResponse := LocationAreas{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationAreas{}, err
	}

	// save cache
	c.cache.Add(fullURL, data)

	return locationResponse, nil
}

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationName
	fullURL := baseURL + endpoint

	///////////////// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok { //cache hit
		locationResponse := LocationArea{}
		err := json.Unmarshal(dat, &locationResponse)
		if err != nil {
			return LocationArea{}, err
		}

		return locationResponse, nil
	}
	/////////////////

	// res, err := http.Get(baseURL + "/location-area/")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationResponse := LocationArea{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationArea{}, err
	}

	// save cache
	c.cache.Add(fullURL, data)

	return locationResponse, nil
}
