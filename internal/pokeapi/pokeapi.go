package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/flying-house/pokedex/internal/pokecache"
)

// LocationArea -
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationResponse -
type LocationResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// PokemonEncounter -
type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []interface{} `json:"version_details"`
}

// LocationAreaDetail -
type LocationAreaDetail struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	GameIndex         int                `json:"game_index"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

// Client -
type Client struct {
	cache *pokecache.Cache
}

// NewClient -
func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(5 * time.Minute),
	}
}

// GetLocationAreas -
func (c *Client) GetLocationAreas(url string) (*LocationResponse, error) {
	if cacheData, found := c.cache.Get(url); found {
		var locResponse LocationResponse
		err := json.Unmarshal(cacheData, &locResponse)
		if err != nil {
			// fallthrough
		} else {
			return &locResponse, nil
		}
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	return &locationResponse, nil
}

// GetLocationAreaDetail -
func (c *Client) GetLocationAreaDetail(locationName string) (*LocationAreaDetail, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName + "/"

	if cacheData, found := c.cache.Get(url); found {
		var locationDetail LocationAreaDetail
		err := json.Unmarshal(cacheData, &locationDetail)
		if err == nil {
			return &locationDetail, nil
		}
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var locationDetail LocationAreaDetail
	err = json.Unmarshal(body, &locationDetail)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	return &locationDetail, nil
}
